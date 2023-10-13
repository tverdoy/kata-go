package main

import (
	"errors"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

var supportOperations = []string{"+", "-", "/", "*"}
var supportRomans = []string{"I", "V", "X", "L", "C", "D", "M"}

func main() {
	//reader := bufio.NewReader(os.Stdin)
	//input, _ := reader.ReadString('\n')
	//input = strings.ReplaceAll(input, "\n", "")

	parse("I + II")
}

func parse(input string) {
	firstArgumentsStr, secondArgumentStr, operation, err := findArguments(input)
	if err != nil {
		panic(err)
	}

	if isRoman(firstArgumentsStr) {
		if isRoman(secondArgumentStr) {
			// first and second is roman

		} else {
			panic("second arguments is not roman")
		}
	}

	if isRoman(secondArgumentStr) {
		panic("first arguments is not arabic")
	}

	firstArguments, err := strconv.Atoi(firstArgumentsStr)
	if err != nil {
		panic(err)
	}

	secondArgument, err := strconv.Atoi(secondArgumentStr)
	if err != nil {
		panic(err)
	}

	result, err := calculate(firstArguments, secondArgument, operation)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
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

func isRoman(symbol string) bool {
	return slices.Contains(supportRomans, symbol)
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

func integerToRoman(number int) string {
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
