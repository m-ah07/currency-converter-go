# Currency Converter (Go)

A simple and efficient currency converter written in Go (Golang) using real-time exchange rates from the ExchangeRate-API.

## Features
- Fetches real-time exchange rates for multiple currencies.
- Lightweight and fast CLI tool.
- Easy to use and extend.

## Directory Structure
```plaintext
currency-converter-go/
├── src/
│   └── main.go              # Main logic for the currency converter
├── examples/
│   └── example.go           # Demonstration of using the currency converter
├── .gitignore               # Git ignore file
└── README.md                # Documentation file
```

## Prerequisites
- Go installed on your system.
- Free API key from [ExchangeRate-API](https://www.exchangerate-api.com).

## Installation
1. Clone the repository:

    ```bash
    git clone https://github.com/marwan-ahmed-23/currency-converter-go.git
    cd currency-converter-go
    ```

2. Run the project: Replace `your_api_key_here` with a valid API key.

    ```bash
    go run src/main.go <API_KEY> <FROM_CURRENCY> <TO_CURRENCY>
    ```

3. Example:

    ```bash
    go run src/main.go your_api_key_here USD EUR
    ```

    Output:

    ```bash
    1 USD = 0.85 EUR
    ```

## Contributing
Contributions are welcome! Feel free to fork the repository and submit a pull request.
