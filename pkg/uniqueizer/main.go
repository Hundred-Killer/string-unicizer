package uniqueizer

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"time"

	"github.com/fatih/color"
	progressbar "github.com/schollz/progressbar/v3"
)

func StringUniqueness(currentBaseFolder string, newBaseFolder string) {
	currentBaseFiles, _ := filepath.Glob(fmt.Sprintf("./%s/*.txt", currentBaseFolder))
	newBaseFiles, _ := filepath.Glob(fmt.Sprintf("./%s/*.txt", newBaseFolder))

	currentBaseArray := []string{}
	newBaseArray := []string{}
	uniqueArray := []string{}

	notUnique := 0
	unique := 0

	for _, newBaseFile := range newBaseFiles {
		file, scanner := openFile(newBaseFile)
		defer file.Close()

		for scanner.Scan() {
			newBaseArray = append(newBaseArray, strings.Split(scanner.Text(), "\n")...)
		}
	}

	for _, currentBaseFile := range currentBaseFiles {
		file, scanner := openFile(currentBaseFile)
		defer file.Close()

		for scanner.Scan() {
			currentBaseArray = append(currentBaseArray, strings.Split(scanner.Text(), "\n")...)
		}
	}

	if _, err := os.Stat("./output"); os.IsNotExist(err) {
		os.Mkdir("./output", os.ModePerm)
	}

	progressBar := progressbar.Default(int64(len(newBaseArray)))

	for _, newBaseString := range newBaseArray {
		progressBar.Add(1)

		if slices.Contains(currentBaseArray, newBaseString) {
			notUnique++
			continue
		}

		unique++
		uniqueArray = append(uniqueArray, newBaseString)
	}

	os.WriteFile("./output/"+time.Now().Format(time.RFC3339), []byte(strings.Join(uniqueArray, "\n")), os.ModePerm)

	fmt.Println(color.GreenString(fmt.Sprintf("Уникализация успешно завершена. Уникальные строки: %d. Повторяющиеся строки: %d", unique, notUnique)))
}

func openFile(filePath string) (*os.File, *bufio.Scanner) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf(color.RedString(fmt.Sprintf("%d", err)))
	}

	scanner := bufio.NewScanner(file)

	return file, scanner
}
