package main

import (
	"github.com/gin-gonic/gin"
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
	c.JSON(200, gin.H{
		"id": c.Param("id"),
	})
}
