// main.go
package main

import "fmt"

const usdToEur = 0.85
const usdToRub = 90.0

const eurToRub = usdToRub / usdToEur

func main() {

	fmt.Println("Курсы:")
	fmt.Printf("1 USD = %.2f EUR\n", usdToEur)
	fmt.Printf("1 USD = %.2f RUB\n", usdToRub)
	fmt.Printf("1 EUR = %.2f RUB\n", eurToRub)
}
