package main

import (
	"fmt"
	"log"

	"github.com/onurhan1337/slugo/pkg/slug"
)

func main() {
	texts := []string{
		"Merhaba Dünya!",
		"İstanbul'da Güzel Bir Gün",
		"Çocuklar & Gençler",
		"Web Sitesi URL'si",
		"Özel Karakterler: @#$%^&*()",
	}

	fmt.Println("=== Basic Slug Generation ===")
	for _, text := range texts {
		slug, err := slug.Generate(text)
		if err != nil {
			log.Printf("Error generating slug for '%s': %v", text, err)
			continue
		}
		fmt.Printf("'%s' -> '%s'\n", text, slug)
	}

	fmt.Println("\n=== Slug with Options ===")
	text := "Merhaba Dünya!"
	slugWithOptions, err := slug.GenerateWithOptions(text, "blog-", "-v2", 20)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	fmt.Printf("'%s' -> '%s'\n", text, slugWithOptions)
}
