package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)

type cord struct {
	y      int
	x      int
	height int
}

func getNeighbors(current cord, cords [][]rune) []cord {
	x := current.x
	y := current.y
	height := current.height
	var output []cord
	if x > 0 {
		nh := int(cords[y][x-1] - '0')
		if nh-1 == height {
			output = append(output, cord{x: x - 1, y: y, height: nh})
		}
	}
	if y > 0 {
		nh := int(cords[y-1][x] - '0')
		if nh-1 == height {
			output = append(output, cord{x: x, y: y - 1, height: nh})
		}
	}
	if x < len(cords[0])-1 {
		nh := int(cords[y][x+1] - '0')
		if nh-1 == height {
			output = append(output, cord{x: x + 1, y: y, height: nh})
		}
	}
	if y < len(cords)-1 {
		nh := int(cords[y+1][x] - '0')
		if nh-1 == height {
			output = append(output, cord{x: x, y: y + 1, height: nh})
		}
	}

	return output
}

func bfs(start cord, cords [][]rune) int {
	queue := list.New()
	visited := make(map[cord]bool)

	queue.PushBack(start)
	visited[start] = true
	total := 0
	for queue.Len() > 0 {
		element := queue.Front()
		node := element.Value.(cord)
		queue.Remove(element)
		for _, neighbor := range getNeighbors(node, cords) {
			if !visited[neighbor] {
				if neighbor.height == 9 {
					total++
				}
				queue.PushBack(neighbor)
				visited[neighbor] = true
			}
		}
	}

	return total
}

func bfs2(start cord, cords [][]rune) int {
	queue := list.New()
	visited := make(map[cord]bool)

	queue.PushBack(start)
	visited[start] = true
	total := 0
	for queue.Len() > 0 {
		element := queue.Front()
		node := element.Value.(cord)
		queue.Remove(element)
		for _, neighbor := range getNeighbors(node, cords) {
			if neighbor.height == 9 {
				total += 1
			}
			queue.PushBack(neighbor)
			visited[neighbor] = true
		}
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
	var amap [][]rune
	var trailheads []cord

	scanner := bufio.NewScanner(file)
	y := 0
	for scanner.Scan() {
		var line []rune
		row := scanner.Text()
		for x, char := range row {
			if char == '0' {
				trailheads = append(trailheads, cord{x: x, y: y, height: 0})
			}
			line = append(line, char)
		}
		amap = append(amap, line)
		y++
	}
	total := 0
	total2 := 0
	for _, c := range trailheads {
		total += bfs(c, amap)
		total2 += bfs2(c, amap)
	}
	fmt.Println("part1:", total)
	fmt.Println("part2:", total2)
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
