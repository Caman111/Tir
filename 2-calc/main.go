package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	var menu = map[string]func([]float64) float64{
		"AVG": average,
		"SUM": sum,
		"MED": median,
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите операцию (AVG, SUM, MED): ")
	opInput, _ := reader.ReadString('\n')
	operation := strings.ToUpper(strings.TrimSpace(opInput))

	fmt.Print("Введите числа через запятую: ")
	numsInput, _ := reader.ReadString('\n')
	numsInput = strings.TrimSpace(numsInput)

	numStrs := strings.Split(numsInput, ",")
	var numbers []float64
	for _, s := range numStrs {
		s = strings.TrimSpace(s)
		num, err := strconv.ParseFloat(s, 64)
		if err != nil {
			fmt.Println("Ошибка преобразования числа:", s)
			return
		}
		numbers = append(numbers, num)
	}

	calkFunc, ok := menu[operation]
	if !ok {
		fmt.Println("Неизвестная операция")
		return
	}
	result := calkFunc(numbers)
	fmt.Println("Результат:", result)
}

func sum(nums []float64) float64 {
	var s float64
	for _, n := range nums {
		s += n
	}
	return s
}

func average(nums []float64) float64 {
	return sum(nums) / float64(len(nums))
}

func median(nums []float64) float64 {
	sort.Float64s(nums)
	n := len(nums)
	if n%2 == 1 {
		return nums[n/2]
	}
	return (nums[n/2-1] + nums[n/2]) / 2
}
