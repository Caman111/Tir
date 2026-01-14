package main

import (
	"fmt"
	"strings"
)

const usdToEur = 0.85
const usdToRub = 90.0
const eurToRub = usdToRub / usdToEur

var rates = map[string]float64{
	"USD": 1.0,
	"EUR": usdToEur,
	"RUB": usdToRub}

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

func convertRates(rates *map[string]float64, amount float64, from string, to string) float64 {

	rateMap := *rates

	rateFrom, okFrom := rateMap[from]
	rateTo, okTo := rateMap[to]

	if !okFrom || !okTo {
		fmt.Println("Ошибка: неверная валюта для конвертации")
		return 0
	}
	amountInUSD := amount / rateFrom
	result := amountInUSD * rateTo
	return result
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

	result := convertRates(&rates, amount, from, to)

	fmt.Println("----------------------------")
	fmt.Printf("Результат: %.2f %s → %.2f %s\n", amount, from, result, to)
}
