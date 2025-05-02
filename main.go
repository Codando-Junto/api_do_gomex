// filepath: /Users/rgomes/code/pessoal/mentoria/api_do_gomex/main.go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Codando-Junto/api_do_gomex/config"
	"github.com/Codando-Junto/api_do_gomex/handlers"
)

// Livro representa a estrutura da tabela no banco de dados
type Livro struct {
	ID            uint `gorm:"primaryKey"`
	Titulo        string
	Autor         string
	AnoPublicacao int
	Genero        string
}

func main() {

	// Conectar ao banco de dados
	db := config.ConnectDatabase()

	// Migração automática para criar a tabela com base na struct Livro
	db.AutoMigrate(&Livro{})

	// Inserir um livro de exemplo
	livro := Livro{Titulo: "Dom Casmurro", Autor: "Machado de Assis", AnoPublicacao: 1899, Genero: "Romance"}
	db.Create(&livro)

	// Buscar o primeiro livro no banco de dados
	var primeiroLivro Livro
	db.First(&primeiroLivro)

	fmt.Printf("Primeiro livro: %s, Autor: %s\n", primeiroLivro.Titulo, primeiroLivro.Autor)

	mux := http.NewServeMux() // Create a new ServeMux

	// Register routes
	mux.HandleFunc("/", handlers.HomeHandler)
	mux.HandleFunc("/health", handlers.HealthHandler)

	// Configure the server
	srv := config.NewServer(mux)

	// Log a message before starting the server
	log.Printf("Server is running on %s", srv.Addr)

	// Start the server
	log.Fatal(srv.ListenAndServe())
}
