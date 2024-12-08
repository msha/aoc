package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
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
		text := scanner.Text()
		re := regexp.MustCompile(`mul\([0-9]*,[0-9]*\)`)
		matches := re.FindAllString(text, -1)
		for _, match := range matches {
			re2 := regexp.MustCompile(`\d+`)
			matches2 := re2.FindAllString(match, -1)
			num, _ := strconv.Atoi(matches2[0])
			num2, _ := strconv.Atoi(matches2[1])
			fmt.Println(num * num2)
			total += num * num2
		}
	}
	fmt.Println(total)
}
