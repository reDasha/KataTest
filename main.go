package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func toRoman(x int) (string, error) {
	if x <= 0 {
		return "", errors.New("ошибка: в римской системе нет отрицательных чисел")
	}
	conversions := []struct {
		Value int
		Latin string
	}{
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}
	romanNumber := ""
	for _, conversion := range conversions {
		for x >= conversion.Value {
			romanNumber += conversion.Latin
			x -= conversion.Value
		}
	}
	return romanNumber, nil
}

func fromRoman(y string) int {
	romanNumbers := map[string]int{
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

	digitNumber := 0
	for romanNumber := range romanNumbers {
		if romanNumber == y {
			digitNumber = romanNumbers[romanNumber]
		}
	}
	if digitNumber == 0 {
		panic("ошибка: введены некорректные данные")
	}
	return digitNumber
}

func arithmetic(a, b int, c string) int {
	switch c {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	}
	return arithmetic(a, b, c)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		text, err := reader.ReadString('\n')
		if err != nil {
			panic("ошибка: введены некорректные данные")
		}
		operation := strings.Split(strings.TrimSpace(text), " ")

		if len(operation) != 3 || (operation[1] != "+" && operation[1] != "-" && operation[1] != "*" && operation[1] != "/") {
			panic("ошибка: введены некорректные данные")
		}

		x, _ := strconv.Atoi(operation[0])
		y, _ := strconv.Atoi(operation[2])

		c := operation[1]

		unitextA := []rune(operation[0])
		unitextB := []rune(operation[2])

		if x > 0 && x <= 10 && y > 0 && y <= 10 {
			a := x
			b := y
			fmt.Println(arithmetic(a, b, c))
		} else if unicode.IsLetter(unitextA[0]) && unicode.IsLetter(unitextB[0]) {
			a := fromRoman(operation[0])
			b := fromRoman(operation[2])
			res, err := toRoman(arithmetic(a, b, c))
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(res)
			}
		} else {
			panic("ошибка: введены некорректные данные")
		}
	}
}
