package bussineslogic

import (
	data "TeamProducts/Data"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ImplementBussines struct {
}

var Data data.Implement

func (i ImplementBussines) GetAll(c *gin.Context) {

	Array, err := Data.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {

		c.JSON(http.StatusOK, gin.H{
			"values": Array,
		})
	}

}
