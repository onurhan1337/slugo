package ui

import (
	"flag"
	"os"

	"github.com/onurhan1337/slugo/internal/slug"
	apperrors "github.com/onurhan1337/slugo/pkg/errors"
)

type InputMode int

const (
	ModeInteractive InputMode = iota
	ModeFile
	ModeStdin
)

type CLIHandler struct {
	mode            InputMode
	filename        string
	copyToClipboard bool
	prefix          string
	suffix          string
	maxLength       int
}

func NewCLIHandler() *CLIHandler {
	return &CLIHandler{}
}

func (ch *CLIHandler) ParseFlags() {
	flag.StringVar(&ch.filename, "file", "", "File path to read from")
	flag.BoolVar(&ch.copyToClipboard, "copy", false, "Copy result to clipboard")
	flag.BoolVar(&ch.copyToClipboard, "c", false, "Copy result to clipboard (short form)")
	flag.StringVar(&ch.prefix, "prefix", "", "Add prefix to slug (e.g., blog-)")
	flag.StringVar(&ch.suffix, "suffix", "", "Add suffix to slug (e.g., -v2)")
	flag.IntVar(&ch.maxLength, "max-length", 0, "Maximum slug length (0 = unlimited)")
	flag.Parse()
}

func (ch *CLIHandler) DetectMode() InputMode {
	if ch.filename != "" {
		return ModeFile
	}

	stat, _ := os.Stdin.Stat()
	isPiped := (stat.Mode() & os.ModeCharDevice) == 0

	if isPiped {
		return ModeStdin
	}

	return ModeInteractive
}

func (ch *CLIHandler) Run() error {
	ch.ParseFlags()
	ch.mode = ch.DetectMode()

	switch ch.mode {
	case ModeFile:
		return ch.runFileMode()
	case ModeStdin:
		return ch.runStdinMode()
	case ModeInteractive:
		return ch.runInteractiveMode()
	default:
		return apperrors.NewValidationError("mode", ch.mode, apperrors.ErrInvalidInput)
	}
}

func (ch *CLIHandler) runFileMode() error {
	processor, err := slug.NewFileBatchProcessor(ch.filename)
	if err != nil {
		return err
	}

	results, err := processor.Process()
	if err != nil {
		return err
	}

	results = ch.applyPrefixAndSuffix(results)
	results = ch.applyMaxLength(results)

	PrintBatchResultsWithSummary(results, ch.copyToClipboard)
	return nil
}

func (ch *CLIHandler) runStdinMode() error {
	processor := slug.NewBatchProcessor(os.Stdin)

	results, err := processor.Process()
	if err != nil {
		return err
	}

	results = ch.applyPrefixAndSuffix(results)
	results = ch.applyMaxLength(results)

	PrintBatchResultsWithSummary(results, ch.copyToClipboard)
	return nil
}

func (ch *CLIHandler) runInteractiveMode() error {
	RunInteractiveMode(ch.copyToClipboard, ch.prefix, ch.suffix, ch.maxLength)
	return nil
}

func (ch *CLIHandler) applyPrefixAndSuffix(results []slug.BatchResult) []slug.BatchResult {
	for i := range results {
		if results[i].Error == nil {
			results[i].Slug = ch.prefix + results[i].Slug + ch.suffix
		}
	}
	return results
}

func (ch *CLIHandler) applyMaxLength(results []slug.BatchResult) []slug.BatchResult {
	if ch.maxLength <= 0 {
		return results
	}

	for i := range results {
		if results[i].Error == nil && len(results[i].Slug) > ch.maxLength {
			results[i].Slug = results[i].Slug[:ch.maxLength]
		}
	}
	return results
}
