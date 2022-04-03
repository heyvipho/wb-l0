package main

import (
	"github.com/gin-gonic/gin"
)

type API struct {
	config *Config
	db     *Database
}

func CreateAPI(config *Config, db *Database) *API {
	api := API{
		config: config,
		db:     db,
	}

	return &api
}

func (v *API) Run() error {
	if v.config.IsProductionMode == true {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	router.StaticFile("/", "./static/index.html")

	router.GET("/api/order/:id", v.getOrderHandler)

	err := router.Run(v.config.APIAddress)

	return err
}

func (v *API) getOrderHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"id": c.Param("id"),
	})
}
