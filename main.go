package main

import (
	bussineslogic "TeamProducts/BussinesLogic"

	"github.com/gin-gonic/gin"
)

func main() {
	var i bussineslogic.ImplementBussines
	c := gin.Default()

	c.GET("/productos", i.GetAll)             //Done
	c.GET("/productos/:id", i.GetProductById) //Done
	// c.POST("/productos", i.PostProducts)      //Done
	c.PUT("/productos/:id", i.Update) //Done
	// c.DELETE("productos/:id", i.DeleteProducts) //Done
	c.Run(":8090")

}
