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
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var list1, list2 []int

	for scanner.Scan() {
		words := strings.Split(scanner.Text(), "   ")
		num1, _ := strconv.Atoi(words[0])
		num2, _ := strconv.Atoi(words[1])
		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}

	sort.Ints(list1)
	sort.Ints(list2)

	total := 0

	for _, value := range list1 {
		found := 0
		for _, value2 := range list2 {
			if value == value2 {
				found += 1
			}

		}
		total += value * found

	}
	fmt.Println(total)
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
