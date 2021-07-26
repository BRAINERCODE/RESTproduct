package data

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type Implement struct {
	client *mongo.Client
}

func (i Implement) GetAllProducts(c *gin.Context) {

}
func (i Implement) GetProductsById(c *gin.Context) {

}
func (i Implement) PostProducts(c *gin.Context) {

}
func (i Implement) UpdateProducts(c *gin.Context) {

}
func (i Implement) DeleteProducts(c *gin.Context) {

}
