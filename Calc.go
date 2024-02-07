package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Введите 'выход' или 'exit' для завершения): ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Попробуйте ввести снова:", err)
			os.Exit(1)
		}

		input = strings.TrimSpace(input)

		if input == "выход" || input == "exit" {
			fmt.Println("Программа завершена.")
			break
		}

		result, err := evaluateExpression(input)
		if err != nil {
			fmt.Println("Ошибка вычисления:", err)
		} else {
			fmt.Println("Результат:", result)
		}
	}
}

func evaluateExpression(input string) (interface{}, error) {
	expression := strings.Fields(input)
	if len(expression) == 1 {
		panic(fmt.Sprintf("Введите математическое выражение"))
	}
	if len(expression) <= 2 {
		panic(fmt.Sprintf("Не является математической операцией"))
	}
	if len(expression) > 3 {
		panic(fmt.Sprintf("формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)."))
	}

	num1, err1 := parseNum(expression[0])
	if err1 != nil {
		panic(fmt.Sprintf("Ошибка ввода числа 1: %v", err1))
	}

	operator := expression[1]
	if operator != "+" && operator != "-" && operator != "*" && operator != "/" {
		panic(fmt.Sprintf("Неверный оператор: %s", operator))
	}

	num2, err2 := parseNum(expression[2])
	if err2 != nil {
		panic(fmt.Sprintf("Ошибка ввода числа 2: %v", err2))
	}

	result := calculate(num1, operator, num2)
	return result, nil
}

func calculate(num1 interface{}, operator string, num2 interface{}) interface{} {
	switch operator {
	case "+":
		return Add(num1, num2)
	case "-":
		return Subtract(num1, num2)
	case "*":
		return Multiply(num1, num2)
	case "/":
		return Divide(num1, num2)
	default:
		return nil
	}
}

func Add(a interface{}, b interface{}) interface{} {
	switch a := a.(type) {
	case int:
		switch b := b.(type) {
		case int:
			return a + b
		case string:
			panic(fmt.Sprintf("Используются разные системы счисления"))
		default:
			return nil
		}
	case string:
		switch b := b.(type) {
		case int:
			panic(fmt.Sprintf("Используются разные системы счисления"))
		case string:
			return AddRomanNum(a, b)
		default:
			return nil
		}
	default:
		return nil
	}
}

func Subtract(a interface{}, b interface{}) interface{} {
	switch a := a.(type) {
	case int:
		switch b := b.(type) {
		case int:
			return a - b
		case string:
			panic(fmt.Sprintf("Используются разные системы счисления"))
		default:
			return nil
		}
	case string:
		switch b := b.(type) {
		case int:
			panic(fmt.Sprintf("Используются разные системы счисления"))
		case string:
			return SubtractRomanNum(a, b)
		default:
			return nil
		}
	default:
		return nil
	}
}

func Multiply(a interface{}, b interface{}) interface{} {
	switch a := a.(type) {
	case int:
		switch b := b.(type) {
		case int:
			return a * b
		case string:
			panic(fmt.Sprintf("Используются разные системы счисления"))
		default:
			return nil
		}
	case string:
		switch b := b.(type) {
		case int:
			panic(fmt.Sprintf("Используются разные системы счисления"))
		case string:
			return MultiplyRomanNum(a, b)
		default:
			return nil
		}
	default:
		return nil
	}
}

func Divide(a interface{}, b interface{}) interface{} {
	switch a := a.(type) {
	case int:
		switch b := b.(type) {
		case int:
			return a / b
		case string:
			panic(fmt.Sprintf("Используются разные системы счисления"))
		default:
			return nil
		}
	case string:
		switch b := b.(type) {
		case int:
			panic(fmt.Sprintf("Используются разные системы счисления"))
		case string:
			return DivideRomanNum(a, b)
		default:
			return nil
		}
	default:
		return nil
	}
}

func parseNum(input string) (interface{}, error) {
	if isNumRoman(input) {
		return strings.ToUpper(input), nil
	}

	num, err := strconv.Atoi(input)
	if num > 10 {
		panic(fmt.Sprintf("Число большое"))
	}
	if err != nil {
		return nil, err
	}
	return num, nil
}

func isNumRoman(s string) bool {
	match, _ := regexp.MatchString(`^(I|II|III|IV|V|VI|VII|VIII|IX|X)$`, strings.ToUpper(s))
	return match
}

func romanToArab(s string) int {
	RomanNum := map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
	}

	val, found := RomanNum[s]
	if !found {
		return 0
	}
	return val
}

func AddRomanNum(a, b string) string {
	num1 := romanToArab(a)
	num2 := romanToArab(b)
	result := num1 + num2
	return arabicToRoman(result)
}
func SubtractRomanNum(a, b string) string {
	num1 := romanToArab(a)
	num2 := romanToArab(b)
	result := num1 - num2
	return arabicToRoman(result)
}

func MultiplyRomanNum(a, b string) string {
	num1 := romanToArab(a)
	num2 := romanToArab(b)
	result := num1 * num2
	return arabicToRoman(result)
}

func DivideRomanNum(a, b string) string {
	num1 := romanToArab(a)
	num2 := romanToArab(b)
	result := num1 / num2
	return arabicToRoman(result)
}

func arabicToRoman(num int) string {
	if num <= 0 || num > 100 {
		panic(fmt.Sprintf("Недопустимое значение"))
	}

	romanNum := ""

	romanSymbols := []string{"C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	values := []int{100, 90, 50, 40, 10, 9, 5, 4, 1}

	for i := 0; i < len(values); i++ {
		for num >= values[i] {
			romanNum += romanSymbols[i]
			num -= values[i]
		}
	}
	return romanNum
}
