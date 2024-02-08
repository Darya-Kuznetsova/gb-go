package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	fmt.Println("Введите предложение")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	// Вставьте ваш код здесь

	text = strings.ToLower(text)
	fmt.Println(text)
	letterCounter := make(map[rune]int)
	var length int
	for _, v := range text {
		if unicode.IsLetter(v) {
			letterCounter[v]++
			length++
		}
	}

	keySlice := make([]rune, 0)
	valueSlice := make([]int, 0)
	percentSlice := make([]float64, 0)

	for key, v := range letterCounter {
		percentSlice = append(percentSlice, float64(v)/float64(length))
		keySlice = append(keySlice, key)
		valueSlice = append(valueSlice, v)
	}

	for i := 0; i < len(letterCounter); i++ {
		fmt.Printf("%c: %v %0.2f \n", keySlice[i], valueSlice[i], percentSlice[i])
	}
}
