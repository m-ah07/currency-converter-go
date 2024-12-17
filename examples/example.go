package main

import (
	"fmt"
	"os/exec"
)

func main() {
	fmt.Println("Running Currency Converter Example...")

	// Replace with your own API key
	cmd := exec.Command("go", "run", "src/main.go", "your_api_key_here", "USD", "EUR")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(output))
}
