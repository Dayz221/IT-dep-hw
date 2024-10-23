package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func ReadFile(fileName string) (string, error) {
	buf, err := os.ReadFile(fileName)
	if err != nil {
		return "", fmt.Errorf("Ошибка при чтении файла: %s", err)
	}

	return string(buf), nil
}

func proccessText(text string) []string {
	text = strings.ReplaceAll(text, "\r", "")
	textLines := strings.Split(text, "\n")

	linesLen := len(textLines)
	if linesLen > 1000 {
		linesLen /= 10
	} else {
		linesLen = 100
	}
	myMap := make(map[string]int, linesLen)

	for _, line := range textLines {
		myMap[line]++
	}

	uniqLen := 0
	for _, cnt := range myMap {
		if cnt == 1 {
			uniqLen++
		}
	}

	uniqTextLines := make([]string, uniqLen)
	for line, cnt := range myMap {
		if cnt == 1 {
			uniqConvertedLine := fmt.Sprintf("%s - %d байт\n", strings.ToUpper(line), len(line))
			uniqTextLines = append(uniqTextLines, uniqConvertedLine)

		}
	}

	slices.Sort(uniqTextLines)

	return uniqTextLines
}

func writeFile(uniqLines []string) error {
	file, err := os.Create("output.txt")
	if err != nil {
		return fmt.Errorf("Ошибка при записи файла: %s", err)
	}
	defer file.Close()

	for _, line := range uniqLines {
		if len(line) != 0 {
			file.WriteString(line)
		}
	}

	return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Передайте в команду запуска путь до входного файла...")
		return
	}

	myText, err := ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	uniqLines := proccessText(myText)

	err = writeFile(uniqLines)
	if err != nil {
		fmt.Println(err)
	}
}
