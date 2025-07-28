package ui

import (
	"fmt"
	"strings"

	"github.com/onurhan/slugo/internal/slug"
)

type BatchOutput struct {
	Results []slug.BatchResult
	Total   int
	Success int
	Errors  int
}

func PrintBatchResult(result slug.BatchResult) {
	if result.Error != nil {
		fmt.Printf("Line %d: ", result.LineNumber)
		PrintError(result.Error)
	} else {
		fmt.Printf("Line %d: %s -> %s\n", result.LineNumber, result.Original, result.Slug)
	}
}

func PrintBatchSummary(output BatchOutput) {
	fmt.Printf("\n--- Processing Summary ---\n")
	fmt.Printf("Total lines: %d\n", output.Total)
	fmt.Printf("Successful: %d\n", output.Success)
	fmt.Printf("Errors: %d\n", output.Errors)
	fmt.Printf("------------------------\n")
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
		fmt.Println("No successful slugs found to copy.")
		return
	}

	clipboardText := strings.Join(successfulSlugs, "\n")

	if err := CopyToSystemClipboard(clipboardText); err != nil {
		fmt.Printf("Clipboard copy error: %v\n", err)
	} else {
		fmt.Printf("✓ %d slugs copied to clipboard\n", len(successfulSlugs))
	}
}
