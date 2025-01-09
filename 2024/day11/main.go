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
		//fmt.Println(stones)
		stones = newStones
		i++
	}

	return len(stones)
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
		fmt.Println(part1(stones))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
