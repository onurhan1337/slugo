package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/onurhan1337/slugo/pkg/slug"
)

func main() {
	input := `Merhaba Dünya!
İstanbul'da Güzel Bir Gün
Çocuklar & Gençler
Web Sitesi URL'si
Özel Karakterler: @#$%^&*()
`

	processor := slug.NewBatchProcessor(strings.NewReader(input))

	fmt.Println("=== Batch Processing ===")

	err := processor.ProcessWithCallback(func(result slug.BatchResult) {
		if result.Error != nil {
			fmt.Printf("Line %d: Error - %v\n", result.LineNumber, result.Error)
		} else {
			fmt.Printf("Line %d: '%s' -> '%s'\n", result.LineNumber, result.Original, result.Slug)
		}
	})

	if err != nil {
		log.Fatalf("Batch processing error: %v", err)
	}

	fmt.Println("\n=== Batch Processing (All Results) ===")

	processor2 := slug.NewBatchProcessor(strings.NewReader(input))
	results, err := processor2.Process()
	if err != nil {
		log.Fatalf("Batch processing error: %v", err)
	}

	successCount := 0
	errorCount := 0

	for _, result := range results {
		if result.Error != nil {
			errorCount++
			fmt.Printf("Line %d: Error - %v\n", result.LineNumber, result.Error)
		} else {
			successCount++
			fmt.Printf("Line %d: '%s' -> '%s'\n", result.LineNumber, result.Original, result.Slug)
		}
	}

	fmt.Printf("\nSummary: %d successful, %d errors\n", successCount, errorCount)
}
