package main

import (
	"bufio"
	"fmt"
	"os"
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

	total := 0
	for scanner.Scan() {
		words := strings.Split(scanner.Text(), " ")
		numbers := make([]int, len(words))
		for i, word := range words {
			num, err := strconv.Atoi(word)
			if err != nil {
				return
			}
			numbers[i] = num
		}

		valid := false

		dird := true
		for i := 1; i < len(numbers); i++ {
			diff := numbers[i-1] - numbers[i]
			if diff <= 0 || diff > 3 {
				dird = false
				break
			}
		}

		diri := true
		for i := 1; i < len(numbers); i++ {
			diff := numbers[i] - numbers[i-1]
			if diff <= 0 || diff > 3 {
				diri = false
				break
			}
		}

		valid = diri || dird

		if !valid {
			for skip := 0; skip < len(numbers); skip++ {
				dird = true
				diri = true

				prev := -1
				for i := 0; i < len(numbers); i++ {
					if i == skip {
						continue
					}
					if prev != -1 {
						diff := numbers[prev] - numbers[i]
						if diff <= 0 || diff > 3 {
							dird = false
						}
						diff = numbers[i] - numbers[prev]
						if diff <= 0 || diff > 3 {
							diri = false
						}
					}
					prev = i
				}
				if dird || diri {
					valid = true
					break
				}
			}
		}

		if valid {
			total += 1
		}
	}
	println("total:", total)
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
