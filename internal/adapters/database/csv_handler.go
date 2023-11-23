package database

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/ahadiihsan/brick_scraper_go/internal/entities"
)

// CSVHandler implements DatabaseHandler
type CSVHandler struct {
}

// SaveProducts saves products to a CSV file
func (h *CSVHandler) SaveProducts(products []entities.Product) error {
	// Create or open the CSV file
	file, err := os.Create("products.csv")
	if err != nil {
		return fmt.Errorf("failed to create or open CSV file: %v", err)
	}
	defer file.Close()

	// Create a CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write CSV header
	header := []string{"Name", "Description", "ImageLink", "Price", "Rating", "Merchant"}
	if err := writer.Write(header); err != nil {
		return fmt.Errorf("failed to write CSV header: %v", err)
	}

	// Write each product to the CSV file
	for _, product := range products {
		row := []string{
			product.Name,
			product.Description,
			product.ImageLink,
			fmt.Sprintf("%.2f", product.Price), // Format price as string
			product.Rating,
			product.Merchant,
		}

		if err := writer.Write(row); err != nil {
			return fmt.Errorf("failed to write CSV row: %v", err)
		}
	}

	return nil
}
