// Purpose of this file is to just compute the success rate of hashcat
// Here we are comparing the recovered password against our original dictonary

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	dictCount := countLines("dictionary.txt")
	recoveredCount := countLines("recovered.txt")

	if dictCount == 0 {
		fmt.Println("Error: dictionary.txt is empty or missing.")
		return
	}

	if recoveredCount == 0 {
		fmt.Println("Warning: recovered.txt is empty. Did Hashcat run successfully?")
	}

	// Compute percentage
	successRate := (float64(recoveredCount) / float64(dictCount)) * 100

	fmt.Println("--- Results ---")
	fmt.Printf("Original Passwords:  %d\n", dictCount)
	fmt.Printf("Recovered Passwords: %d\n", recoveredCount)
	fmt.Printf("Success Rate:        %.2f%%\n", successRate)
}

// count non-empty lines in a file
func countLines(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		return 0
	}
	defer file.Close()

	count := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() != "" {
			count++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error reading %s: %v", filename, err)
	}

	return count
}
