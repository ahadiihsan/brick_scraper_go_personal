package usecases

import "github.com/ahadiihsan/brick_scraper_go/internal/entities"

// Scraper interface for different scrapers
type Scraper interface {
	Scrape(url string) ([]entities.Product, error)
}

// DatabaseHandler interface for different database handlers
type DatabaseHandler interface {
	SaveProducts(products []entities.Product) error
}
