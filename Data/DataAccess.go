package data

import (
	"TeamProducts/Database"
	"TeamProducts/Models"
	"database/sql"
	"fmt"
)

type Implement struct {
	db *sql.DB
}

var productModel Models.Products
var productos []Models.Products

func (i Implement) GetAllProducts() ([]Models.Products, error) {
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

	for filas.Next() {
		err = filas.Scan(&productModel.IdProduct, &productModel.Nombre, &productModel.Stock, &productModel.PrecioU, &productModel.IdCiudad)
		if err != nil {
			return nil, err
		}
		productos = append(productos, productModel)
	}
	return productos, nil

}

func (i Implement) GetProductsById(id int64) (Models.Products, error) {

	db, err := Database.GetMySqlClient()
	if err != nil {
		return Models.Products{}, err
	}
	defer db.Close()
	filas, err := db.Query("SELECT * FROM products WHERE IdProduct = ?", id)

	if err != nil {
		return Models.Products{}, err
	}
	defer filas.Close()

	for filas.Next() {

		err = filas.Scan(&productModel.IdProduct, &productModel.Nombre, &productModel.Stock, &productModel.PrecioU, &productModel.IdCiudad)
		if err != nil {

			return Models.Products{}, err
		}

	}
	return productModel, nil

}
func (i Implement) PostProducts(c Models.Products) error {

	db, err := Database.GetMySqlClient()
	if err != nil {
		return err
	}
	defer db.Close()

	sentencia, err := db.Prepare("INSERT INTO products (IdProduct,Nombre, Stock, PrecioU, IdCiudad) VALUES(?,?, ?, ?,?)")
	if err != nil {
		return err
	}
	defer sentencia.Close()
	_, err = sentencia.Exec(c.IdProduct, c.Nombre, c.Stock, c.PrecioU, c.IdCiudad)

	if err != nil {
		return err
	}

	return nil

}
func (i Implement) UpdateProducts(c Models.Products, id int) error {
	db, err := Database.GetMySqlClient()
	if err != nil {
		return err
	}
	defer db.Close()

	sentenciaPreparada, err := db.Prepare("UPDATE products SET Nombre = ?, Stock = ?, PrecioU = ?,IdCiudad = ? WHERE IdProduct = ?")
	fmt.Println("Sentencia", sentenciaPreparada)
	if err != nil {
		return err
	}

	defer sentenciaPreparada.Close()
	// Pasar argumentos en el mismo orden que la consulta
	_, err = sentenciaPreparada.Exec(c.Nombre, c.Stock, c.PrecioU, c.IdCiudad, id)
	return err // Ya sea nil o sea un error, lo manejaremos desde donde hacemos la llamada
}
func (i Implement) DeleteProducts(id int) error {
	db, err := Database.GetMySqlClient()
	if err != nil {
		return err
	}
	defer db.Close()

	sentenciaPreparada, err := db.Prepare("DELETE FROM products WHERE IdProduct = ?")
	if err != nil {
		return err
	}
	defer sentenciaPreparada.Close()

	_, err = sentenciaPreparada.Exec(id)
	if err != nil {
		return err
	}
	return nil
}
