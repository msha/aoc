package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var operators = []string{"+", "*"}

func getCombinations(n int) [][]string {
	if n == 0 {
		return [][]string{}
	}
	result := [][]string{}
	current := []string{}
	var helper func(int)
	helper = func(pos int) {
		if pos == n {
			combination := make([]string, len(current))
			copy(combination, current)
			result = append(result, combination)
			return
		}
		for _, op := range operators {
			current = append(current, op)
			helper(pos + 1)
			current = current[:len(current)-1]
		}
	}
	helper(0)
	return result
}

func checkExpr(numbers []int, ops []string) int {
	result := numbers[0]
	for i, op := range ops {
		num := numbers[i+1]
		if op == "+" {
			result += num
		} else if op == "*" {
			result *= num
		}
	}
	return result
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	total := 0
	for scanner.Scan() {
		words := strings.Split(scanner.Text(), ":")
		comp, _ := strconv.Atoi(words[0])
		inputArrays := []int{}
		numberStrings := strings.Fields(strings.TrimSpace(words[1]))
		for _, n := range numberStrings {
			num, _ := strconv.Atoi(n)
			inputArrays = append(inputArrays, num)
		}
		//fmt.Printf("Numbers: %v\n", inputArrays)
		operatorSlots := len(inputArrays) - 1
		combinations := getCombinations(operatorSlots)
		for _, ops := range combinations {
			value := checkExpr(inputArrays, ops)
			if value == comp {
				total += comp
				break
			}
		}
	}

	fmt.Println(total)
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
