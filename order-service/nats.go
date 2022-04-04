package main

import (
	"encoding/json"
	"github.com/nats-io/stan.go"
	"log"
)

const (
	NATS_SUBJECT_ORDER = "order"
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
	_, err := v.sc.Subscribe(NATS_SUBJECT_ORDER, func(m *stan.Msg) {
		var order Order
		err := json.Unmarshal(m.Data, &order)
		if err != nil {
			log.Println(err)
			return
		}

		err = v.db.AddOrder(order)
		if err != nil {
			log.Println(err)
			return
		}
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
