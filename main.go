package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Operation string

const (
	OpAvg Operation = "AVG"
	OpSum Operation = "SUM"
	OpMed Operation = "MED"
)

func main() {
	if err := run(os.Stdin, os.Stdout); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func run(in io.Reader, out io.Writer) error {
	scanner := bufio.NewScanner(in)

	op, err := readOperation(scanner, out)
	if err != nil {
		return err
	}
	nums, err := readNumbers(scanner, out)
	if err != nil {
		return err
	}
	result, err := op.Calc(nums)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(out, "Result: %g\n", result)
	return err
}

func readOperation(scanner *bufio.Scanner, out io.Writer) (Operation, error) {
	for {
		_, err := fmt.Fprintf(out, "Введите желаемую операцию (%s/%s/%s): ", OpAvg, OpSum, OpMed)
		if err != nil {
			return "", err
		}

		if !scanner.Scan() {
			return "", scanErr(scanner)
		}

		op, err := parseOperation(scanner.Text())
		if err == nil {
			return op, err
		}

		_, _ = fmt.Fprintln(out, "Некорректная операция")
	}
}

func readNumbers(scanner *bufio.Scanner, out io.Writer) ([]float64, error) {
	for {
		_, err := fmt.Fprintln(out, "Введите числа через запятую:")
		if err != nil {
			return nil, err
		}

		if !scanner.Scan() {
			return nil, scanErr(scanner)
		}

		nums, err := parseNumbers(scanner.Text())
		if err == nil {
			return nums, nil
		}

		_, _ = fmt.Fprintln(out, "Некорректный ввод")
	}
}

func parseOperation(s string) (Operation, error) {
	op := Operation(strings.ToUpper(strings.TrimSpace(s)))

	switch op {
	case OpAvg, OpSum, OpMed:
		return op, nil
	default:
		return "", fmt.Errorf("unknown operation: %q", s)
	}
}

func parseNumbers(s string) ([]float64, error) {
	parts := strings.Split(strings.TrimSpace(s), ",")

	if len(parts) == 0 {
		return nil, fmt.Errorf("empty input")
	}

	nums := make([]float64, 0, len(parts))
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			return nil, fmt.Errorf("empty number")
		}

		n, err := strconv.ParseFloat(part, 64)
		if err != nil {
			return nil, fmt.Errorf("parse float %q: %w", part, err)
		}

		nums = append(nums, n)
	}
	if len(nums) == 0 {
		return nil, fmt.Errorf("no numbers")
	}

	return nums, nil
}

func (op Operation) Calc(nums []float64) (float64, error) {
	if len(nums) == 0 {
		return 0, errors.New("no numbers")
	}

	switch op {
	case OpAvg:
		return average(nums), nil
	case OpSum:
		return sum(nums), nil
	case OpMed:
		return median(nums), nil
	default:
		return 0, fmt.Errorf("unsupported operation: %s", op)
	}
}

func average(nums []float64) float64 {
	return sum(nums) / float64(len(nums))
}

func sum(nums []float64) float64 {
	var total float64
	for _, n := range nums {
		total += n
	}
	return total
}

func median(nums []float64) float64 {
	sorted := append([]float64(nil), nums...)
	sort.Float64s(sorted)

	mid := len(sorted) / 2
	if len(sorted)%2 == 1 {
		return sorted[mid]
	}
	return (sorted[mid-1] + sorted[mid]) / 2
}

func scanErr(scanner *bufio.Scanner) error {
	if err := scanner.Err(); err != nil {
		return err
	}
	return io.EOF
}
