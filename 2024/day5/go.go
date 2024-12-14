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
func findValidOrder(update []int, rules []Tuple) []int {

	graph := make(map[int][]int)
	nodes := make(map[int]bool)

	for _, rule := range rules {
		before := rule.before
		after := rule.after
		inBefore := false
		inAfter := false
		for _, num := range update {
			if num == before {
				inBefore = true
				nodes[before] = true
			}
			if num == after {
				inAfter = true
				nodes[after] = true
			}
		}
		if inBefore && inAfter {
			if _, exists := graph[before]; !exists {
				graph[before] = make([]int, 0)
			}
			graph[before] = append(graph[before], after)
		}
	}

	inDegree := make(map[int]int)
	for node := range nodes {
		inDegree[node] = 0
	}
	for _, edges := range graph {
		for _, v := range edges {
			inDegree[v]++
		}
	}

	var queue []int
	for node := range nodes {
		if inDegree[node] == 0 {
			queue = append(queue, node)
		}
	}

	var result []int
	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		result = append(result, u)

		for _, v := range graph[u] {
			inDegree[v]--
			if inDegree[v] == 0 {
				queue = append(queue, v)
			}
		}
	}

	if len(result) != len(nodes) {
		return nil
	}

	return result
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var pairs []Tuple
	var pages [][]string
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

	totalPart1 := 0
	totalPart2 := 0
	for _, pageStr := range pages {
		update := make([]int, len(pageStr))
		for i, str := range pageStr {
			num, _ := strconv.Atoi(str)
			update[i] = num
		}

		if isValidUpdate(update, pairs) {
			middleIdx := len(update) / 2
			totalPart1 += update[middleIdx]
		} else {
			validOrder := findValidOrder(update, pairs)
			if validOrder != nil {
				middleIdx := len(validOrder) / 2
				totalPart2 += validOrder[middleIdx]
			}
		}
	}
	fmt.Println("total:", totalPart1)
	fmt.Println("total:", totalPart2)
}
