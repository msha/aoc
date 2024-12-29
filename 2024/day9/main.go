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
	organizeBlocks(blocks)
	//str := getBlockString(blocks)
	//fmt.Println(str)
	fmt.Println("part 1:", calcHash(blocks))

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
