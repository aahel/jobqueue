package main

import (
	"fmt"
	"time"

	"github.com/aahel/jobqueue/mq"
)

var (
	topic = "topic"
)

func main() {
	SubscribeTopic()
}

func SubscribeTopic() {
	m := mq.NewClient()
	m.SetCapacity(10)
	ch, err := m.Subscribe(topic)
	if err != nil {
		fmt.Println("subscribe failed")
		return
	}
	go PublishToTopic(m)
	GetMessagesFromTopic(ch, m)
	defer m.Close()
}

func PublishToTopic(c *mq.Client) {
	i := 0
	t := time.NewTicker(2 * time.Second)
	defer t.Stop()
	for range t.C {
		err := c.Publish(topic, fmt.Sprintf("messge %d", i+1))
		if err != nil {
			fmt.Println("pub message failed")
		}
		i++
	}
}

func GetMessagesFromTopic(m <-chan interface{}, c *mq.Client) {
	for {
		val := c.GetPayLoad(m)
		fmt.Printf("message is %s\n", val)
	}
}
