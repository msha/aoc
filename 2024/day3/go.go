package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func getToDont(text string) (string, int) {
	splet := strings.SplitN(text, "don't()", 2)
	count := countMul(splet[0])
	if len(splet) < 2 {
		return "", count
	}
	return splet[1], count
}

func getToDo(text string) (string, int) {
	splet := strings.SplitN(text, "do()", 2)
	if len(splet) < 2 {
		return "", 0
	}
	return splet[1], 0
}

func countMul(text string) int {
	re := regexp.MustCompile(`mul\([0-9]*,[0-9]*\)`)
	matches := re.FindAllString(text, -1)
	total := 0
	for _, match := range matches {
		re2 := regexp.MustCompile(`\d+`)
		matches2 := re2.FindAllString(match, -1)
		num, _ := strconv.Atoi(matches2[0])
		num2, _ := strconv.Atoi(matches2[1])
		total += num * num2
	}
	return total
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
		text := scanner.Text()
		for len(text) > 0 {
			var val int
			text, val = getToDont(text)
			total += val
			text, _ = getToDo(text)
		}
	}
	fmt.Println(total)
}
