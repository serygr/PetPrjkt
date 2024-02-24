package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите текст: ")
	input, err := reader.ReadString('\n')

	input = strings.TrimSpace(input)

	if err != nil {
		fmt.Println("Ошибка", err)
		return
	}
	varibl := strings.Fields(input)

	var sep string
	varibl1 := varibl[0]
	oper := varibl[1]
	varibl2 := varibl[2]

	if _, err := strconv.Atoi(varibl1); err == nil {
		panic(fmt.Sprintf("Первое значение не строка"))
	}
	if _, err := strconv.ParseFloat(varibl1, 64); err == nil {
		panic(fmt.Sprintf("Первое значение не строка"))
	}

	if len(varibl) > 3 {
		for _, val := range varibl {
			if val == "+" || val == "-" || val == "*" || val == "/" {
				sep = val
			}
		}
		joinedString := strings.Join(varibl, " ")
		res := strings.Split(joinedString, sep)
		varibl1 = strings.Replace(res[0], "\"", "", -1)
		varibl2 = strings.Replace(res[1], "\"", "", -1)
		oper = sep
	}
	if len(strings.Replace(varibl1, "\"", "", -1)) > 10 {
		panic(fmt.Sprintf("Слишком много символов"))
	}
	if len(strings.Replace(varibl2, "\"", "", -1)) > 10 {
		panic(fmt.Sprintf("Слишком много символов"))
	}

	result := conc(varibl1, oper, varibl2)

	if len(result) > 40 {
		fmt.Println("Вывод:", result[:40]+"...")
	} else {
		fmt.Println("Вывод:", result)
	}

}
func conc(varibl1 string, oper string, varibl2 string) string {
	switch oper {
	case "+":
		return "\"" + strings.Replace(varibl1, "\"", "", -1) + strings.Replace(varibl2, "\"", "", -1) + "\""
	case "-":
		return strings.Replace(varibl1, varibl2, "", -1)
	case "*":
		num, err := strconv.Atoi(varibl2)
		if err != nil {
			fmt.Println("Ошибка умножения:", err)
		}
		if num > 10 {
			panic(fmt.Sprintf("Число большое"))
		}
		return "\"" + strings.Repeat(strings.Replace(varibl1, "\"", "", -1), num) + "\""
	case "/":
		num, err := strconv.Atoi(strings.TrimSpace(varibl2))
		if err != nil {
			fmt.Println("Ошибка деления:", err)
		}
		if num > 10 {
			panic(fmt.Sprintf("Число большое"))
		}
		divis := len(strings.Replace(varibl1, "\"", "", -1))/num + 1
		return varibl1[:divis] + "\""
	default:
		panic(fmt.Sprintf("Что то пошло не так", oper))
		return ""
	}
}
