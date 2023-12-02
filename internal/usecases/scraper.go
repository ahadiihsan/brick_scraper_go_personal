// initial
package usecases

import (
	"log"
)

// ProductScraperUsecase contains the business logic for handling products
type ProductScraperUsecase struct {
	Scraper    Scraper
	SQLHandler DatabaseHandler
	CSVHandler DatabaseHandler
}

// NewProductScraperUsecase creates a new instance of ProductScraperUsecase with the provided dependencies
func NewProductScraperUsecase(scraper Scraper, sqlHandler DatabaseHandler, csvHandler DatabaseHandler) *ProductScraperUsecase {
	return &ProductScraperUsecase{
		Scraper:    scraper,
		SQLHandler: sqlHandler,
		CSVHandler: csvHandler,
	}
}

// ScrapeAndSaveProducts scrapes products using the provided Scraper and saves them using the provided DatabaseHandler
func (uc *ProductScraperUsecase) ScrapeAndSaveProducts(url string) {
	// Scrape data
	products, err := uc.Scraper.Scrape(url)
	if err != nil {
		log.Printf("Error scraping %s: %v", url, err)
		return
	}

	// Save to database
	err = uc.SQLHandler.SaveProducts(products)
	if err != nil {
		log.Printf("Error saving to database: %v", err)
		return
	}

	// Save to csv
	err = uc.CSVHandler.SaveProducts(products)
	if err != nil {
		log.Printf("Error saving to database: %v", err)
		return
	}
}
