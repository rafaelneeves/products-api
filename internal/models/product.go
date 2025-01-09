package models

import (
	"database/sql"
	"log"
)

type Product struct {
	Id     int
	Name   string
	Price  float64
	IdUser int
}

func GetProductsAll(db *sql.DB) ([]Product, error) {
	query := "SELECT * FROM products"
	rows, err := db.Query(query)
	if err != nil {
		log.Printf("Erro ao listar produtos: %v", err)
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var product Product
		if err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.IdUser); err != nil {
			log.Printf("Erro ao escanear produto: %v", err)
			continue
		}
		products = append(products, product)
	}

	return products, nil
}

func GetProductByID(db *sql.DB, id int) (Product, error) {
	query := "SELECT * FROM products WHERE id = ?"
	row := db.QueryRow(query, id)

	var product Product
	if err := row.Scan(&product.Id, &product.Name, &product.Price, &product.IdUser); err != nil {
		log.Printf("Erro ao buscar produto: %v", err)
		return Product{}, err
	}

	return product, nil
}

func CreateProduct(db *sql.DB, product Product) (int, error) {
	query := "INSERT INTO products (name, price, id_user) VALUES (?, ?, ?)"
	result, err := db.Exec(query, product.Name, product.Price, product.IdUser)
	if err != nil {
		log.Printf("Erro ao inserir produto: %v", err)
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("Erro ao buscar ID do produto inserido: %v", err)
		return 0, err
	}

	return int(id), nil
}

// TODO: Corrigir o método Update está zerando os valores quando não passado
func UpdateProduct(db *sql.DB, product Product) error {
	query := "UPDATE products SET name = ?, price = ?, id_user = ? WHERE id = ?"
	_, err := db.Exec(query, product.Name, product.Price, product.IdUser, product.Id)
	if err != nil {
		log.Printf("Erro ao atualizar produto: %v", err)
		return err
	}

	return nil
}

func DeleteProduct(db *sql.DB, id int) error {
	query := "DELETE FROM products WHERE id = ?"
	_, err := db.Exec(query, id)
	if err != nil {
		log.Printf("Erro ao deletar produto: %v", err)
		return err
	}

	return nil
}
