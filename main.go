package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"strings"

	"github.com/chai2010/webp"
	webpDecode "golang.org/x/image/webp"
)

// Функция для конвертирования изображения в формат WebP
func convertToWebP(inputPath, outputPath string) error {
	// Открываем файл изображения
	file, err := os.Open(inputPath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Определяем формат изображения и декодируем
	var img image.Image
	if strings.HasSuffix(inputPath, ".jpeg") || strings.HasSuffix(inputPath, ".jpg") {
		img, err = jpeg.Decode(file)
	} else if strings.HasSuffix(inputPath, ".png") {
		img, err = png.Decode(file)
	} else {
		return fmt.Errorf("unsupported image format")
	}
	if err != nil {
		return fmt.Errorf("failed to decode image: %w", err)
	}

	// Создаем выходной файл для WebP
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer outputFile.Close()

	// Кодируем изображение в формат WebP
	if err := webp.Encode(outputFile, img, &webp.Options{Lossless: true}); err != nil {
		return fmt.Errorf("failed to encode to webp: %w", err)
	}

	return nil
}

// Функция для декодирования изображения из формата WebP
func decodeWebP(inputPath string) (image.Image, error) {
	// Открываем WebP файл
	file, err := os.Open(inputPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Декодируем изображение из WebP
	img, err := webpDecode.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("failed to decode webp: %w", err)
	}

	return img, nil
}

func main() {
	inputImagePath := "testdata/architecture.png" // Путь к входному файлу
	outputWebPPath := "testdata/output.webp"      // Путь к выходному файлу WebP
	decodedImageOutput := "decoded.png"           // Путь для сохранения декодированного изображения

	// Конвертируем изображение в WebP
	if err := convertToWebP(inputImagePath, outputWebPPath); err != nil {
		fmt.Println("Conversion to WebP failed:", err)
		return
	}
	fmt.Println("Conversion to WebP successful")

	// Декодируем изображение из WebP
	img, err := decodeWebP(outputWebPPath)
	if err != nil {
		fmt.Println("Decoding WebP failed:", err)
		return
	}
	fmt.Println("Decoding WebP successful")

	// Сохраняем декодированное изображение как PNG
	outputFile, err := os.Create(decodedImageOutput)
	if err != nil {
		fmt.Println("Failed to create decoded image file:", err)
		return
	}
	defer outputFile.Close()
	if err := png.Encode(outputFile, img); err != nil {
		fmt.Println("Failed to save decoded image as PNG:", err)
	} else {
		fmt.Println("Decoded image saved successfully as PNG")
	}
}
