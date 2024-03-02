package utils

import (
	"fmt"
	"io"
	"log"
	"os"
)

func HtmlString(path string) (string, error) {
	htmlFilePath := path // "path/to/your/file.html" // 替换为实际文件路径
	file, err := os.Open(htmlFilePath)
	if err != nil {
		log.Fatalf("Error opening HTML file: %s", err)
	}
	defer file.Close()

	htmlString, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Error reading HTML file: %s", err)
		return "", err
	}

	fmt.Println(string(htmlString))
	return string(htmlString), nil
}
