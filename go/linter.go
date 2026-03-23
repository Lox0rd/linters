package main
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
type Calculation struct {
	Expression string
	Result     float64
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	history := []Calculation{}
	fmt.Println("Простой калькулятор (введите 'exit' для выхода, 'history' для просмотра истории)")
	fmt.Println("Поддерживаемые операции: +, -, *, /")
	for {
		fmt.Print("\nВведите выражение (например, 5 + 3): ")
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())
		if input == "exit" {
			fmt.Println("До свидания!")
			break
		}
		if input == "history" {
			printHistory(history)
			continue
		}
		result, err := calculate(input)
		if err != nil {
			fmt.Printf("Ошибка: %v\n", err)
			continue
		}
		calc := Calculation{
			Expression: input,
			Result:   result,
		}
		history = append(history, calc)
		fmt.Printf("Результат: %.2f\n", result)
	}
}
func calculate(expression string) (float64, error) {
	parts := strings.Fields(expression)
	if len(parts) != 3 {
		return 0, fmt.Errorf("неверный формат выражения. Используйте: число операция число")
	}
	num1, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return 0, fmt.Errorf("первое значение не является числом")
	}
	num2, err := strconv.ParseFloat(parts[2], 64)
	if err != nil {
		return 0, fmt.Errorf("второе значение не является числом")
	}
	operator := parts[1]
	var result float64
	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			return 0, fmt.Errorf("деление на ноль")
		}
		result = num1 / num2
	default:
		return 0, fmt.Errorf("неподдерживаемая операция: %s", operator)
	}
	return result, nil
}
func printHistory(history []Calculation) {
	if len(history) == 0 {
		fmt.Println("История пуста")
		return
	}
	fmt.Println("\nИстория вычислений:")
	for i, calc := range history {
		fmt.Printf("%d. %s = %.2f\n", i+1, calc.Expression, calc.Result)
	}
}