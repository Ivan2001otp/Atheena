package util

import (
	"fmt"
	"log"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	uuid_generator uuid.UUID
	mu sync.Mutex
)

func GenerateRandomUUID() string {
	mu.Lock()
	defer mu.Unlock();
	
	uuid_generator = uuid.New()
	log.Println("Generated new UUID : ", uuid_generator.String())
	return uuid_generator.String();
}

func GenerateObjectID() primitive.ObjectID {
	return primitive.NewObjectID()
}

func GenerateCreateDateTime()( time.Time, error) {
	return time.Parse(time.RFC3339, time.Now().Format(time.RFC3339));
}


func ToUpper(str string) string {
	return strings.ToUpper(str);
}

func FormatDateTime(t time.Time) string {
	return t.Format("2006-01-02 03:04:05 PM");
}


// decoding url algorithm
var base62Chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
var charIndex = make(map[rune]int)

func Init() {
	for i,c := range base62Chars {
		charIndex[c] = i;
	}
}

func fromBase62(s string) *big.Int{
	result := big.NewInt(0)
	base := big.NewInt(62)

	for _,c := range s {
		result.Mul(result, base);
		result.Add(result, big.NewInt(int64(charIndex[c])))
	}

	return result;
}

func DecodeToHex(input string) string {
	num := fromBase62(strings.TrimLeft(input, "0"));
	hexStr := fmt.Sprintf("%x", num);

	
	if (len(hexStr) %2 != 0) {
		fmt.Println("Even length of hex id")
		hexStr = "0" + hexStr;
	} else {
		fmt.Println("Not even length of hex");
	}

	return hexStr;
}