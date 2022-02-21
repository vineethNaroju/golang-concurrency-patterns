package main

import (
	"fmt"
	"reflect"
)

type Shape interface {
	Area() float64
}

type Triangle struct {
	base, height float64
}

func (t *Triangle) Area() float64 {
	return 0.5 * t.base * t.height
}

func NewTriangle(base, height float64) *Triangle {
	return &Triangle{
		base: base, height: height,
	}
}

type Square struct {
	side float64
}

func NewSquare(side float64) *Square {
	return &Square{
		side: side,
	}
}

func (s *Square) Area() float64 {
	return s.side * s.side
}

func main() {
	var shapeSlice []Shape

	sq := NewSquare(10)

	tri := NewTriangle(5, 100)

	shapeSlice = append(shapeSlice, sq)
	shapeSlice = append(shapeSlice, tri)

	for idx, val := range shapeSlice {
		fmt.Println(idx, reflect.TypeOf(val), val.Area())

		switch val.(type) {
		case *Square:
			fmt.Println("Square")
		// case *Triangle:
		// 	fmt.Println("Triangle")
		default:
			fmt.Println("weird shape")
		}

	}
}
