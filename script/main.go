package main

import (
	"fmt"
	"github.com/nats-io/stan.go"
	"log"
	"os"
)

func main() {
	st, err := stan.Connect("spiridonov", "daniil", stan.NatsURL("nats//localhost:4222"))
	if err != nil {
		log.Println(err)
	}

	for {
		var name string
		_, err := fmt.Scan(&name)
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
