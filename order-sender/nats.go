package main

import (
	"encoding/json"
	"github.com/nats-io/stan.go"
)

const (
	NATS_SUBJECT_ORDER = "order"
)

type NATS struct {
	sc     stan.Conn
	config *Config
}

func CreateNATS(config *Config) *NATS {
	n := NATS{
		config: config,
	}

	return &n
}

func (v *NATS) PublishOrder(obj Order) error {
	jsonObj, err := json.Marshal(obj)
	if err != nil {
		return err
	}

	err = v.sc.Publish(NATS_SUBJECT_ORDER, jsonObj)

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
