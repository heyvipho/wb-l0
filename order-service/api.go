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
	order, err := v.db.GetOrder(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{})
		return
	}

	c.JSON(200, order)
}
