package main

import (
	"log"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	config, err := CreateConfig()
	if err != nil {
		log.Panicln(err)
	}

	nats := CreateNATS(config)
	err = nats.Connect()
	if err != nil {
		log.Panicln(err)
	}
	defer nats.Close()

	api := CreateAPI(config, nats)
	err = api.Run()
	if err != nil {
		log.Panicln(err)
	}
}
