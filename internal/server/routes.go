package server

import (
	"go-product/internal/migrations"

	"github.com/gin-gonic/gin"
)

func SetupRouter(migrationService *migrations.MigrationService) *gin.Engine {
	r := gin.Default()

	r.GET("/migrate", func(c *gin.Context) {
		if err := migrationService.MigrateUp(); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "Migrações aplicadas com sucesso"})
	})

	r.GET("/migrate/down", func(c *gin.Context) {
		if err := migrationService.MigrateDown(); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "Migrações desfeitas com sucesso"})
	})

	return r
}
