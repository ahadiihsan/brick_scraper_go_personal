package scraping

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"

	"github.com/ahadiihsan/brick_scraper_go/internal/entities"
	"github.com/gocolly/colly/v2"
)

// TokopediaScraper implements Scraper using Colly
type TokopediaScraper struct {
}

func (s *TokopediaScraper) Scrape(pageToScrape string) ([]entities.Product, error) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var products []entities.Product

	// current page
	i := 1
	// max products to scrape
	limit := 100

	// initializing a Colly instance
	c := colly.NewCollector(
		// turning on the asynchronous request mode in Colly
		colly.Async(true),
	)

	// make sure at least 5 requests are allowed to run in parallel
	_ = c.Limit(&colly.LimitRule{
		// limit the parallel requests to 5 request at a time
		Parallelism: 5,
	})

	// setting a valid User-Agent header
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36"

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting ", r.URL)
	})

	// Extract product details from the product list page
	c.OnHTML(".css-54k5sq", func(e *colly.HTMLElement) {

		wg.Add(1)
		// Extract the product detail link from each product asynchronously
		go func(e *colly.HTMLElement) {
			defer wg.Done()

			detailURL := e.Request.AbsoluteURL(e.Attr("href"))
			if strings.Contains(detailURL, "ta.tokopedia.com/promo/v1/clicks") {
				return
			}

			merchant := e.ChildText(".css-vbihp9 span:nth-child(2)")

			// Create a new collector for the detail page
			detailCollector := colly.NewCollector()
			detailCollector.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36"

			detailCollector.OnRequest(func(r *colly.Request) {
				fmt.Println("Visiting Detail Page ", r.URL)
			})

			// Extract product details from the detail page
			detailCollector.OnHTML(".css-1m5sihj", func(ed *colly.HTMLElement) {

				priceStr := strings.Replace(strings.TrimSpace(strings.TrimSpace(ed.ChildText("div[data-testid=lblPDPDetailProductPrice]"))), "Rp", "", -1)
				price, err := strconv.ParseFloat(strings.Replace(priceStr, ".", "", -1), 64)
				if err != nil {
					log.Printf("Error parsing price: %v", err)
				}

				rating := ""
				rating = ed.ChildText(".css-bczdt6 [data-testid=lblPDPDetailProductRatingNumber]")

				product := entities.Product{
					Name:        ed.ChildText("h1[data-testid=lblPDPDetailProductName]"),
					ImageLink:   ed.ChildAttr("img[data-testid=PDPMainImage]", "src"),
					Description: ed.ChildText("div[data-testid=lblPDPDescriptionProduk]"),
					Rating:      rating,
					Price:       price,
					Merchant:    merchant,
				}

				mu.Lock()
				if len(products) < limit {
					products = append(products, product)
				}
				mu.Unlock()
			})

			// Visit the detail page
			err := detailCollector.Visit(detailURL)
			if err != nil {
				log.Println(err)
			}

		}(e)
	})

	c.OnScraped(func(response *colly.Response) {
		// product scraping limit
		if len(products) < limit {
			// incrementing the iteration counter
			i++

			// visiting a new page
			_ = c.Visit(fmt.Sprintf("%s?page=%d&ob=5", pageToScrape, i))
		}
	})

	// registering all pages to scrape
	// deploys 5 goroutines to scrape the pages
	for n := 0; n < 5; n++ {
		err := c.Visit(fmt.Sprintf("%s?page=%d&ob=5", pageToScrape, i))
		i++
		if err != nil {
			continue
		}

	}

	// wait for tColly to visit all pages
	c.Wait()

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Println("Scraping completed.")
	return products, nil
}
