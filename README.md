# Tokopedia Scraper

This Golang program extracts the top 100 products of the "Mobile Phones/Handphones" category from Tokopedia. It stores the product information in both a CSV file and a PostgreSQL database.

## Requirements

- Go (Golang)
- PostgreSQL

## Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/ahadiihsan/brick_scraper_go.git
    cd brick_scraper_go
    ```

2. Install dependencies:

    ```bash
    go mod tidy
    go mod vendor
    ```

3. Set up PostgreSQL:
   
    - Create a PostgreSQL database.
    - Update the database connection string in `cmd/main.go` by updating the constants
    ```go
    const (
        dbHost     = "server.buatbesok.com"
        dbPort     = 5432
        dbUser     = "dbmaster"
        dbPassword = "32ThicRofRafro&UWufr"
        dbName     = "sandbox"
    )
    ```

4. Run the program:

    ```bash
    go run cmd/main.go
    ```

## Configuration

- Database connection string: Update the PostgreSQL connection string in `cmd/main.go`.
- Tokopedia URL: Update the Tokopedia URL in `cmd/main.go`.

## Clean Architecture

The code follows a clean architectural design philosophy, separating concerns into distinct layers:

- **Entities**: Contains the data structures used in the application.
- **Adapters**: Adapts the application to external components (database, scraping library).
- **Usecases**: Implements the application's business logic.
  
### Project Structure
```lua
    your_project/
    |-- cmd/
    |   |-- main.go
    |
    |-- internal/
    |   |-- entities/
    |   |   |-- product.go
    |   |
    |   |-- usecases/
    |   |   |-- scraper.go
    |   |
    |   |-- adapters/
    |       |-- database/
    |       |   |-- postgres_handler.go
    |       |   |-- csv_handler.go
    |       |
    |       |-- scraping/
    |           |-- tokopedia_scraper.go
    |
    |-- go.mod
    |-- go.sum
    |-- README.md
```

## Multithreading

The scraping process is executed with a minimum of 5 multithreading processes simultaneously. 

Scraping Function Explanation:

1. **TokopediaScraper Struct:**
    ```go
    type TokopediaScraper struct {
    }
    ```
    This struct represents the Tokopedia scraper and is currently empty. It serves as a placeholder for potential future functionalities.

2. **Scrape Function:**
    ```go
    func (s *TokopediaScraper) Scrape(pageToScrape string) ([]entities.Product, error) {
    ```
    The `Scrape` function is the main method responsible for initiating the scraping process. It takes a Tokopedia URL as an argument and returns a slice of `entities.Product` and an error.

3. **Synchronization:**
    ```go
    var wg sync.WaitGroup
    var mu sync.Mutex
    ```
    The `sync.WaitGroup` (`wg`) is used to wait for all goroutines to finish, and `sync.Mutex` (`mu`) is used to safely handle concurrent writes to the `products` slice.

4. **Colly Configuration:**
    ```go
    c := colly.NewCollector(
        colly.Async(true),
    )
    ```
    A new Colly collector is created with asynchronous mode enabled to allow concurrent scraping. Additionally, a `LimitRule` is set to control the parallelism to 5 requests at a time.

5. **Request Handling:**
    ```go
    c.OnRequest(func(r *colly.Request) {
        fmt.Println("Visiting ", r.URL)
    })
    ```
    The `OnRequest` callback is used to print the URL being visited.

6. **Product List Page Scraping:**
    ```go
    c.OnHTML(".css-54k5sq", func(e *colly.HTMLElement) {
        // ...
    })
    ```
    The `OnHTML` callback extracts product details from the product list page. It also initiates goroutines for visiting each product's detail page concurrently.

7. **Detail Page Scraping:**
    ```go
    detailCollector.OnHTML(".css-1m5sihj", func(ed *colly.HTMLElement) {
        // ...
    })
    ```
    Another Colly collector is created for the product detail page, and the `OnHTML` callback extracts additional details such as price and description.

8. **Concurrency:**
    ```go
    go func(e *colly.HTMLElement) {
        // ...
    }(e)
    ```
    The `go` statement launches a goroutine for each product to scrape its detail page concurrently.

9. **Scraping Pagination:**
    ```go
    c.OnScraped(func(response *colly.Response) {
        // ...
    })
    ```
    The `OnScraped` callback is triggered after each page is scraped, and it initiates scraping of the next page until the product limit is reached.

10. **Initial Page Visit:**
    ```go
    err := c.Visit(fmt.Sprintf("%s?page=%d&ob=5", pageToScrape, 1))
    ```
    The scraper starts by visiting the initial page.

11. **Concurrent Page Visits:**
    ```go
    for n := 0; n < 5; n++ {
        err := c.Visit(fmt.Sprintf("%s?page=%d&ob=5", pageToScrape, i))
        // ...
    }
    ```
    Additional pages are visited concurrently to speed up the scraping process.

12. **Waiting for Goroutines:**
    ```go
    c.Wait()
    ```
    The `Wait` method is called to wait for all concurrent scraping operations to finish.

13. **Returning Results:**
    ```go
    return products, nil
    ```
    Finally, the scraped products are returned.

The function demonstrates a concurrent web scraping approach using Colly with a focus on handling asynchronous requests, goroutine synchronization, and pagination for scraping multiple pages.

## Scraping Library

The program uses Colly as the scraping library. If you want to switch to Selenium or another library, modify the implementation in `internal/adapters/scraping/tokopedia_scraper.go`.

## Database

The program supports both CSV and PostgreSQL as storage options. The `CSVHandler` and `PostgresHandler` in `internal/adapters/database` handle saving products to CSV and PostgreSQL, respectively.

## Usage

- Run the program using the steps mentioned in the installation section.
- Check the CSV file and PostgreSQL database for the extracted product information.
- Sample result can be found in `products.csv`

## License
