package main

import (
	"errors"
	"fmt"
)

// Функция hello
func hello(name string) string {
	return "Привет, " + name + "!"
}

// Функция printEven
func printEven(a, b int64) error {
	if a > b {
		return errors.New("левая граница больше правой")
	}
	for i := a; i <= b; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	return nil
}

// Функция apply
func apply(a, b float64, operator string) (float64, error) {
	switch operator {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errors.New("деление на ноль")
		}
		return a / b, nil
	default:
		return 0, errors.New("действие не поддерживается")
	}
}

func main() {
	// Тестирование функции hello
	fmt.Println(hello("Мир"))

	// Тестирование функции printEven
	err := printEven(1, 10)
	if err != nil {
		fmt.Println("Ошибка:", err)
	}

	err = printEven(10, 1)
	if err != nil {
		fmt.Println("Ошибка:", err)
	}

	// Тестирование функции apply
	result, err := apply(3, 5, "+")
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Printf("Результат: %f\n", result)
	}

	result, err = apply(7, 10, "*")
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Printf("Результат: %f\n", result)
	}

	result, err = apply(3, 5, "#")
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Printf("Результат: %f\n", result)
	}
}
