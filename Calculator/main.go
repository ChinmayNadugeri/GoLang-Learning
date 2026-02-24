package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"unicode"
)

func main() {

	// Serve CSS
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Serve HTML
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./templates/index.html")
	})

	// API endpoint
	http.HandleFunc("/calculate", calculateHandler)

	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

type Request struct {
	Expression string `json:"expression"`
}

type Response struct {
	Result float64 `json:"result,omitempty"`
	Error  string  `json:"error,omitempty"`
}

func calculateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		json.NewEncoder(w).Encode(Response{Error: "Invalid request"})
		return
	}

	result, err := evaluateExpression(req.Expression)
	if err != nil {
		json.NewEncoder(w).Encode(Response{Error: err.Error()})
		return
	}

	json.NewEncoder(w).Encode(Response{Result: result})
}

func evaluateExpression(expr string) (float64, error) {
	var numbers []float64
	var operators []rune

	i := 0
	for i < len(expr) {
		ch := rune(expr[i])

		if ch == '=' {
			return 0, errors.New("invalid expression: '=' sign encountered")
		}

		if unicode.IsSpace(ch) {
			i++
			continue
		}

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

	n2 := (*numbers)[len(*numbers)-1]
	n1 := (*numbers)[len(*numbers)-2]
	*numbers = (*numbers)[:len(*numbers)-2]

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
