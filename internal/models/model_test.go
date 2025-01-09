package models

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3" // Import para usar SQLite
)

func setupTestDB(t *testing.T) *sql.DB {
	// Conecta a um banco SQLite em memória
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}

	// Cria a tabela de teste
	createTableQuery := `
	CREATE TABLE products (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		price REAL NOT NULL,
		id_user INTEGER NOT NULL
	);`
	_, err = db.Exec(createTableQuery)
	if err != nil {
		t.Fatalf("Erro ao criar tabela: %v", err)
	}

	return db
}

func seedTestData(db *sql.DB, t *testing.T) {
	// Insere dados fictícios na tabela
	insertQuery := `
	INSERT INTO products (name, price, id_user) VALUES
		('Produto 1', 10.50, 1),
		('Produto 2', 20.75, 2),
		('Produto 3', 15.00, 1);`
	_, err := db.Exec(insertQuery)
	if err != nil {
		t.Fatalf("Erro ao inserir dados de teste: %v", err)
	}
}

func TestGetAll(t *testing.T) {
	// Configura o banco de dados de teste
	db := setupTestDB(t)
	defer db.Close()

	// Popula o banco com dados fictícios
	seedTestData(db, t)

	// Executa a função que será testada
	products, err := GetProductsAll(db)
	if err != nil {
		t.Fatalf("Erro ao executar GetAll: %v", err)
	}

	// Valida os resultados
	expectedCount := 3
	if len(products) != expectedCount {
		t.Errorf("Esperava %d produtos, mas obteve %d", expectedCount, len(products))
	}

}

func TestGetByID(t *testing.T) {
	// Configura o banco de dados de teste
	db := setupTestDB(t)
	defer db.Close()

	// Popula o banco com dados fictícios
	seedTestData(db, t)

	// Executa a função que será testada
	product, err := GetProductByID(db, 1)
	if err != nil {
		t.Fatalf("Erro ao executar GetByID: %v", err)
	}

	// Valida os resultados
	expectedId := 1
	if product.Id != expectedId {
		t.Errorf("Esperava ID %d, mas obteve %d", expectedId, product.Id)
	}
}

func TestCreate(t *testing.T) {
	// Configura o banco de dados de teste
	db := setupTestDB(t)
	defer db.Close()

	// Cria um novo produto
	newProduct := Product{
		Name:   "Produto Novo",
		Price:  99.99,
		IdUser: 3,
	}

	// Executa a função que será testada
	id, err := CreateProduct(db, newProduct)
	if err != nil {
		t.Fatalf("Erro ao executar Create: %v", err)
	}

	// Valida os resultados
	if id == 0 {
		t.Error("ID do produto criado é inválido")
	}
}

func TestUpdate(t *testing.T) {
	// Configura o banco de dados de teste
	db := setupTestDB(t)
	defer db.Close()

	// Popula o banco com dados fictícios
	seedTestData(db, t)

	// Cria um produto para atualizar
	productToUpdate := Product{
		Id:     1,
		Name:   "Produto Atualizado",
		Price:  49.99,
		IdUser: 2,
	}

	// Executa a função que será testada
	err := UpdateProduct(db, productToUpdate)
	if err != nil {
		t.Fatalf("Erro ao executar Update: %v", err)
	}
}

func TestDelete(t *testing.T) {
	// Configura o banco de dados de teste
	db := setupTestDB(t)
	defer db.Close()

	// Popula o banco com dados fictícios
	seedTestData(db, t)

	// Executa a função que será testada
	err := DeleteProduct(db, 1)
	if err != nil {
		t.Fatalf("Erro ao executar Delete: %v", err)
	}
}
