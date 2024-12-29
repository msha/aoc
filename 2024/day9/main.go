package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type block struct {
	num   int
	empty bool
}

func deepCopy(original []block) []block {
	copy := make([]block, len(original))
	for i := range original {
		copy[i] = original[i]
	}
	return copy
}

func getBlockString(blocks []block) string {
	output := ""
	for _, bl := range blocks {
		if !bl.empty {
			value := strconv.Itoa(bl.num)
			output += value
		} else {
			output += "."
		}
	}
	return output
}

func organizeBlocks(blocks []block) []block {
	last := len(blocks) - 1
	for index, bl := range blocks {
		if index >= last {
			break
		}
		if bl.empty {
			ndex := last
			for i := ndex; i > index; i-- {
				if !blocks[i].empty {
					blocks[index] = blocks[i]
					blocks[i].empty = true
					blocks[i].num = 0
					break
				}
			}
			last = ndex
		}
	}
	return blocks
}

func findAndMoveBlocks(blocks []block, num int, end int, start int) []block {
	size := end - start + 1

	bestStart := -1
	currentStart := -1
	currentLength := 0

	for i := 0; i < start; i++ {
		if blocks[i].empty {
			if currentStart == -1 {
				currentStart = i
			}
			currentLength++
			if currentLength == size {
				bestStart = currentStart
				break
			}
		} else {
			currentStart = -1
			currentLength = 0
		}
	}

	if bestStart != -1 {
		for i := 0; i < size; i++ {
			blocks[bestStart+i].num = num
			blocks[bestStart+i].empty = false
		}
		for i := start; i <= end; i++ {
			blocks[i].num = 0
			blocks[i].empty = true
		}
	}

	return blocks
}

func organizeBlocks2(blocks []block) []block {
	maxID := 99999
	for id := maxID; id >= 0; id-- {
		start := -1
		end := -1

		for i := 0; i < len(blocks); i++ {
			if !blocks[i].empty && blocks[i].num == id {
				if start == -1 {
					start = i
				}
				end = i
			} else if start != -1 {
				blocks = findAndMoveBlocks(blocks, id, end, start)
				start = -1
				end = -1
			}
		}
		if start != -1 {
			blocks = findAndMoveBlocks(blocks, id, end, start)
		}
	}
	return blocks
}

func calcHash(blocks []block) int {
	total := 0
	for i, bl := range blocks {
		total += i * bl.num
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
	var blocks []block
	curValue := 0
	for scanner.Scan() {
		words := strings.Split(scanner.Text(), "")
		for i, n := range words {
			if i%2 == 0 {
				value, _ := strconv.Atoi(n)
				for i := 0; i < value; i++ {
					newBlock := block{num: curValue, empty: false}
					blocks = append(blocks, newBlock)
				}
			} else {
				curValue++
				value, _ := strconv.Atoi(n)
				for i := 0; i < value; i++ {
					newBlock := block{num: 0, empty: true}
					blocks = append(blocks, newBlock)
				}
			}
		}
	}
	var blocks2 []block
	blocks2 = deepCopy(blocks)
	organizeBlocks(blocks)
	//str := getBlockString(blocks)
	//fmt.Println(str)
	organizeBlocks2(blocks2)
	//str = getBlockString(blocks2)
	//fmt.Println(str)
	fmt.Println("part 1:", calcHash(blocks))
	fmt.Println("part 2:", calcHash(blocks2))
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
