// internal/adapters/database/postgres_handler.go
// initial
package database

import (
	"errors"

	"github.com/ahadiihsan/brick_scraper_go/internal/entities"
	"github.com/jinzhu/gorm"
)

// PostgresHandler implements DatabaseHandler
type PostgresHandler struct {
	DB *gorm.DB
}

func (h *PostgresHandler) SaveProducts(products []entities.Product) error {
	if h.DB == nil {
		return errors.New("database connection is nil")
	}

	for _, product := range products {
		if err := h.DB.Create(&product).Error; err != nil {
			return err
		}
	}

	return nil
}
