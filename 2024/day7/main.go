package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var operators = []string{"+", "*"}
var operators2 = []string{"+", "*", "|"}

func getCombinations(n int, ops []string) [][]string {
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
		for _, op := range ops {
			current = append(current, op)
			helper(pos + 1)
			current = current[:len(current)-1]
		}
	}
	helper(0)
	return result
}

func checkExpr(numbers []int, ops []string) int {
	if len(numbers) == 0 {
		return 0
	}

	result := numbers[0]

	for i, op := range ops {
		num := numbers[i+1]
		switch op {
		case "+":
			result += num
		case "*":
			result *= num
		case "|":
			concatenated, _ := strconv.Atoi(fmt.Sprintf("%d%d", result, num))
			result = concatenated
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
	total2 := 0
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
		combinations := getCombinations(operatorSlots, operators)
		for _, ops := range combinations {
			value := checkExpr(inputArrays, ops)
			if value == comp {
				total += comp
				break
			}
		}
		combinations2 := getCombinations(operatorSlots, operators2)
		for _, ops := range combinations2 {
			value := checkExpr(inputArrays, ops)
			if value == comp {
				total2 += comp
				break
			}
		}
	}

	fmt.Println("part1 total:", total)
	fmt.Println("part2 total:", total2)
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
