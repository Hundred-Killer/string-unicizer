package selector

import (
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
)

func SelectFolder(label string) string {
	folders := []string{}

	entries, err := os.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		if e.IsDir() {
			folders = append(folders, e.Name())
		}
	}

	if len(folders) < 2 {
		log.Fatalf(color.RedString("Создайте папку с существующей базой и базу, которую нужно проверить."))
	}

	prompt := promptui.Select{
		Label: label,
		Items: folders,
	}

	_, result, err := prompt.Run()

	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}

	return result
}
