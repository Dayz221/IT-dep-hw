package calculator

import "fmt"

type Shape interface {
	GetArea() (float64, error)
}

func TotalArea(shapes []Shape) (float64, error) {
	var sum float64 = 0
	for _, el := range shapes {
		curArea, err := el.GetArea()
		if err != nil {
			return 0, fmt.Errorf("Ошибка при вычислении площади фигуры: %s", err)
		}
		sum += curArea
	}
	return sum, nil
}
