package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var supportOperations = []string{"+", "-", "/", "*"}
var supportRomans = []string{"I", "V", "X", "L", "C", "D", "M"}

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.ReplaceAll(input, "\n", "")

	parse(input)
}

func parse(input string) {
	firstArgumentsStr, secondArgumentStr, operation, err := findArguments(input)
	if err != nil {
		panic(err)
	}

	firstArgument, secondArgument, isRoman := parseArguments(firstArgumentsStr, secondArgumentStr)

	if firstArgument < 1 || firstArgument > 10 || secondArgument < 1 || secondArgument > 10 {
		panic("numbers greater than 10 or less than 1")
	}

	result, err := calculate(firstArgument, secondArgument, operation)
	if err != nil {
		panic(err)
	}

	if isRoman {
		if result < 1 {
			panic("result cannot be negative in roman")
		}
		fmt.Println(intToRoman(result))
	} else {
		fmt.Println(result)
	}
}

func calculate(a int, b int, operation string) (int, error) {
	switch operation {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "/":
		return a / b, nil
	case "*":
		return a * b, nil
	default:
		return 0, errors.New("not support operation")
	}
}

func findArguments(input string) (string, string, string, error) {
	input = strings.ReplaceAll(input, " ", "")
	for _, operation := range supportOperations {
		splitInput := strings.Split(input, operation)
		if len(splitInput) == 2 {
			return splitInput[0], splitInput[1], operation, nil
		}
	}

	return "", "", "", errors.New("syntax error")
}

func parseArguments(first string, second string) (int, int, bool) {
	if isRoman(first) {
		if isRoman(second) {
			return romanToInt(first), romanToInt(second), true
		} else {
			panic("second arguments is not roman")
		}
	}

	if isRoman(second) {
		panic("first arguments is not arabic")
	}

	firstArguments, err := strconv.Atoi(first)
	if err != nil {
		panic(err)
	}

	secondArgument, err := strconv.Atoi(second)
	if err != nil {
		panic(err)
	}

	return firstArguments, secondArgument, false
}

func isRoman(symbol string) bool {
	for _, roman := range supportRomans {
		if strings.Contains(symbol, roman) {
			return true
		}
	}

	return false
}

func romanToInt(s string) int {
	romans := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	s = strings.Replace(s, "CM", "CCCCCCCCC", -1) // 900
	s = strings.Replace(s, "CD", "CCCC", -1)      // 400
	s = strings.Replace(s, "XC", "XXXXXXXXX", -1) // 90
	s = strings.Replace(s, "XL", "XXXX", -1)      // 40
	s = strings.Replace(s, "IX", "IIIIIIIII", -1) // 9
	s = strings.Replace(s, "IV", "IIII", -1)      // 4

	var sum int
	for _, roman := range s {
		sum += romans[roman]
	}
	return sum
}

func intToRoman(number int) string {
	maxRomanNumber := 3999
	if number > maxRomanNumber {
		return strconv.Itoa(number)
	}

	conversions := []struct {
		value int
		digit string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
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

	var roman strings.Builder
	for _, conversion := range conversions {
		for number >= conversion.value {
			roman.WriteString(conversion.digit)
			number -= conversion.value
		}
	}

	return roman.String()
}
