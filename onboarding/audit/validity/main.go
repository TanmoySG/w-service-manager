package main

import (
	"fmt"

	"encoding/json"
	"validity/pkg/kafka"
)

func main() {

	broker := "localhost:9092"
	clientID := "hsgfd"

	kc := kafka.Client{
		Brokers:  &broker,
		ClientID: &clientID,
	}

	tpc := "intake"
	msg := make(map[string]interface{})

	msg["test"] = "pass"
	msgData, err := json.Marshal(msg)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	kc.Producer(tpc, []byte("1"), []byte(msgData))

	kc.Consumer(tpc)
	fmt.Print("test")
}
