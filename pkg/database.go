package pkg

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

// Configuração da conexão com o banco de dados PostgreSQL
type DBConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	DBName   string
	SSLMode  string
}

// Função para gerar a string de conexão
func (config *DBConfig) ConnectionString() string {
	return "postgres://" + config.User + ":" + config.Password + "@" + config.Host + ":" + config.Port + "/" + config.DBName + "?sslmode=" + config.SSLMode
}

// Função para criar uma nova conexão ao banco de dados
func NewPostgresDB(config *DBConfig) (*sql.DB, error) {
	db, err := sql.Open("postgres", config.ConnectionString())
	if err != nil {
		return nil, err
	}

	// Configurar o timeout para a conexão
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	// Verificar se a conexão está funcionando corretamente
	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Conexão com o banco de dados PostgreSQL estabelecida com sucesso.")
	return db, nil
}
