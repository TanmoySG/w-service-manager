package kafka

import (
	"context"
	"time"
	"os"
	"os/signal"
	"syscall"
	"fmt"

	kafka "github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
)

type Client struct {
	Brokers  []string
	ClientID string
	ReadDeadline time.Time
}

func (k Client) Consumer(topic string, callback func([]byte, []byte)) {

	signals := make(chan os.Signal, 1)

	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())

	// go routine for getting signals asynchronously
	go func() {
		sig := <-signals
		fmt.Println("Got signal: ", sig)
		cancel()
	}()


	config := kafka.ReaderConfig{
		Brokers:  k.Brokers,
		GroupID:  k.ClientID,
		Topic:    topic,
		MaxWait:  500 * time.Millisecond,
		MinBytes: 1,
		MaxBytes: 10e3,
	}

	r := kafka.NewReader(config)

	// fmt.Println("Consumer configuration: ", config)

	defer func() {
		err := r.Close()
		if err != nil {
			fmt.Println("Error closing consumer: ", err)
			return
		}
		fmt.Println("Consumer closed")
	}()

	for {
		m, err := r.ReadMessage(ctx)
		if err != nil {
			fmt.Println("Error reading message: ", err)
			break
		}
		
		callback(m.Key, m.Value)
	}
}

func (k Client) Producer(topic string, key []byte, value []byte) {
	conn, err := kafka.DialLeader(context.Background(), "tcp", k.Brokers[0], topic, 0)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, err = conn.WriteMessages(
		kafka.Message{
			Key:   key,
			Value: value,
		},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}

}
