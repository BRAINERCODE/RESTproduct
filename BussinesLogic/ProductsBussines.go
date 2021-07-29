package bussineslogic

import (
	data "TeamProducts/Data"
	"TeamProducts/Models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ImplementBussines struct {
	Data data.Implement
}

var p Models.Products
var Array []Models.Products

func (i ImplementBussines) GetAll(c *gin.Context) {

	Array, err := i.Data.GetAllProducts()
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

func (i ImplementBussines) GetProductById(c *gin.Context) {
	idstr := c.Param("id")
	id, err := strconv.ParseInt(idstr, 10, 64)
	GetById, err := i.Data.GetProductsById(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {

		c.JSON(http.StatusOK, gin.H{
			"values": GetById,
		})
	}
}
func (i ImplementBussines) SaveProduct(c *gin.Context) {

	err := c.ShouldBindJSON(&p)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	i.Data.PostProducts(p)
}

func (i ImplementBussines) Update(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.ParseInt(idParam, 10, 64)
	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	i.Data.UpdateProducts(p, int(id))

}

func (i ImplementBussines) Delete(c *gin.Context) {

	idParam := c.Param("id")
	id, _ := strconv.ParseInt(idParam, 10, 64)
	err := i.Data.DeleteProducts(int(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
}
