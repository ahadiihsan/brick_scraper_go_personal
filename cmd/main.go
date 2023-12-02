// initial
package main

import (
	"fmt"
	"log"

	"github.com/ahadiihsan/brick_scraper_go/internal/adapters/database"
	"github.com/ahadiihsan/brick_scraper_go/internal/adapters/scraping"
	"github.com/ahadiihsan/brick_scraper_go/internal/entities"
	"github.com/ahadiihsan/brick_scraper_go/internal/usecases"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	dbHost     = "server.buatbesok.com"
	dbPort     = 5432
	dbUser     = "dbmaster"
	dbPassword = "32ThicRofRafro&UWufr"
	dbName     = "sandbox"
)

func main() {

	// Initialize concrete implementations
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)
	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.AutoMigrate(&entities.Product{})

	postgresHandler := &database.PostgresHandler{DB: db}
	csvHandler := &database.CSVHandler{}
	scraper := &scraping.TokopediaScraper{}

	// Create instances of the use case with the concrete implementations
	productScraperUsecase := usecases.NewProductScraperUsecase(scraper, postgresHandler, csvHandler)

	// Use the use case to save products to both CSV and PostgreSQL
	productScraperUsecase.ScrapeAndSaveProducts("https://www.tokopedia.com/p/handphone-tablet/handphone")

	fmt.Println("Processing completed.")
}
