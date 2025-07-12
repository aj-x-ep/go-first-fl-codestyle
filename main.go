// Package main - используемый пакет
package main

import (
	"fmt"
	"math"
)

// Message - Приветствие не меняется, поэтому это константа. Замена одинарных кавычек на двойные.
const Message = "Добро пожаловать в самую лучшую программу для вычисления квадратного корня из заданного числа"

// CalculateSquareRoot - экспортируемая функция, вычисляет квадратный корень.
func CalculateSquareRoot(myNumber float64) float64 {
	return math.Sqrt(myNumber)
}

// Calc - проверяет чтобы число было больше нуля
func Calc(number float64) {
	if number < 0 {
		return
	}
	fmt.Println("Мы вычислили квадратный корень из введённого вами числа. Это будет:", CalculateSquareRoot(number))
	return
}

func main() {
	fmt.Println(Message)
	Calc(25.5)
}
