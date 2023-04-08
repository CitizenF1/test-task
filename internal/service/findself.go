package service

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func FindSelf(search string) ([]string, error) {
	// Получаем список всех файлов и папок в текущей директории и её поддиректориях
	var identifiers []string

	// рекурсивно обходим все папки с файлами
	err := filepath.Walk("./", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// проверяем только go-файлы и файлы без префикса test
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".go") && !strings.Contains(info.Name(), "_test") {
			content, err := os.ReadFile(path)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Failed to read file %s: %v\n", path, err)
				return nil
			}
			// регулярное выражение для поиска идентификаторов
			re := regexp.MustCompile(`\b\w+\b`)
			for _, match := range re.FindAllString(string(content), -1) {
				if strings.Contains(match, search) {
					identifiers = append(identifiers, match)
				}
			}
		}
		return nil
	})
	return identifiers, err
}
