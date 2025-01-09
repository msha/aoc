package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Machine struct {
	Ax, Ay int
	Bx, By int
	Px, Py int
}

func parseLine(line string) (x, y int) {
	re := regexp.MustCompile(`X([=+\-]?)(\d+).*Y([=+\-]?)(\d+)`)
	matches := re.FindStringSubmatch(line)

	getValue := func(op string, numStr string) int {
		sign := 1
		if op == "-" {
			sign = -1
		}
		value, _ := strconv.Atoi(numStr)
		return sign * value
	}

	xVal := getValue(matches[1], matches[2])
	yVal := getValue(matches[3], matches[4])

	return xVal, yVal
}

func calcMin(m Machine) int {
	val := -1
	for a := 0; a <= 100; a++ {
		for b := 0; b <= 100; b++ {
			xPos := m.Ax*a + m.Bx*b
			yPos := m.Ay*a + m.By*b
			if xPos == m.Px && yPos == m.Py {
				cost := 3*a + b
				if val == -1 || cost < val {
					val = cost
				}
			}
		}
	}
	return val
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var machines []Machine
	var buffer []string

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		buffer = append(buffer, line)
		if len(buffer) == 3 {
			Ax, Ay := parseLine(buffer[0])
			Bx, By := parseLine(buffer[1])
			Px, Py := parseLine(buffer[2])

			machines = append(machines, Machine{
				Ax: Ax, Ay: Ay,
				Bx: Bx, By: By,
				Px: Px, Py: Py,
			})
			buffer = buffer[:0]
		}
	}

	total := 0
	winnableCount := 0
	for _, m := range machines {
		cost := calcMin(m)
		if cost != -1 {
			winnableCount++
			total += cost
		}
	}

	fmt.Printf("Part1: %d\n", total)
}
