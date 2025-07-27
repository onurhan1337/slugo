# Slugo 🐌

A fast and efficient command-line URL slug generator written in Go, with support for Turkish characters and multiple input modes.

## Features ✨

- **Turkish Character Support**: Automatically converts Turkish characters (ç, ğ, ı, ö, ş, ü) to their Latin equivalents
- **Multiple Input Modes**:
  - Interactive mode for real-time slug generation
  - File processing mode for batch operations
  - Standard input (pipe) mode for integration with other tools
- **Clipboard Integration**: Copy generated slugs directly to system clipboard
- **Customization Options**:
  - Add prefixes and suffixes to slugs
  - Set maximum slug length
  - Batch processing with detailed results
- **Cross-Platform**: Works on macOS, Linux, and Windows
- **Error Handling**: Comprehensive validation and error reporting

## Installation 📦

### Prerequisites
- Go 1.23.4 or higher

### Build from Source
```bash
git clone https://github.com/yourusername/slugo.git
cd slugo
go build -o slugo cmd/main.go
```

### Install Globally
```bash
go install ./cmd/main.go
```

## Usage 🚀

### Interactive Mode
Start the interactive slug generator:
```bash
./slugo
# or
go run cmd/main.go
```

### File Processing
Process a text file with one text per line:
```bash
./slugo --file input.txt
# or
go run cmd/main.go --file input.txt
```

### Pipe Input
Process text from standard input:
```bash
echo "Merhaba Dünya" | ./slugo
# or
cat input.txt | ./slugo
```

### Command Line Options

| Option | Short | Description | Example |
|--------|-------|-------------|---------|
| `--file` | - | Process text from file | `--file input.txt` |
| `--copy` | `-c` | Copy result to clipboard | `--copy` |
| `--prefix` | - | Add prefix to slug | `--prefix blog-` |
| `--suffix` | - | Add suffix to slug | `--suffix -v2` |
| `--max-length` | - | Set maximum slug length | `--max-length 50` |

## Examples 📝

### Basic Usage
```bash
$ ./slugo
----- Slugo Slug Generator -----
Metin girin ve Enter'a basın.
Çıkmak için 'exit' veya 'quit' yazın.
--------------------------------
Metin girin: Merhaba Dünya
Durum: Başarılı
Slug: merhaba-dunya
Mesaj: Slug başarıyla oluşturuldu.
```

### Turkish Character Conversion
```bash
$ echo "Türkçe Karakterler" | ./slugo
turkce-karakterler
```

### With Prefix and Suffix
```bash
$ echo "Web Sitesi URL'si" | ./slugo --prefix blog- --suffix -2024
blog-web-sitesi-urlsi-2024
```

### Batch Processing
```bash
$ cat test.txt | ./slugo --copy
Satır 1: Merhaba Dünya -> merhaba-dunya
Satır 2: Türkçe Karakterler -> turkce-karakterler
Satır 3: Web Sitesi URL'si -> web-sitesi-urlsi
Satır 4: Özel Karakterler: @#$%^&*() -> ozel-karakterler
Satır 5: Çok Uzun Bir Başlık Bu Slug Çok Uzun Olacak ve Maksimum Uzunluk Testi İçin Kullanılacak -> cok-uzun-bir-baslik-bu-slug-cok-uzun-olacak-ve-maksimum-uzunluk-testi-icin-kullanilacak
Satır 6: Kısa -> kisa

--- İşlem Özeti ---
Toplam satır: 6
Başarılı: 6
Hatalı: 0
------------------
✓ 6 slug clipboard'a kopyalandı
```

## Project Structure 📁

```
slugo/
├── cmd/
│   └── main.go              # Application entry point
├── internal/
│   ├── slug/
│   │   ├── slug.go          # Core slug generation logic
│   │   └── batch.go         # Batch processing functionality
│   └── ui/
│       ├── cli.go           # Interactive mode and output formatting
│       ├── cli_handler.go   # Command-line interface handler
│       ├── batch.go         # Batch output formatting
│       └── clipboard.go     # Cross-platform clipboard operations
├── pkg/
│   └── errors/
│       └── errors.go        # Custom error types and handling
└── test.txt                 # Sample input file for testing
```

## Slug Generation Rules 🔧

1. **Character Conversion**:
   - Turkish characters: ç→c, ğ→g, ı→i, ö→o, ş→s, ü→u
   - All text converted to lowercase
   - Non-alphanumeric characters replaced with spaces

2. **Formatting**:
   - Multiple spaces collapsed to single space
   - Leading/trailing spaces removed
   - Spaces replaced with hyphens
   - Multiple consecutive hyphens collapsed to single hyphen

3. **Validation**:
   - Empty or whitespace-only input rejected
   - Input containing only invalid characters rejected

## Error Handling 🛡️

The application provides detailed error reporting:

- **Validation Errors**: Clear messages for invalid input
- **File Errors**: Descriptive messages for file access issues
- **Clipboard Errors**: Graceful handling of clipboard operations

## Contributing 🤝

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License 📄

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments 🙏

- Built with Go 1.23.4
- Cross-platform clipboard support using native system commands
- Turkish language support for international users

---

**Slugo** - Making URL slugs simple and efficient! 🐌✨ 