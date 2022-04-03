package main

import (
	"github.com/nats-io/stan.go"
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
