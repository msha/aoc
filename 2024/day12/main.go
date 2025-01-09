package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)

type cord struct {
	y int
	x int
	t rune
}

type region struct {
	plots     []cord
	plant     rune
	area      int
	perimeter int
}

func getNeighbors(current cord, cords [][]cord) []cord {
	x := current.x
	y := current.y
	plant := current.t
	var output []cord
	if x > 0 {
		nh := cords[y][x-1].t
		if nh == plant {
			output = append(output, cords[y][x-1])
		}
	}
	if y > 0 {
		nh := cords[y-1][x].t
		if nh == plant {
			output = append(output, cords[y-1][x])
		}
	}
	if x < len(cords[0])-1 {
		nh := cords[y][x+1].t
		if nh == plant {
			output = append(output, cords[y][x+1])
		}
	}
	if y < len(cords)-1 {
		nh := cords[y+1][x].t
		if nh == plant {
			output = append(output, cords[y+1][x])
		}
	}

	return output
}

func bfs(start cord, cords [][]cord, visited map[cord]bool) []cord {
	var region []cord
	queue := list.New()

	queue.PushBack(start)
	visited[start] = true
	region = append(region, start)

	for queue.Len() > 0 {
		element := queue.Front()
		node := element.Value.(cord)
		queue.Remove(element)

		for _, neighbor := range getNeighbors(node, cords) {
			if !visited[neighbor] {
				queue.PushBack(neighbor)
				visited[neighbor] = true
				region = append(region, neighbor)
			}
		}
	}

	return region
}

func getRegions(plots []cord, cords [][]cord) int {
	perimeter := 0
	plotMap := make(map[cord]bool)

	for _, plot := range plots {
		plotMap[plot] = true
	}

	for _, plot := range plots {
		x, y := plot.x, plot.y

		if x == 0 || !plotMap[cord{x: x - 1, y: y, t: plot.t}] {
			perimeter++
		}
		if x == len(cords[0])-1 || !plotMap[cord{x: x + 1, y: y, t: plot.t}] {
			perimeter++
		}
		if y == 0 || !plotMap[cord{x: x, y: y - 1, t: plot.t}] {
			perimeter++
		}
		if y == len(cords)-1 || !plotMap[cord{x: x, y: y + 1, t: plot.t}] {
			perimeter++
		}
	}
	return perimeter
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var amap [][]cord
	scanner := bufio.NewScanner(file)
	y := 0
	for scanner.Scan() {
		var line []cord
		row := scanner.Text()
		for x, char := range row {
			place := cord{x: x, y: y, t: char}
			line = append(line, place)
		}
		amap = append(amap, line)
		y++
	}

	visited := make(map[cord]bool)
	var regions []region
	totalPrice := 0

	for y := 0; y < len(amap); y++ {
		for x := 0; x < len(amap[0]); x++ {
			current := amap[y][x]
			if !visited[current] {
				plots := bfs(current, amap, visited)
				perimeter := getRegions(plots, amap)
				area := len(plots)
				price := area * perimeter

				regions = append(regions, region{
					plots:     plots,
					plant:     current.t,
					area:      area,
					perimeter: perimeter,
				})
				totalPrice += price
			}
		}
	}

	fmt.Println("Part 1:", totalPrice)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
