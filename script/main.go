package main

import (
	"fmt"
	"log"
	"os"

	"github.com/nats-io/stan.go"
)

func main() {
	st, err := stan.Connect("test-cluster", "Daniil-pub", stan.NatsURL("http://localhost:4222"))
	if err != nil {
		log.Println(err)
	}

	for {
		var name string
		_, err := fmt.Scanln(&name)
		data, err := os.ReadFile(name)
		if err != nil {
			log.Println(err)
		}
		err = st.Publish("message", data)
		if err != nil {
			log.Println(err)
		}
	}
}
