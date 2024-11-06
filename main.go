package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Dayz221/IT-dep-hw/calculator"
	"github.com/Dayz221/IT-dep-hw/myshapes"
)

func ReadFile(name string) (string, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		return "", fmt.Errorf("возникла ошибка при чтении файла %s: %w", name, err)
	}
	strdata := string(data)
	strdata = strings.ReplaceAll(strdata, "\r", "")
	strdata = strings.Trim(strdata, "\n")
	return strdata, nil
}

func SplitShapes(data string) ([]calculator.Shape, error) {
	splittedLines := strings.Split(data, "\n")
	shapeList := make([]calculator.Shape, len(splittedLines))

	for i, el := range splittedLines {
		if strings.HasPrefix(el, "Rect") {
			params := strings.Split(el, " ")
			if len(params) != 3 {
				return nil, fmt.Errorf("введено неверное количество параметров для Rect: %s", el)
			}

			width, err := strconv.ParseFloat(params[1], 64)

			if err != nil {
				return nil, fmt.Errorf("ошибка в параметрах фигуры: %s", el)
			}

			height, err := strconv.ParseFloat(params[2], 64)

			if err != nil {
				return nil, fmt.Errorf("ошибка в параметрах фигуры: %s", el)
			}

			if width < 0 || height < 0 {
				return nil, fmt.Errorf("ошибка в параметрах фигуры: высота и ширина не могут быть меньше 0: %s", el)
			}

			shapeList[i] = myshapes.Rect{Width: width, Height: height}

		} else if strings.HasPrefix(el, "Circle") {
			params := strings.Split(el, " ")
			if len(params) != 2 {
				return nil, fmt.Errorf("введено неверное количество параметров для Circle: %s", el)
			}

			radius, err := strconv.ParseFloat(params[1], 64)

			if err != nil {
				return nil, fmt.Errorf("ошибка в параметрах фигуры: %s", el)
			}

			if radius < 0 {
				return nil, fmt.Errorf("ошибка в параметрах фигуры: радиус не может быть меньше 0: %s", el)
			}

			shapeList[i] = myshapes.Circle{Radius: radius}

		} else {
			return nil, fmt.Errorf("неправильно задана фигура: %s", el)
		}
	}

	return shapeList, nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Введите название файла с фигурами в строку запуска!")
		return
	}

	fileData, err := ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	shapes, err := SplitShapes(fileData)
	if err != nil {
		fmt.Println(err)
		return
	}

	totalArea, err := calculator.TotalArea(shapes)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Суммарная площадь фигур: %f", totalArea)
}
