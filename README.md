# Slugo 🐌

A fast and efficient command-line URL slug generator written in Go, with support for Turkish characters and multiple input modes.

## Features

- ✅ Turkish character support (ç, ğ, ı, ö, ş, ü)
- ✅ Batch processing from files or stdin
- ✅ Configurable prefix, suffix, and max length
- ✅ Clipboard integration
- ✅ Interactive CLI mode
- ✅ Comprehensive error handling

## Installation

### As a Library

```bash
go get github.com/onurhan1337/slugo
```

### As a CLI Tool

```bash
go install github.com/onurhan1337/slugo/cmd/slugo@latest
```

## Usage

### As a Library

#### Basic Usage

```go
package main

import (
    "fmt"
    "github.com/onurhan1337/slugo/pkg/slug"
)

func main() {
    // Generate a simple slug
    slug, err := slug.Generate("Merhaba Dünya!")
    if err != nil {
        panic(err)
    }
    fmt.Println(slug) // Output: merhaba-dunya
}
```

#### With Options

```go
package main

import (
    "fmt"
    "github.com/onurhan1337/slugo/pkg/slug"
)

func main() {
    // Generate slug with prefix, suffix, and max length
    slug, err := slug.GenerateWithOptions("Merhaba Dünya!", "blog-", "-v2", 20)
    if err != nil {
        panic(err)
    }
    fmt.Println(slug) // Output: blog-merhaba-dunya-v2
}
```

#### Batch Processing

```go
package main

import (
    "fmt"
    "os"
    "github.com/onurhan1337/slugo/pkg/slug"
)

func main() {
    // Process from file
    processor, err := slug.NewFileBatchProcessor("input.txt")
    if err != nil {
        panic(err)
    }

    results, err := processor.Process()
    if err != nil {
        panic(err)
    }

    for _, result := range results {
        if result.Error != nil {
            fmt.Printf("Line %d: Error - %v\n", result.LineNumber, result.Error)
        } else {
            fmt.Printf("Line %d: %s -> %s\n", result.LineNumber, result.Original, result.Slug)
        }
    }
}
```

#### Batch Processing with Callback

```go
package main

import (
    "os"
    "github.com/onurhan1337/slugo/pkg/slug"
)

func main() {
    processor := slug.NewBatchProcessor(os.Stdin)
    
    err := processor.ProcessWithCallback(func(result slug.BatchResult) {
        if result.Error != nil {
            fmt.Printf("Error on line %d: %v\n", result.LineNumber, result.Error)
        } else {
            fmt.Printf("%s -> %s\n", result.Original, result.Slug)
        }
    })
    
    if err != nil {
        panic(err)
    }
}
```

### As a CLI Tool

#### Interactive Mode

```bash
slugo
```

This starts an interactive session where you can enter text and get slugs generated.

#### File Processing

```bash
slugo --file input.txt
```

Process a file with one text per line.

#### With Options

```bash
slugo --file input.txt --prefix "blog-" --suffix "-v2" --max-length 50 --copy
```

#### From Stdin

```bash
echo "Merhaba Dünya!" | slugo
```

#### Available Flags

- `--file, -f`: Input file path
- `--copy, -c`: Copy results to clipboard
- `--prefix`: Add prefix to slugs
- `--suffix`: Add suffix to slugs
- `--max-length`: Maximum slug length

## API Reference

### Package: `github.com/onurhan1337/slugo/pkg/slug`

#### Functions

- `Generate(text string) (string, error)`: Generate a basic slug
- `GenerateWithOptions(text, prefix, suffix string, maxLength int) (string, error)`: Generate slug with options

#### Types

- `BatchResult`: Result of batch processing
- `BatchProcessor`: Handles batch processing

#### Methods

- `NewBatchProcessor(reader io.Reader) *BatchProcessor`
- `NewFileBatchProcessor(filename string) (*BatchProcessor, error)`
- `Process() ([]BatchResult, error)`
- `ProcessWithCallback(callback func(BatchResult)) error`

### Package: `github.com/onurhan1337/slugo/pkg/errors`

- `ValidationError`: Custom validation error type
- `NewValidationError(field, value, err)`: Create validation error
- `IsValidationError(err)`: Check if error is validation error

## Examples

### Turkish Text Examples

```go
// Input: "Merhaba Dünya!"
// Output: "merhaba-dunya"

// Input: "İstanbul'da Güzel Bir Gün"
// Output: "istanbulda-guzel-bir-gun"

// Input: "Çocuklar & Gençler"
// Output: "cocuklar-gencler"
```

### Error Handling

```go
slug, err := slug.Generate("")
if err != nil {
    fmt.Println("Error:", err.Error())
    // Output: "Error: text consists only of whitespace"
}
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## License

This project is licensed under the MIT License - see the LICENSE file for details. 