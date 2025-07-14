// Package mycodestyle содержит функции для стиля кода и вычислений.
package mycodestyle

import (
	"fmt"
	"math"
)

// Message - константа приветствия.
const Message = "Добро пожаловать в самую лучшую программу для вычисления квадратного корня из заданного числа"

// CalculateSquareRoot - экспортируемая функция, вычисляет квадратный корень.
func CalculateSquareRoot(myNumber float64) float64 {
	return math.Sqrt(myNumber)
}

// Calc - проверка, чтобы число было больше нуля и вывод результата.
func Calc(number float64) {
	if number < 0 {
		fmt.Println("Число должно быть неотрицательным")
		return
	}
	fmt.Println("Мы вычислили квадратный корень из введённого вами числа. Это будет:", CalculateSquareRoot(number))
}
