package main

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
)

type RegisterRequest struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
	Data    struct {
		ID       uint   `json:"id"`
		FullName string `json:"full_name"`
		Email    string `json:"email"`
		Role     string `json:"role"`
	} `json:"data,omitempty"`
}

func main() {
	client := resty.New()

	// List of 10 real Indonesian names
	names := []string{
		"Budi Santoso",
		"Siti Aminah",
		"Joko Widodo",
		"Dewi Lestari",
		"Ahmad Dhani",
		"Dian Sastrowardoyo",
		"Bambang Pamungkas",
		"Sri Mulyani",
		"Agus Harimurti",
		"Megawati Soekarnoputri",
	}

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Starting registration of 10 customer users...")

	for i, name := range names {
		// Create a clean email name based on the full name
		cleanName := strings.ToLower(strings.ReplaceAll(name, " ", "."))
		// Use a random number suffix to avoid email conflicts on multiple runs
		email := fmt.Sprintf("%s.%d@example.com", cleanName, rand.Intn(100000))
		password := "password123"

		reqBody := RegisterRequest{
			FullName: name,
			Email:    email,
			Password: password,
		}

		var result RegisterResponse

		resp, err := client.R().
			SetHeader("Content-Type", "application/json").
			SetHeader("x-api-key", "juara-coding-super-secret").
			SetBody(reqBody).
			SetResult(&result).
			Post("http://localhost:8080/api/v1/register")

		if err != nil {
			log.Fatalf("Error registering user %d (%s): %v", i+1, name, err)
		}

		if resp.IsError() {
			fmt.Printf("[%d/10] FAILED to register %s. Status: %d, Response: %s\n", i+1, name, resp.StatusCode(), resp.String())
		} else {
			fmt.Printf("[%d/10] SUCCESS: Registered ID %d | Name: %s | Email: %s | Role: %s\n", 
				i+1, result.Data.ID, result.Data.FullName, result.Data.Email, result.Data.Role)
		}
	}

	fmt.Println("Finished registration process.")
}
