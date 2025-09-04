package websockets

// models for socket implementation
import (
	_entities "atheena/internals/entities"
	_util "atheena/internals/util"
	"encoding/json"
	"log"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	once sync.Once
	singleton *Hub
)
type Client struct {
	ID string
	Conn *websocket.Conn
	Hub *Hub
	UserID primitive.ObjectID
	Send chan []byte
}

type Hub struct {
	clients map[*Client]bool
	broadcast chan []byte
	Register chan *Client
	Unregister chan *Client;
	mutex sync.Mutex
}

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second
	
	// Time allowed to read the next pong message from the peer.
    pongWait = 60 * time.Second

    // Send pings to peer with this period. Must be less than pongWait.
    pingPeriod = (pongWait * 9) / 10

    // Maximum message size allowed from peer.
    maxMessageSize = 512
)

func (c *Client) ReadPump() {
	defer func () {
		c.Hub.Unregister <- c
		c.Conn.Close()
	} ()


	c.Conn.SetReadLimit(maxMessageSize)
	c.Conn.SetReadDeadline(time.Now().Add(pongWait))
	c.Conn.SetPongHandler(func(string) error {
		c.Conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil;
	})


	for {
		_, message , err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error on read-pump : %v", err)
			}

			break;
		}

		c.Hub.broadcast <- message
	}
}


func (c *Client) WritePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func(){
		ticker.Stop()
		c.Conn.Close()
	}()


	for {
		select {
		case message , ok :=  <- c.Send :
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))

			if (!ok) {
				log.Println("something went wrong in write pump, some parsing issue")
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return;
			}

			w, err := c.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return;
			}
			w.Write(message)

			n := len(c.Send)
			for i:= 0; i< n;i++ {
				w.Write([]byte{'\n'})
				w.Write(<- c.Send)
			}


			if err := w.Close(); err != nil {
				log.Println("Something went wrong in write pump while writing message")
				return;
			}
		
		case <- ticker.C: 
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				log.Println(err.Error())
				return;
			}
		}
	}

}

type WebSocketMessage struct {
	Type string `json:"type"`
	Payload interface{} `json:"payload"`
}


func newHub() *Hub {
	log.Println("new hub instantiated.");
	return & Hub{
		clients: make(map[*Client]bool),
		broadcast: make(chan []byte),
		Register: make(chan *Client),
		Unregister: make(chan *Client),
	}
}


func GetSocketHub() *Hub {
	once.Do(func ()  {
		singleton = newHub()
		go singleton.Run()
	})

	log.Println("Got a new socket hub !");
	return singleton;
}


func (h *Hub) Run() {
	for {
		select {
		case client := <- h.Register : 
			h.mutex.Lock()
			h.clients[client] = true;
			log.Println("✅ client registered.")
			h.mutex.Unlock()

		
		case client := <- h.Unregister:
			h.mutex.Lock()

			if _,ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.Send)
				log.Println("✅ unregistered client.");
			}
			h.mutex.Unlock()


		case message := <- h.broadcast:
			h.mutex.Lock()
			for client := range h.clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(h.clients, client)
				}
			}

			h.mutex.Unlock()
		}
	}
}


func (h *Hub) BroadcastNotification(notification *_entities.ApprovalTypeNotification) {
	message := WebSocketMessage {
		Type: _util.APPROVAL_NOTIFICATION,
		Payload : notification,
	}

	jsonMessage, err := json.Marshal(message);
	if err != nil {
		log.Printf("error marshaling notification: %v", err)
		return;
	}

	h.broadcast <- jsonMessage
}


func (h *Hub) SendToUser(userID primitive.ObjectID, notification interface{}) {
	message := WebSocketMessage{								// *_entities.ApprovalTypeNotification
		Type: _util.APPROVAL_NOTIFICATION,
		Payload: notification,
	}

	jsonMessage, err := json.Marshal(message)
	if err != nil {
		log.Printf("Error marshaling notification : %v", err)
		return;
	}

	h.mutex.Lock();
	defer h.mutex.Unlock();

	for client := range h.clients {
		if client.UserID == userID {
			client.Send <- jsonMessage
			log.Println("✅ sent the message to user")
		}
	}
}