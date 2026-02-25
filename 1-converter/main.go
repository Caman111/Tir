package main

import (
	"fmt"
	"strings"
)

const usdToEur = 0.85
const usdToRub = 90.0
const eurToRub = usdToRub / usdToEur


func readCurrency(prompt string) string {
	for {
		fmt.Print(prompt)
		var cur string
		fmt.Scan(&cur)

		cur = strings.ToUpper(cur)

		if cur == "USD" || cur == "EUR" || cur == "RUB" {
			return cur
		}

		fmt.Println("Ошибка: доступные варианты USD, EUR, RUB. Попробуйте снова.")
	}
}

func readAmount() float64 {
	for {
		fmt.Print("Введите сумму: ")

		var amount float64
		_, err := fmt.Scan(&amount)

		if err == nil && amount >= 0 {
			return amount
		}

		fmt.Println("Ошибка: введите корректное число.")
	}
}

func convert(amount float64, from string, to string) float64 {
	if from == to {
		return amount
	}

	var inUSD float64

	switch from {
	case "USD":
		inUSD = amount
	case "EUR":
		inUSD = amount / usdToEur
	case "RUB":
		inUSD = amount / usdToRub
	}

	switch to {
	case "USD":
		return inUSD
	case "EUR":
		return inUSD * usdToEur
	case "RUB":
		return inUSD * usdToRub
	}

	return 0
}

func main() {
	fmt.Println("Курсы:")
	fmt.Printf("1 USD = %.2f EUR\n", usdToEur)
	fmt.Printf("1 USD = %.2f RUB\n", usdToRub)
	fmt.Printf("1 EUR = %.2f RUB\n", eurToRub)
	fmt.Println("----------------------------")

	from := readCurrency("Из какой валюты (USD/EUR/RUB): ")

	amount := readAmount()

	to := readCurrency("В какую валюту (USD/EUR/RUB): ")

	result := convert(amount, from, to)

	fmt.Println("----------------------------")
	fmt.Printf("Результат: %.2f %s → %.2f %s\n", amount, from, result, to)
}
