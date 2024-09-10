package migrations

import (
	"database/sql"
	"log"
	"os"
	"path/filepath" // Importar o pacote path/filepath

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type MigrationService struct {
	db *sql.DB
}

func NewMigrationService(db *sql.DB) *MigrationService {
	return &MigrationService{db: db}
}

func (ms *MigrationService) MigrateUp() error {
	// Debug: Verificar o diretório de trabalho atual
	workingDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Erro ao obter diretório de trabalho: %v", err)
	}

	// Construir o caminho absoluto para o arquivo de migração
	absMigrationPath := filepath.Join(workingDir, "../db/migrations/")
	absMigrationPath = "file://" + absMigrationPath

	driver, err := postgres.WithInstance(ms.db, &postgres.Config{})
	if err != nil {
		return err
	}
	m, err := migrate.NewWithDatabaseInstance(absMigrationPath, "postgres", driver)
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	log.Println("Migrações aplicadas com sucesso")
	return nil
}

func (ms *MigrationService) MigrateDown() error {
	// Debug: Verificar o diretório de trabalho atual
	workingDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Erro ao obter diretório de trabalho: %v", err)
	}
	log.Printf("Diretório de trabalho atual: %s", workingDir)

	driver, err := postgres.WithInstance(ms.db, &postgres.Config{})
	if err != nil {
		return err
	}

	// Construir o caminho absoluto para o arquivo de migração
	absMigrationPath := filepath.Join(workingDir, "../db/migrations/")
	absMigrationPath = "file://" + absMigrationPath

	m, err := migrate.NewWithDatabaseInstance(absMigrationPath, "postgres", driver)
	if err != nil {
		return err
	}

	if err := m.Down(); err != nil {
		return err
	}

	log.Println("Migrações revertidas com sucesso")
	return nil
}
