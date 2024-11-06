package myshapes

import (
	"errors"
	"math"
)

const Pi = 3.141592

type Circle struct {
	Radius float64
}

func (circle Circle) GetArea() (float64, error) {
	if circle.Radius < 0 {
		return 0, errors.New("радиус круга не может быть меньше 0")
	}
	return Pi * math.Pow(circle.Radius, 2), nil
}

type Rect struct {
	Width  float64
	Height float64
}

func (rect Rect) GetArea() (float64, error) {
	if rect.Width < 0 || rect.Height < 0 {
		return 0, errors.New("Ширина и высота не могут быть меньше 0")
	}
	return rect.Width * rect.Height, nil
}
