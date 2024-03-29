package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Укажите полный путь до файла вторым аргументом")

	}

	filePth := os.Args[1]

	var fileName, fileExt string

	// Напишите код, который выведет следующее
	// filename: <name>
	// extension: <extension>

	// Подсказка. Возможно вам понадобится функция strings.LastIndex
	// Для проверки своего решения используйте функции filepath.Base() filepath.Ext(
	// ) Они могут помочь для проверки решения

	dotIndex := strings.LastIndex(filePth, ".")
	fileExt = filePth[dotIndex+1:]
	solidusIndex := strings.LastIndex(filePth, "/")
	fileName = filePth[solidusIndex+1 : dotIndex]

	fmt.Printf("filename: %s\n", fileName)
	fmt.Printf("extension: %s\n", fileExt)
}
