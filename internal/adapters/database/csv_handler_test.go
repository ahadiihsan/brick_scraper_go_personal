// database/csv_handler_test.go
// initial
package database

import (
	"encoding/csv"
	"fmt"
	"os"
	"testing"

	"github.com/ahadiihsan/brick_scraper_go/internal/entities"
	"github.com/stretchr/testify/assert"
)

func TestCSVHandler_SaveProducts(t *testing.T) {
	// Arrange
	handler := &CSVHandler{}
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
	err := handler.SaveProducts(products)

	// Assert
	assert.NoError(t, err, "SaveProducts should not return an error")

	// Verify the content of the generated CSV file
	file, err := os.Open("products.csv")
	assert.NoError(t, err, "Failed to open the generated CSV file")
	defer file.Close()

	// Read the CSV file
	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	assert.NoError(t, err, "Failed to read CSV rows")

	// Check the header
	assert.Equal(t, []string{"Name", "Description", "ImageLink", "Price", "Rating", "Merchant"}, rows[0])

	// Check the content of each row
	for i, product := range products {
		expectedRow := []string{
			product.Name,
			product.Description,
			product.ImageLink,
			// Format the price as a string for comparison
			fmt.Sprintf("%.2f", product.Price),
			product.Rating,
			product.Merchant,
		}
		assert.Equal(t, expectedRow, rows[i+1], "CSV row content mismatch for product %d", i+1)
	}

	// Clean up: Remove the generated CSV file
	err = os.Remove("products.csv")
	assert.NoError(t, err, "Failed to remove the generated CSV file")
}
