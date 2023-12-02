// database/postgres_handler_test.go
// initial
package database

import (
	"testing"

	"github.com/ahadiihsan/brick_scraper_go/internal/entities"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPostgresHandler_SaveProducts(t *testing.T) {
	// Create an in-memory SQLite database for testing
	db, err := gorm.Open("sqlite3", ":memory:")
	require.NoError(t, err, "Failed to connect to the database")
	defer db.Close()

	// AutoMigrate the Product model
	db.AutoMigrate(&entities.Product{})

	// Create the PostgresHandler with the in-memory database
	handler := &PostgresHandler{DB: db}

	// Arrange
	products := []entities.Product{
		{
			Name:        "Product1",
			Description: "Description1",
			ImageLink:   "ImageLink1",
			Price:       10.5,
			Rating:      "4.5",
			Merchant:    "Merchant1",
		},
		{
			Name:        "Product2",
			Description: "Description2",
			ImageLink:   "ImageLink2",
			Price:       20.75,
			Rating:      "3.8",
			Merchant:    "Merchant2",
		},
	}

	// Act
	err = handler.SaveProducts(products)

	// Assert
	assert.NoError(t, err, "SaveProducts should not return an error")

	// Retrieve the saved products from the database for verification
	var savedProducts []entities.Product
	db.Find(&savedProducts)

	// Check the number of saved products
	assert.Len(t, savedProducts, len(products), "Number of saved products mismatch")

	// Check the content of each saved product
	for i, savedProduct := range savedProducts {
		expectedProduct := products[i]
		assert.Equal(t, expectedProduct.Name, savedProduct.Name, "Name mismatch for product %d", i+1)
		assert.Equal(t, expectedProduct.Description, savedProduct.Description, "Description mismatch for product %d", i+1)
		assert.Equal(t, expectedProduct.ImageLink, savedProduct.ImageLink, "ImageLink mismatch for product %d", i+1)
		assert.Equal(t, expectedProduct.Price, savedProduct.Price, "Price mismatch for product %d", i+1)
		assert.Equal(t, expectedProduct.Rating, savedProduct.Rating, "Rating mismatch for product %d", i+1)
		assert.Equal(t, expectedProduct.Merchant, savedProduct.Merchant, "Merchant mismatch for product %d", i+1)
	}
}
