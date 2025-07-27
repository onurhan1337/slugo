package ui

import (
	"fmt"
	"slugo/internal/slug"
	"strings"
)

type BatchOutput struct {
	Results []slug.BatchResult
	Total   int
	Success int
	Errors  int
}

func PrintBatchResult(result slug.BatchResult) {
	if result.Error != nil {
		fmt.Printf("Satır %d: ", result.LineNumber)
		PrintError(result.Error)
	} else {
		fmt.Printf("Satır %d: %s -> %s\n", result.LineNumber, result.Original, result.Slug)
	}
}

func PrintBatchSummary(output BatchOutput) {
	fmt.Printf("\n--- İşlem Özeti ---\n")
	fmt.Printf("Toplam satır: %d\n", output.Total)
	fmt.Printf("Başarılı: %d\n", output.Success)
	fmt.Printf("Hatalı: %d\n", output.Errors)
	fmt.Printf("------------------\n")
}

func ProcessBatchResults(results []slug.BatchResult) BatchOutput {
	output := BatchOutput{
		Results: results,
		Total:   len(results),
	}

	for _, result := range results {
		if result.Error != nil {
			output.Errors++
		} else {
			output.Success++
		}
		PrintBatchResult(result)
	}

	return output
}

func PrintBatchResultsWithSummary(results []slug.BatchResult, copyToClipboard bool) {
	output := ProcessBatchResults(results)
	PrintBatchSummary(output)

	if copyToClipboard {
		copyResultsToClipboard(results)
	}
}

func copyResultsToClipboard(results []slug.BatchResult) {
	var successfulSlugs []string

	for _, result := range results {
		if result.Error == nil {
			successfulSlugs = append(successfulSlugs, result.Slug)
		}
	}

	if len(successfulSlugs) == 0 {
		fmt.Println("Kopyalanacak başarılı slug bulunamadı.")
		return
	}

	clipboardText := strings.Join(successfulSlugs, "\n")

	if err := CopyToSystemClipboard(clipboardText); err != nil {
		fmt.Printf("Clipboard'a kopyalama hatası: %v\n", err)
	} else {
		fmt.Printf("✓ %d slug clipboard'a kopyalandı\n", len(successfulSlugs))
	}
}
