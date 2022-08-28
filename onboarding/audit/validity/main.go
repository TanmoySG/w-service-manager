package main

import (
	"fmt"
	"time"

	// "encoding/json"
	"validity/pkg/kafka"
)

func main() {

	broker := []string{
		0: "localhost:9092",
	}
	clientID := "hsgfsd"
	deadline := time.Now().Add(10 * time.Second)

	kc := kafka.Client{
		Brokers:  broker,
		ClientID: clientID,
		ReadDeadline: deadline,
	}

	tpc := "intake"
	// msg := make(map[string]interface{})

	// msg["test"] = "pass"
	// msgData, err := json.Marshal(msg)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }
	// kc.Producer(tpc, []byte("1c"), []byte(msgData))

	kc.Consumer(tpc, func(key, value []byte) {
		fmt.Printf("%s : %s", key, value)
	})
	fmt.Print("test")
}
