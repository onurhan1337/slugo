package main

import (
	"fmt"
	"os"

	"slugo/internal/ui"
	apperrors "slugo/pkg/errors"
)

func main() {
	handler := ui.NewCLIHandler()

	if err := handler.Run(); err != nil {
		if validationErr, ok := apperrors.IsValidationError(err); ok {
			fmt.Printf("Hata: %s\n", validationErr.Message)
		} else {
			fmt.Printf("Beklenmeyen hata: %v\n", err)
		}
		os.Exit(1)
	}
}
