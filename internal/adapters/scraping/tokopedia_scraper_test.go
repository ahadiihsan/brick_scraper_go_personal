package scraping

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokopediaScraper_Scrape(t *testing.T) {
	// Arrange
	scraper := &TokopediaScraper{}
	url := "https://www.tokopedia.com/p/handphone-tablet/handphone"

	// Act
	products, err := scraper.Scrape(url)

	// Assert
	assert.NoError(t, err, "Scraping should not return an error")
	assert.NotEmpty(t, products, "Scraping should return at least one product")

	// Check details of each scraped product
	for i, product := range products {
		t.Logf("Product %d:", i+1)
		t.Logf("Name: %s", product.Name)
		t.Logf("Description: %s", product.Description)
		t.Logf("ImageLink: %s", product.ImageLink)
		t.Logf("Price: %f", product.Price)
		t.Logf("Rating: %s", product.Rating)
		t.Logf("Merchant: %s", product.Merchant)
		t.Logf("--------")

		// Additional assertions for each product
		assert.NotEmpty(t, product.Name, "Product name should not be empty")
		assert.NotEmpty(t, product.ImageLink, "Product image link should not be empty")
		assert.NotEmpty(t, product.Merchant, "Product merchant should not be empty")
	}
}
