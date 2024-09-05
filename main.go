package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	binance "github.com/binance/binance-connector-go" // Import the correct package
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file containing API keys
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	apiKey := os.Getenv("BINANCE_API_KEY")
	secretKey := os.Getenv("BINANCE_SECRET_KEY")
	baseURL := "https://api.binance.com" // Use the real Binance API URL

	// Initialize client
	client := binance.NewClient(apiKey, secretKey, baseURL)

	// Fetch account information using the correct service method
	response, err := client.NewGetAccountService().Do(context.Background())
	if err != nil {
		log.Fatalf("Error fetching account information: %v", err)
	}

	// Calculate total balance in USDT
	var totalUSDT float64
	for _, balance := range response.Balances {
		if balance.Asset == "USDT" {
			free, err := strconv.ParseFloat(balance.Free, 64)
			if err != nil {
				log.Fatalf("Error converting free balance: %v", err)
			}
			locked, err := strconv.ParseFloat(balance.Locked, 64)
			if err != nil {
				log.Fatalf("Error converting locked balance: %v", err)
			}
			totalUSDT = free + locked
		}
	}

	// Print total balance in USDT
	fmt.Printf("Estimated balance in USDT: %.2f\n", totalUSDT)
}
