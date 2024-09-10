package main

import (
	"go-product/internal/migrations"
	"go-product/internal/server"
	"go-product/pkg"
	"log"
)

func main() {
	// Configuração do banco de dados
	dbConfig := &pkg.DBConfig{
		User:     "postgres",
		Password: "1234",
		Host:     "localhost",
		Port:     "5432",
		DBName:   "postgres",
		SSLMode:  "disable", // ou "require", se necessário
	}

	// Conectar ao banco de dados PostgreSQL
	db, err := pkg.NewPostgresDB(dbConfig)
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}

	// Criar serviço de migração
	migrationService := migrations.NewMigrationService(db)

	// Configurar e iniciar servidor Gin
	r := server.SetupRouter(migrationService)

	log.Println("Servidor rodando na porta :8080")
	r.Run(":8080")
}
