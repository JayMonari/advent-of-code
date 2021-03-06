package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	part1()
}

// part1 makes a graph of all line segments and their intersections and counts
// the amount of intersections.
func part1() {
	f, err := os.Open("input")
	check(err)
	graph := make([][]int, 1000)
	for i := range graph {
		graph[i] = make([]int, 1000)
	}
	scr := bufio.NewScanner(f)
	for scr.Scan() {
		// Need to parse line segments
		x1, x2, y1, y2 := Coords(scr.Text())
		// part1 shenanigans
		// if x1 != x2 && y1 != y2 {
		// 	continue
		// }
		step := 1
		switch {
		case x1 != x2 && y1 != y2:
			xstep := 1
			ystep := 1
			if x1 > x2 {
				xstep = -1
			}
			if y1 > y2 {
				ystep = -1
			}
			for i := x1; i != x2+xstep; i += xstep {
				graph[y1][x1]++
				y1 += ystep
				x1 += xstep
			}
		case x1 == x2:
			if y1 > y2 {
				step = -1
			}
			for i := y1; i != y2+step; i += step {
				graph[i][x1]++
			}
		case y1 == y2:
			if x1 > x2 {
				step = -1
			}
			for i := x1; i != x2+step; i += step {
				graph[y1][i]++
			}
		}
	}
	sum := 0
	for _, row := range graph {
		for _, n := range row {
			if n > 1 {
				sum++
			}
		}
	}
	fmt.Println(sum)
}

// Coords parses an input line of form `%d,%d -> %d,%d` into 4 coordinates.
func Coords(input string) (x1, x2, y1, y2 int) {
	f := strings.Fields(input)
	p1 := strings.Split(f[0], ",")
	p2 := strings.Split(f[2], ",")
	x1, _ = strconv.Atoi(p1[0])
	y1, _ = strconv.Atoi(p1[1])
	x2, _ = strconv.Atoi(p2[0])
	y2, _ = strconv.Atoi(p2[1])
	return
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
