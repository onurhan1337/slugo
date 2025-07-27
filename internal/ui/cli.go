package ui

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"slugo/internal/slug"
	apperrors "slugo/pkg/errors"
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
	fmt.Printf("Durum: Başarılı\n")
	fmt.Printf("Slug: %s\n", output.Slug)
	fmt.Printf("Mesaj: %s\n\n", output.Message)
}

func PrintError(err error) {
	if validationErr, ok := apperrors.IsValidationError(err); ok {
		errOutput := ErrorOutput{
			Type:    "ValidationError",
			Message: "Giriş metni doğrulama hatası. Lütfen kontrol edin.",
			Details: validationErr,
		}
		fmt.Printf("Durum: %s\n", errOutput.Type)
		fmt.Printf("Mesaj: %s\n", errOutput.Message)
		fmt.Printf("Detay: Alan: '%s', Değer: '%v', Neden: %s\n\n",
			validationErr.Field, validationErr.Value, validationErr.Message)
	} else {
		errOutput := ErrorOutput{
			Type:    "UnknownError",
			Message: "Beklenmeyen bir hata oluştu.",
			Details: err.Error(),
		}
		fmt.Printf("Durum: %s\n", errOutput.Type)
		fmt.Printf("Mesaj: %s\n", errOutput.Message)
		fmt.Printf("Detay: %s\n\n", errOutput.Details)
	}
}

func RunInteractiveMode(copyToClipboard bool, prefix string, suffix string, maxLength int) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("----- Slugo Slug Generator -----")
	fmt.Println("Metin girin ve Enter'a basın.")
	fmt.Println("Çıkmak için 'exit' veya 'quit' yazın.")
	fmt.Println("Dosyadan okumak için: go run cmd/main.go --file dosya.txt")
	if copyToClipboard {
		fmt.Println("✓ Clipboard kopyalama aktif")
	}
	if prefix != "" {
		fmt.Printf("✓ Önek: %s\n", prefix)
	}
	if suffix != "" {
		fmt.Printf("✓ Sonek: %s\n", suffix)
	}
	if maxLength > 0 {
		fmt.Printf("✓ Maksimum uzunluk: %d karakter\n", maxLength)
	}
	fmt.Println("--------------------------------")

	for {
		fmt.Print("Metin girin: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if strings.ToLower(input) == "exit" || strings.ToLower(input) == "quit" {
			fmt.Println("Çıkış yapılıyor. Güle güle!")
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
			Message: "Slug başarıyla oluşturuldu.",
		}
		PrintSuccess(successOutput)

		if copyToClipboard {
			if err := CopyToSystemClipboard(finalSlug); err != nil {
				fmt.Printf("Clipboard'a kopyalama hatası: %v\n", err)
			} else {
				fmt.Printf("✓ Slug clipboard'a kopyalandı\n")
			}
		}
	}
}
