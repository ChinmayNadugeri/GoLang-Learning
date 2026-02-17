package main

import (
	"errors"
	"fmt"
	"strconv"
	"unicode"
)

func main() {
	var input string
	fmt.Print("Enter expression: ")
	fmt.Scanln(&input)

	result, err := evaluateExpression(input)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Result:", result)
}

func evaluateExpression(expr string) (float64, error) {
	var numbers []float64
	var operators []rune

	i := 0
	for i < len(expr) {
		ch := rune(expr[i])

		// Terminate if '=' encountered
		if ch == '=' {
			return 0, errors.New("invalid expression: '=' sign encountered")
		}

		// Skip spaces
		if unicode.IsSpace(ch) {
			i++
			continue
		}

		// If digit, parse full number
		if unicode.IsDigit(ch) {
			start := i
			for i < len(expr) && (unicode.IsDigit(rune(expr[i])) || expr[i] == '.') {
				i++
			}
			num, err := strconv.ParseFloat(expr[start:i], 64)
			if err != nil {
				return 0, err
			}
			numbers = append(numbers, num)
			continue
		}

		// If operator
		if isOperator(ch) {
			for len(operators) > 0 &&
				precedence(operators[len(operators)-1]) >= precedence(ch) {

				err := applyOperator(&numbers, &operators)
				if err != nil {
					return 0, err
				}
			}
			operators = append(operators, ch)
		} else {
			return 0, fmt.Errorf("invalid character: %c", ch)
		}

		i++
	}

	// Apply remaining operators
	for len(operators) > 0 {
		err := applyOperator(&numbers, &operators)
		if err != nil {
			return 0, err
		}
	}

	if len(numbers) != 1 {
		return 0, errors.New("invalid expression")
	}

	return numbers[0], nil
}

func isOperator(ch rune) bool {
	return ch == '+' || ch == '-' || ch == '*' || ch == '/'
}

func precedence(op rune) int {
	switch op {
	case '+', '-':
		return 1
	case '*', '/':
		return 2
	}
	return 0
}

func applyOperator(numbers *[]float64, operators *[]rune) error {
	if len(*numbers) < 2 {
		return errors.New("invalid expression")
	}

	// Pop last two numbers
	n2 := (*numbers)[len(*numbers)-1]
	n1 := (*numbers)[len(*numbers)-2]
	*numbers = (*numbers)[:len(*numbers)-2]

	// Pop operator
	op := (*operators)[len(*operators)-1]
	*operators = (*operators)[:len(*operators)-1]

	var result float64

	switch op {
	case '+':
		result = n1 + n2
	case '-':
		result = n1 - n2
	case '*':
		result = n1 * n2
	case '/':
		if n2 == 0 {
			return errors.New("division by zero")
		}
		result = n1 / n2
	}

	*numbers = append(*numbers, result)
	return nil
}
