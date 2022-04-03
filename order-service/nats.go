package main

import (
	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
)

type NATS struct {
	nc *nats.Conn
	sc stan.Conn
}

func CreateNATS() *NATS {
	n := NATS{}

	return &n
}

func (v *NATS) Run(c *Config) error {
	nc, err := nats.Connect(c.NATSURL)
	if err != nil {
		return err
	}
	v.nc = nc
	sc, err := stan.Connect(c.ClusterID, c.ClientID, stan.NatsConn(v.nc))
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
	if v.nc != nil {
		v.nc.Close()
	}
}
