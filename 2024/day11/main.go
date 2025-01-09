package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1(stones []int) int {
	i := 0
	for i < 25 {
		newStones := make([]int, 0)
		for _, stone := range stones {
			stoneStr := strconv.Itoa(stone)
			if stone == 0 {
				newStones = append(newStones, 1)
			} else if len(stoneStr)%2 == 0 {
				leftHalf, _ := strconv.Atoi(stoneStr[:len(stoneStr)/2])
				rightHalf, _ := strconv.Atoi(stoneStr[len(stoneStr)/2:])
				newStones = append(newStones, leftHalf, rightHalf)
			} else {
				stone *= 2024
				newStones = append(newStones, stone)
			}
		}
		stones = newStones
		i++
	}
	return len(stones)
}

type stackItem struct {
	face  int
	depth int
}

func part2(stones []int) int {
	memo := make(map[string]int)
	totalStones := 0

	for _, initialFace := range stones {
		stack := []stackItem{{face: initialFace, depth: 0}}
		results := make(map[string]int)

		for len(stack) > 0 {
			current := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			key := strconv.Itoa(current.face) + "_" + strconv.Itoa(current.depth)
			if count, exists := memo[key]; exists {
				results[key] = count
				continue
			}
			if current.depth == 75 {
				results[key] = 1
				memo[key] = 1
				continue
			}

			canCompute := true
			var sum int

			if current.face == 0 {
				childKey := "1_" + strconv.Itoa(current.depth+1)
				if childResult, exists := results[childKey]; exists {
					sum = childResult
				} else {
					stack = append(stack, current)
					stack = append(stack, stackItem{face: 1, depth: current.depth + 1})
					canCompute = false
				}
			} else {
				str := strconv.Itoa(current.face)
				if len(str)%2 == 0 {
					mid := len(str) / 2
					left, _ := strconv.Atoi(str[0:mid])
					right, _ := strconv.Atoi(str[mid:])

					leftKey := strconv.Itoa(left) + "_" + strconv.Itoa(current.depth+1)
					rightKey := strconv.Itoa(right) + "_" + strconv.Itoa(current.depth+1)

					leftResult, leftExists := results[leftKey]
					rightResult, rightExists := results[rightKey]

					if leftExists && rightExists {
						sum = leftResult + rightResult
					} else {
						stack = append(stack, current)
						if !rightExists {
							stack = append(stack, stackItem{face: right, depth: current.depth + 1})
						}
						if !leftExists {
							stack = append(stack, stackItem{face: left, depth: current.depth + 1})
						}
						canCompute = false
					}
				} else {
					newFace := current.face * 2024
					childKey := strconv.Itoa(newFace) + "_" + strconv.Itoa(current.depth+1)
					if childResult, exists := results[childKey]; exists {
						sum = childResult
					} else {
						stack = append(stack, current)
						stack = append(stack, stackItem{face: newFace, depth: current.depth + 1})
						canCompute = false
					}
				}
			}
			if canCompute {
				results[key] = sum
				memo[key] = sum
			}
		}
		startKey := strconv.Itoa(initialFace) + "_0"
		totalStones += results[startKey]
	}
	return totalStones
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	stones := make([]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words := strings.Split(scanner.Text(), " ")
		for _, word := range words {
			num, _ := strconv.Atoi(word)
			stones = append(stones, num)
		}
		fmt.Println("part1: ", part1(stones))
		fmt.Println("part2: ", part2(stones))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
