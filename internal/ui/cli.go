package ui

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/onurhan/slugo/internal/slug"
	apperrors "github.com/onurhan/slugo/pkg/errors"
)

type SuccessOutput struct {
	Slug    string `json:"slug"`
	Message string `json:"message,omitempty"`
}

type ErrorOutput struct {
	Type    string `json:"type"`
	Message string `json:"message"`
	Details any    `json:"details,omitempty"`
}

func PrintSuccess(output SuccessOutput) {
	fmt.Printf("Status: Success\n")
	fmt.Printf("Slug: %s\n", output.Slug)
	fmt.Printf("Message: %s\n\n", output.Message)
}

func PrintError(err error) {
	if validationErr, ok := apperrors.IsValidationError(err); ok {
		errOutput := ErrorOutput{
			Type:    "ValidationError",
			Message: "Input text validation error. Please check your input.",
			Details: validationErr,
		}
		fmt.Printf("Status: %s\n", errOutput.Type)
		fmt.Printf("Message: %s\n", errOutput.Message)
		fmt.Printf("Details: Field: '%s', Value: '%v', Reason: %s\n\n",
			validationErr.Field, validationErr.Value, validationErr.Message)
	} else {
		errOutput := ErrorOutput{
			Type:    "UnknownError",
			Message: "An unexpected error occurred.",
			Details: err.Error(),
		}
		fmt.Printf("Status: %s\n", errOutput.Type)
		fmt.Printf("Message: %s\n", errOutput.Message)
		fmt.Printf("Details: %s\n\n", errOutput.Details)
	}
}

func RunInteractiveMode(copyToClipboard bool, prefix string, suffix string, maxLength int) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("----- Slugo Slug Generator -----")
	fmt.Println("Enter text and press Enter.")
	fmt.Println("Type 'exit' or 'quit' to exit.")
	fmt.Println("To read from file: slugo --file filename.txt")
	if copyToClipboard {
		fmt.Println("✓ Clipboard copying enabled")
	}
	if prefix != "" {
		fmt.Printf("✓ Prefix: %s\n", prefix)
	}
	if suffix != "" {
		fmt.Printf("✓ Suffix: %s\n", suffix)
	}
	if maxLength > 0 {
		fmt.Printf("✓ Max length: %d characters\n", maxLength)
	}
	fmt.Println("--------------------------------")

	for {
		fmt.Print("Enter text: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if strings.ToLower(input) == "exit" || strings.ToLower(input) == "quit" {
			fmt.Println("Exiting. Goodbye!")
			break
		}

		generatedSlug, err := slug.Generate(input)

		if err != nil {
			PrintError(err)
			continue
		}

		finalSlug, err := slug.Generate(prefix + generatedSlug + suffix)

		if err != nil {
			PrintError(err)
			continue
		}

		if maxLength > 0 && len(finalSlug) > maxLength {
			finalSlug = finalSlug[:maxLength]
		}

		successOutput := SuccessOutput{
			Slug:    finalSlug,
			Message: "Slug generated successfully.",
		}
		PrintSuccess(successOutput)

		if copyToClipboard {
			if err := CopyToSystemClipboard(finalSlug); err != nil {
				fmt.Printf("Clipboard copy error: %v\n", err)
			} else {
				fmt.Printf("✓ Slug copied to clipboard\n")
			}
		}
	}
}
