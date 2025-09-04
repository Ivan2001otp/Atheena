import React from 'react'

interface WebSocketContextType{
    isConnected : boolean;
    sendMessage : (message : any) => void;
}




const WebSocketProvider = () => {
  return (


    <div>WebSocketProvider</div>
  )
}



export default WebSocketProvider