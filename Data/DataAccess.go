package data

import (
	"TeamProducts/Database"
	"TeamProducts/Models"
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Implement struct {
	db *sql.DB
	// c  *gin.Context
}

func (i Implement) GetAllProducts() ([]Models.Products, error) {
	productos := []Models.Products{}
	db, err := Database.GetMySqlClient()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	filas, _ := db.Query("SELECT * FROM products")

	if err != nil {
		return nil, err
	}
	defer filas.Close()

	var p Models.Products

	for filas.Next() {
		err = filas.Scan(&p.IdProduct, &p.Nombre, &p.Stock, &p.PrecioU)
		// Al escanear puede haber un error
		if err != nil {

			fmt.Println(err)
			return nil, err
		}
		// Y si no, entonces agregamos lo leído al arreglo
		productos = append(productos, p)
		// fmt.Println(productos)
		// Vacío o no, regresamos el arreglo de contactos
	}
	return productos, nil

}

func (i Implement) GetProductsById(c *gin.Context) {

}
func (i Implement) PostProducts(c *gin.Context) {

}
func (i Implement) UpdateProducts(c *gin.Context) {

}
func (i Implement) DeleteProducts(c *gin.Context) {

}
