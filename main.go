package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	romanToInt := map[string]int{
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

	intToRoman := map[int]string{
		1:  "I",
		2:  "II",
		3:  "III",
		4:  "IV",
		5:  "V",
		6:  "VI",
		7:  "VII",
		8:  "VIII",
		9:  "IX",
		10: "X",
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Привет, введи математическую операцию:")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	parts := strings.Fields(input)

	if len(parts) != 3 {
		panic("Неверный формат ввода")
	}

	isRoman := func(s string) bool {
		_, exists := romanToInt[s]
		return exists
	}

	var operand1 int
	operator := parts[1]
	var operand2 int
	var isRim bool

	if isRoman(parts[0]) && isRoman(parts[2]) {
		isRim = true
		operand1 = romanToInt[parts[0]]
		operand2 = romanToInt[parts[2]]
	} else if !isRoman(parts[0]) && !isRoman(parts[2]) {
		var err error
		operand1, err = strconv.Atoi(parts[0])
		if err != nil {
			panic("Калькулятор принимает только целые числа")
		}
		operand2, err = strconv.Atoi(parts[2])
		if err != nil {
			panic("Калькулятор принимает только целые числа")
		}
	} else {
		panic("Калькулятор принимает только целые числа и римские числа")
	}

	if operand1 > 10 || operand2 > 10 {
		panic("Калькулятор не может работать с числами больше 10")
	}

	var result int
	if operator == "+" {
		result = operand1 + operand2
	} else if operator == "-" {
		result = operand1 - operand2
	} else if operator == "*" {
		result = operand1 * operand2
	} else if operator == "/" {
		if operand1 == 0 || operand2 == 0 {
			panic("Ошибка, делить на 0 нельзя")
		}
		result = operand1 / operand2
	} else {
		panic("Калькулятор принимает следующие математические операции: +, -, *, /")
	}

	if isRim {
		if result < 1 {
			panic("Невалидный результат для римских чисел, число меньше единицы")
		}

		if result > 10 {
			panic("Калькулятор не может работать с числами больше 10")
		}
		romanResult, exists := intToRoman[result]
		if !exists {
			panic("Результат не может быть представлен римским числом")
		}
		fmt.Println(romanResult)
	} else {
		fmt.Print(result)
	}
}
