package kafka

import (
	"context"
	"time"

	kafka "github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
)

type Client struct {
	Brokers  *string
	ClientID *string
}

func (k Client) Consumer() {
	return
}

func (k Client) Producer(topic string, key []byte, ) {
	conn, err := kafka.DialLeader(context.Background(), "tcp", *k.Brokers, topic, 0)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, err = conn.WriteMessages(
		kafka.Message{Value: []byte("one!")},
		kafka.Message{Value: []byte("two!")},
		kafka.Message{Value: []byte("three!")},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}

}
