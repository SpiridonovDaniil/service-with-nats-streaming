package main

import (
	"github.com/nats-io/stan.go"
	"log"
	"os"
)

func main() {
	st, err := stan.Connect("spiridonov", "daniil", stan.NatsURL("nats://localhost:4222"))
	if err != nil {
		log.Println(err)
	}

	for {
		data, err := os.ReadFile("model.json")
		if err != nil {
			log.Println(err)
		}
		err = st.Publish("message", data)
		if err != nil {
			log.Println(err)
		}
	}
}
