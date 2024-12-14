package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Tuple struct {
	before int
	after  int
}

func isValidUpdate(update []int, rules []Tuple) bool {
	for _, rule := range rules {
		beforePos := -1
		afterPos := -1
		for i, num := range update {
			if num == rule.before {
				beforePos = i
			}
			if num == rule.after {
				afterPos = i
			}
		}
		if beforePos != -1 && afterPos != -1 {

			if beforePos > afterPos {
				return false
			}
		}
	}
	return true
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var pairs []Tuple
	var pages [][]string
	total := 0
	before := true

	for scanner.Scan() {
		if strings.TrimSpace(scanner.Text()) == "" {
			before = false
		}
		if before {

			b := strings.Split(scanner.Text(), "|")
			first, _ := strconv.Atoi(b[0])
			second, _ := strconv.Atoi(b[1])
			pairs = append(pairs, Tuple{before: first, after: second})
		} else {
			a := strings.Split(scanner.Text(), ",")

			if len(a) > 1 {
				pages = append(pages, a)
			}

		}

	}

	total = 0
	for _, pageStr := range pages {
		update := make([]int, len(pageStr))
		for i, str := range pageStr {
			num, _ := strconv.Atoi(str)
			update[i] = num
		}

		if isValidUpdate(update, pairs) {
			middleIdx := len(update) / 2
			total += update[middleIdx]
		}
	}
	fmt.Println("total:", total)
}
