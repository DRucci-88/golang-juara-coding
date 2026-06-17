package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
)

type Rating struct {
	Rate  float64 `json:"rate"`
	Count int     `json:"count"`
}

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Image       string  `json:"image"`
	Rating      Rating  `json:"rating"`
}

func main() {
	// 1. Initialize Resty Client
	client := resty.New()
	client.SetTimeout(15 * time.Second)

	fmt.Println("==================================================")
	fmt.Println("Fetching products from FakeStore API...")
	fmt.Println("==================================================")

	// 2. Fetch Data from API
	var products []Product
	resp, err := client.R().
		SetResult(&products).
		Get("https://fakestoreapi.com/products")

	if err != nil {
		log.Fatalf("Error executing request: %v", err)
	}

	if resp.IsError() {
		log.Fatalf("API returned error status: %s", resp.Status())
	}

	// 3. Display Data in an ASCII Table Format
	printProductsTable(products)
}

// truncateString cuts a string if it exceeds maxLen and adds "..."
func truncateString(s string, maxLen int) string {
	if len(s) > maxLen {
		return s[:maxLen-3] + "..."
	}
	return s
}

// printProductsTable formats and prints the products in a styled ASCII table
func printProductsTable(products []Product) {
	hID := "ID"
	hTitle := "Title"
	hPrice := "Price"
	hCategory := "Category"
	hRating := "Rating"

	// Initial widths based on header length
	wID := len(hID)
	wTitle := len(hTitle)
	wPrice := len(hPrice)
	wCategory := len(hCategory)
	wRating := len(hRating)

	// Struct to hold formatted row string representations
	type Row struct {
		id       string
		title    string
		price    string
		category string
		rating   string
	}

	var rows []Row
	for _, p := range products {
		title := truncateString(p.Title, 40)
		price := fmt.Sprintf("$%.2f", p.Price)
		rating := fmt.Sprintf("%.1f (⭐) / %d", p.Rating.Rate, p.Rating.Count)
		id := fmt.Sprintf("%d", p.ID)

		rows = append(rows, Row{
			id:       id,
			title:    title,
			price:    price,
			category: p.Category,
			rating:   rating,
		})

		// Update column widths if values are longer
		if len(id) > wID {
			wID = len(id)
		}
		if len(title) > wTitle {
			wTitle = len(title)
		}
		if len(price) > wPrice {
			wPrice = len(price)
		}
		if len(p.Category) > wCategory {
			wCategory = len(p.Category)
		}
		if len(rating) > wRating {
			wRating = len(rating)
		}
	}

	// Function to print horizontal border lines
	printLine := func() {
		fmt.Printf("+-%s-+-%s-+-%s-+-%s-+-%s-+\n",
			strings.Repeat("-", wID),
			strings.Repeat("-", wTitle),
			strings.Repeat("-", wPrice),
			strings.Repeat("-", wCategory),
			strings.Repeat("-", wRating),
		)
	}

	// Print Table
	printLine()
	fmt.Printf("| %-*s | %-*s | %-*s | %-*s | %-*s |\n",
		wID, hID,
		wTitle, hTitle,
		wPrice, hPrice,
		wCategory, hCategory,
		wRating, hRating,
	)
	printLine()

	for _, r := range rows {
		fmt.Printf("| %-*s | %-*s | %*s | %-*s | %-*s |\n",
			wID, r.id,
			wTitle, r.title,
			wPrice, r.price, // Right align the price
			wCategory, r.category,
			wRating, r.rating,
		)
	}
	printLine()
}
