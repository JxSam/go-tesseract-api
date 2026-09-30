package tesseract

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/otiai10/gosseract/v2"
)

func GetText(pdfPath string) (*string, error) {
	// Временная директория для изображений страниц
	tempDir, err := os.MkdirTemp("", "pdf-ocr-*")
	if err != nil {
		return nil, fmt.Errorf("create temp dir: %w", err)
	}
	defer os.RemoveAll(tempDir)

	// PDF -> PNG
	outputPrefix := filepath.Join(tempDir, "page")

	cmd := exec.Command(
		"pdftoppm",
		"-png",
		"-r", "300",
		pdfPath,
		outputPrefix,
	)

	if output, err := cmd.CombinedOutput(); err != nil {
		return nil, fmt.Errorf(
			"convert pdf to images: %w: %s",
			err,
			strings.TrimSpace(string(output)),
		)
	}

	// Получаем список страниц
	files, err := filepath.Glob(filepath.Join(tempDir, "page-*.png"))
	if err != nil {
		return nil, fmt.Errorf("find pages: %w", err)
	}

	if len(files) == 0 {
		return nil, fmt.Errorf("pdf contains no pages")
	}

	// Tesseract
	client := gosseract.NewClient()
	defer client.Close()

	client.SetLanguage("rus", "eng")

	var allText strings.Builder

	for _, file := range files {
		if err := client.SetImage(file); err != nil {
			return nil, fmt.Errorf(
				"set image %s: %w",
				file,
				err,
			)
		}

		text, err := client.Text()
		if err != nil {
			return nil, fmt.Errorf(
				"tesseract page %s: %w",
				file,
				err,
			)
		}

		allText.WriteString(text)
		allText.WriteString(" ")
	}

	// Объединенный текст
	cleanText := allText.String()

	// \n, \r\n -> пробел
	reNewLines := regexp.MustCompile(`[\r\n]+`)
	cleanText = reNewLines.ReplaceAllString(cleanText, " ")

	// Несколько пробелов -> один
	reSpaces := regexp.MustCompile(`\s+`)
	cleanText = reSpaces.ReplaceAllString(cleanText, " ")

	// Убираем пробелы в начале и конце
	cleanText = strings.TrimSpace(cleanText)

	return &cleanText, nil
}
