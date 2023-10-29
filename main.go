package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calculate(operation string) (string, error) {
	romanNumerals := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}

	items := strings.Split(operation, " ")
	if len(items) != 3 {
		return "", errors.New("Вывод ошибки, так как строка не является математической операцией. ")
	}

	operand1 := items[0]
	operator := items[1]
	operand2 := items[2]

	if (contains(romanNumerals, operand1) && !contains(romanNumerals, operand2)) || (!contains(romanNumerals, operand1) && contains(romanNumerals, operand2)) {
		return "", errors.New("Вывод ошибки, так как используются одновременно разные системы счисления. ")
	}

	var num1, num2 int
	var err error

	if contains(romanNumerals, operand1) {
		num1, err = romanToArabic(operand1)
		if err != nil {
			return "", err
		}
		num2, err = romanToArabic(operand2)
		if err != nil {
			return "", err
		}
	} else {
		num1, err = strconv.Atoi(operand1)

		if err != nil {
			return "", errors.New("Вывод ошибки, так как строка не является математической операцией. ")
		}
		num2, err = strconv.Atoi(operand2)
		if err != nil {
			return "", errors.New("Неподходящее условие ")
		}
		if num1 <= 0 || num1 > 10 || num2 <= 0 || num2 > 10 {
			return "", errors.New("Неподходящее условие ")
		}
	}

	var result int
	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			return "", errors.New("Деление на ноль ")
		}
		result = num1 / num2
	}

	resultStr := ""
	if contains(romanNumerals, operand1) {
		resultStr = arabicToRoman(result)

		if resultStr == "" {
			return "", errors.New("Вывод ошибки, так как в римской системе нет отрицательных чисел. ")
		}
	} else {
		resultStr = strconv.Itoa(result)
	}

	return resultStr, nil
}

func contains(arr []string, item string) bool {
	for _, val := range arr {
		if val == item {
			return true
		}
	}
	return false
}

func romanToArabic(roman string) (int, error) {
	romanNumerals := map[string]int{
		"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
	}

	if val, ok := romanNumerals[roman]; ok {
		return val, nil
	} else {
		return 0, errors.New("Неверное число ")
	}
}

func arabicToRoman(arabic int) string {
	arabicNumerals := []int{100, 90, 50, 40, 10, 9, 5, 4, 1}
	romanNumerals := []string{"C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	result := ""
	for i, val := range arabicNumerals {
		for arabic >= val {
			result += romanNumerals[i]
			arabic -= val
		}
	}

	return result
}

func main() {
	fmt.Print("Введите арифметическую операцию (например, 1 + 2): ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	result, err := calculate(input)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Результат:", result)
	}
}
