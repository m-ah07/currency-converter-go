package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// ExchangeRate holds the API response data
type ExchangeRate struct {
	Rates map[string]float64 `json:"rates"`
}

// fetchExchangeRate fetches exchange rates from the API
func fetchExchangeRate(apiURL string) (*ExchangeRate, error) {
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var rate ExchangeRate
	if err := json.NewDecoder(resp.Body).Decode(&rate); err != nil {
		return nil, err
	}
	return &rate, nil
}

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: go run main.go <API_KEY> <FROM_CURRENCY> <TO_CURRENCY>")
		return
	}

	apiKey := os.Args[1]
	fromCurrency := os.Args[2]
	toCurrency := os.Args[3]

	apiURL := fmt.Sprintf("https://v6.exchangerate-api.com/v6/%s/latest/%s", apiKey, fromCurrency)

	rates, err := fetchExchangeRate(apiURL)
	if err != nil {
		fmt.Println("Error fetching exchange rates:", err)
		return
	}

	exchangeRate, ok := rates.Rates[toCurrency]
	if !ok {
		fmt.Println("Error: Unsupported currency")
		return
	}

	fmt.Printf("1 %s = %.2f %s\n", fromCurrency, exchangeRate, toCurrency)
}
