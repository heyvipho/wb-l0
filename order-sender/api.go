package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
)

type API struct {
	config *Config
	nats   *NATS
}

func CreateAPI(config *Config, nats *NATS) *API {
	api := API{
		config: config,
		nats:   nats,
	}

	return &api
}

func (v *API) Run() error {
	if v.config.IsProductionMode == true {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	router.StaticFile("/", "./static/index.html")

	router.POST("/api/order", v.postOrderHandler)

	err := router.Run(v.config.APIAddress)

	return err
}

func (v *API) postOrderHandler(c *gin.Context) {
	jsonData, err := c.GetRawData()
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{})
		return
	}

	var order Order
	err = json.Unmarshal(jsonData, &order)
	if err != nil {
		c.JSON(400, gin.H{})
		return
	}

	err = v.nats.PublishOrder(order)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{})
		return
	}

	c.JSON(201, gin.H{
		"id": 0,
	})
}
