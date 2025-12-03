package main

import "fmt"

const usdToEur = 0.85
const usdToRub = 90.0

const eurToRub = usdToRub / usdToEur

func readInput() (float64, string, string) {
	var amount float64
	var from, to string

	fmt.Print("Введите сумму: ")
	fmt.Scan(&amount)

	fmt.Print("Из какой валюты (USD/EUR/RUB): ")
	fmt.Scan(&from)

	fmt.Print("В какую валюту (USD/EUR/RUB): ")
	fmt.Scan(&to)

	return amount, from, to
}

func convert(amount float64, from string, to string) float64 {
	return 0
}

func main() {
	fmt.Println("Курсы:")
	fmt.Printf("1 USD = %.2f EUR\n", usdToEur)
	fmt.Printf("1 USD = %.2f RUB\n", usdToRub)
	fmt.Printf("1 EUR = %.2f RUB\n", eurToRub)

	amount, from, to := readInput()

	result := convert(amount, from, to)

	fmt.Printf("Результат: %.2f %s -> %.2f %s\n", amount, from, result, to)
}
