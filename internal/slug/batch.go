package slug

import (
	"bufio"
	"io"
	"os"

	apperrors "github.com/onurhan1337/slugo/pkg/errors"
)

type BatchResult struct {
	LineNumber int
	Original   string
	Slug       string
	Error      error
}

type BatchProcessor struct {
	reader io.Reader
}

func NewBatchProcessor(reader io.Reader) *BatchProcessor {
	return &BatchProcessor{
		reader: reader,
	}
}

func NewFileBatchProcessor(filename string) (*BatchProcessor, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, apperrors.NewValidationError("filename", filename, err)
	}

	return &BatchProcessor{
		reader: file,
	}, nil
}

func (bp *BatchProcessor) Process() ([]BatchResult, error) {
	var results []BatchResult
	scanner := bufio.NewScanner(bp.reader)
	lineNumber := 0

	for scanner.Scan() {
		lineNumber++
		text := scanner.Text()

		if text == "" {
			continue
		}

		slug, err := Generate(text)
		result := BatchResult{
			LineNumber: lineNumber,
			Original:   text,
			Slug:       slug,
			Error:      err,
		}

		results = append(results, result)
	}

	if err := scanner.Err(); err != nil {
		return results, err
	}

	return results, nil
}

func (bp *BatchProcessor) ProcessWithCallback(callback func(BatchResult)) error {
	scanner := bufio.NewScanner(bp.reader)
	lineNumber := 0

	for scanner.Scan() {
		lineNumber++
		text := scanner.Text()

		if text == "" {
			continue
		}

		slug, err := Generate(text)
		result := BatchResult{
			LineNumber: lineNumber,
			Original:   text,
			Slug:       slug,
			Error:      err,
		}

		callback(result)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
