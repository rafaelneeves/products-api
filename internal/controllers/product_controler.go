package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"products-api/internal/models"
	"strconv"
)

type ProductController struct {
	DB *sql.DB
}

func NewProductController(db *sql.DB) *ProductController {
	return &ProductController{DB: db}
}

func (p *ProductController) GetProductsAll(w http.ResponseWriter, r *http.Request) {
	products, err := models.GetProductsAll(p.DB)

	if err != nil {
		http.Error(w, "Erro ao listar os Produtos", http.StatusInternalServerError)
		return
	}

	if len(products) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func (p *ProductController) GetProductByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}
	product, err := models.GetProductByID(p.DB, id)

	if err != nil {
		http.Error(w, "Erro ao buscar o Produto", http.StatusInternalServerError)
		return
	}

	if product.Id == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func (p *ProductController) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)

	if err != nil {
		http.Error(w, "Erro ao decodificar o Produto", http.StatusBadRequest)
		return
	}

	id, err := models.CreateProduct(p.DB, product)

	if err != nil {
		http.Error(w, "Erro ao criar o Produto", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int{"Criado o protudo de ID: ": id})
}

func (p *ProductController) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)

	if err != nil {
		http.Error(w, "Erro ao decodificar o Produto", http.StatusBadRequest)
		return
	}

	err = models.UpdateProduct(p.DB, product)

	if err != nil {
		http.Error(w, "Erro ao atualizar o Produto", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Produto atualizado com sucesso!")
}
func (p *ProductController) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	err = models.DeleteProduct(p.DB, id)

	if err != nil {
		http.Error(w, "Erro ao deletar o Produto", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Produto deletado com sucesso!")
}
