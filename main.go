package main

import (
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
)

func ReadFile(fileName string) string {
	var text string

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Не смог открыть файл...")
		return ""
	}
	defer file.Close()

	buf := make([]byte, 64)

	for {
		n, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		text += string(buf[:n])

	}

	return text
}

func main() {
	myText := ReadFile("input.txt")
	myText = strings.ReplaceAll(myText, "\r", "")

	myTextLines := strings.Split(myText, "\n")

	myMap := make(map[string]int, 10)

	for _, line := range myTextLines {
		myMap[line] += 1
	}

	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Не смог создать файл...")
		os.Exit(1)
	}
	defer file.Close()

	uniqTextLines := []string{}
	for line, cnt := range myMap {
		if cnt == 1 && len(line) != 0 {
			uniqTextLines = append(uniqTextLines, line)
		}
	}
	slices.Sort(uniqTextLines)
	slices.Reverse(uniqTextLines)

	if len(uniqTextLines) == 0 {
		file.WriteString("Не нашель уникальных строк...")
	}
	for _, line := range uniqTextLines {
		file.WriteString(strings.ToUpper(line) + " - " + strconv.Itoa(len(line)) + " байт \n")
	}
}
