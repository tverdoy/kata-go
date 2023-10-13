package main

import (
	"errors"
	"fmt"
	"strings"
)

var supportOperations = []string{"+", "-", "/", "*"}

func main() {
	parse("1 + 2 + 3")
}

func parse(input string) {
	firstArguments, secondArgument, err := findArguments(input)

	if err != nil {
		panic(err)
	}

	fmt.Println(firstArguments, secondArgument)
}

func findArguments(input string) (string, string, error) {
	input = strings.Replace(input, " ", "", -1)
	for _, operation := range supportOperations {
		splitInput := strings.Split(input, operation)
		if len(splitInput) == 2 {
			return splitInput[0], splitInput[1], nil
		}
	}

	return "", "", errors.New("Syntax error")
}
