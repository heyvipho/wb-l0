package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	c, err := CreateConfig()
	if err != nil {
		log.Panicln(err)
	}

	//db, err := CreateDatabase(c)
	//if err != nil {
	//	log.Panicln(err)
	//}

	nats := CreateNATS()
	err = nats.Run(c)
	if err != nil {
		log.Panicln(err)
	}

	if c.IsProductionMode == true {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	router.StaticFile("/", "./static/index.html")

	router.GET("/api/order/:id", getOrderHandler)

	err = router.Run(c.APIAddress)
	if err != nil {
		log.Panicln(err)
	}
}
