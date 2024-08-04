package main

import (
	"log"
	selector "uniqueizer/pkg/selector"
	uniqueness "uniqueizer/pkg/uniqueizer"

	"github.com/fatih/color"
)

func main() {
	currentBaseFolder := selector.SelectFolder(color.GreenString("Выберите папку основной существующей базы"))
	newBaseFolder := selector.SelectFolder(color.GreenString("Выберите папку новой базы, которую нужно сравнивать"))

	if currentBaseFolder == newBaseFolder {
		log.Fatalf(color.RedString("Вы не можете сравнивать одну и ту же базу!"))
	}

	uniqueness.StringUniqueness(currentBaseFolder, newBaseFolder)
}
