package main

import (
	"fmt"
	"github.com/nats-io/stan.go"
)

type NATS struct {
	sc     stan.Conn
	config *Config
	db     *Database
}

func CreateNATS(config *Config, db *Database) *NATS {
	n := NATS{
		config: config,
		db:     db,
	}

	return &n
}

func (v *NATS) Subscribe() error {
	_, err := v.sc.Subscribe("foo", func(m *stan.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})

	return err
}

func (v *NATS) Connect() error {
	sc, err := stan.Connect(v.config.ClusterID, v.config.ClientID, stan.NatsURL(v.config.NATSURL))
	if err != nil {
		return err
	}
	v.sc = sc
	return nil
}

func (v *NATS) Close() {
	if v.sc != nil {
		v.sc.Close()
	}
}
