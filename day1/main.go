package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("input not provided")
		os.Exit(1)
	}

	masses, err := loadMasses(os.Args[1])
	if err != nil {
		fmt.Printf("couldn't load masses: %v\n", err)
		os.Exit(1)
	}

	fuel := 0
	total := 0
	for _, m := range masses {
		f := calcFuel(m)
		fuel += f
		total += f

		for {
			f = calcFuel(f)
			if f <= 0 {
				break
			}
			total += f
		}
	}

	fmt.Printf("fuel for modules: %d\n", fuel)
	fmt.Printf("total fuel: %d\n", total)
}

func loadMasses(path string) ([]int, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open file: %s: %w", path, err)
	}
	defer f.Close()

	result := make([]int, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, fmt.Errorf("convert line %s to int: %w", scanner.Text(), err)
		}

		result = append(result, i)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("unexpected scan error: %w", err)
	}

	return result, nil
}

func calcFuel(mass int) int {
	m := math.Floor(float64(mass) / 3)
	return int(m - 2)
}
