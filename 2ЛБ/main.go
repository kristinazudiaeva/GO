// 1
package main

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

// formatIP принимает массив из 4 байт и возвращает строку в формате "127.0.0.1".
func formatIP(ip [4]byte) string {
	var parts []string
	for _, b := range ip {
		parts = append(parts, fmt.Sprintf("%d", b))
	}
	return strings.Join(parts, ".")
}

// listEven возвращает срез четных чисел в диапазоне [start, end] и ошибку, если start > end.
func listEven(start, end int) ([]int, error) {
	if start > end {
		return nil, errors.New("левая граница диапазона больше правой")
	}

	var evens []int
	for i := start; i <= end; i++ {
		if i%2 == 0 {
			evens = append(evens, i)
		}
	}
	return evens, nil
}

// 2
// countCharacters подсчитывает количество каждого символа в строке.
func countCharacters(input string) map[rune]int {
	charCount := make(map[rune]int)
	for _, char := range input {
		charCount[char]++
	}
	return charCount
}

// 3
type Point struct {
	X, Y float64
}

type Line struct {
	Start, End Point
}

// Length возвращает длину отрезка.
func (l Line) Length() float64 {
	dx := l.End.X - l.Start.X
	dy := l.End.Y - l.Start.Y
	return math.Sqrt(dx*dx + dy*dy)
}

type Triangle struct {
	A, B, C Point
}

// Area рассчитывает площадь треугольника по формуле Герона.
func (t Triangle) Area() float64 {
	a := Line{t.A, t.B}.Length()
	b := Line{t.B, t.C}.Length()
	c := Line{t.C, t.A}.Length()
	s := (a + b + c) / 2
	return math.Sqrt(s * (s - a) * (s - b) * (s - c))
}

type Circle struct {
	Center Point
	Radius float64
}

// Area рассчитывает площадь круга.
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Shape interface {
	Area() float64
}

// printArea выводит площадь фигуры.
func printArea(s Shape) {
	fmt.Printf("Площадь фигуры: %.2f\n", s.Area())
}

func main() {
	// Пример использования formatIP
	ip := [4]byte{127, 0, 0, 1}
	fmt.Println("IP адрес:", formatIP(ip))

	// Пример использования listEven
	evens, err := listEven(1, 10)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Четные числа:", evens)
	}

	// Пример подсчета символов
	text := "hello world"
	fmt.Println("Подсчет символов:", countCharacters(text))

	// Пример работы с фигурами
	triangle := Triangle{
		A: Point{0, 0},
		B: Point{0, 3},
		C: Point{4, 0},
	}
	circle := Circle{
		Center: Point{0, 0},
		Radius: 5,
	}

	printArea(triangle) 
	printArea(circle) 
}
