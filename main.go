package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	AVG = "AVG"
	SUM = "SUM"
	MED = "MED"
)

func main() {
	op := scanUserOperation()
	nums := scanUserNumbers()
	res := getCalculation(op, nums)

	fmt.Printf("%g", res)
}

func scanUserOperation() string {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("Введите желаемую операцию (%s/%s/%s): ", AVG, SUM, MED)
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Некорректная операция")
			continue
		}
		op := strings.ToUpper(strings.TrimSpace(line))
		switch op {
		case AVG, SUM, MED:
			return op
		default:
			fmt.Println("Некорректная операция")
		}
	}
}

func scanUserNumbers() []float64 {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Введите числа через запятую:")
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Некорректный ввод")
			continue
		}

		parts := strings.Split(strings.TrimSpace(line), ",")

		nums := make([]float64, 0, len(parts))
		valid := true

		for _, part := range parts {
			part = strings.TrimSpace(part)
			if part == "" {
				valid = false
				break
			}
			num, err := strconv.ParseFloat(part, 64)
			if err != nil {
				valid = false
				break
			}
			nums = append(nums, num)
		}
		if !valid || len(nums) == 0 {
			fmt.Println("Некорректный ввод")
			continue
		}

		return nums
	}
}

func getCalculation(op string, nums []float64) float64 {
	var res float64
	switch op {
	case AVG:
		res = getAverage(nums)
	case SUM:
		res = getSum(nums)
	case MED:
		res = getMedian(nums)

	}
	return res
}

func getAverage(nums []float64) float64 {
	sum := getSum(nums)
	return sum / float64(len(nums))
}

func getSum(nums []float64) float64 {
	var sum float64

	for _, num := range nums {
		sum += num
	}
	return sum
}

func getMedian(nums []float64) float64 {
	l := len(nums)
	sorted := make([]float64, l)
	copy(sorted, nums)
	sort.Float64s(sorted)
	mid := l / 2

	if l%2 == 1 {
		return sorted[mid]
	}

	return (sorted[mid-1] + sorted[mid]) / 2
}
