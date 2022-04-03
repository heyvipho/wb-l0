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

	db := CreateDatabase(config)

	nats := CreateNATS(config, db)
	err = nats.Connect()
	if err != nil {
		log.Panicln(err)
	}
	defer nats.Close()

	err = nats.Subscribe()
	if err != nil {
		log.Panicln(err)
	}

	api := CreateAPI(config, db)
	err = api.Run()
	if err != nil {
		log.Panicln(err)
	}
}
