package main

import (
	"fmt"
	"os"

	"github.com/onurhan1337/slugo/internal/ui"
	apperrors "github.com/onurhan1337/slugo/pkg/errors"
)

func main() {
	handler := ui.NewCLIHandler()

	if err := handler.Run(); err != nil {
		if validationErr, ok := apperrors.IsValidationError(err); ok {
			fmt.Printf("Error: %s\n", validationErr.Message)
		} else {
			fmt.Printf("Unexpected error: %v\n", err)
		}
		os.Exit(1)
	}
}
