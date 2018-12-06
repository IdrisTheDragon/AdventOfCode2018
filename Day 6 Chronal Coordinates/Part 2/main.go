package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/IdrisTheDragon/AdventOfCode2018/utils"
)

func main() {
	lines := utils.GetLines("../myInput.txt")

	x := make([]int, len(lines))
	y := make([]int, len(lines))

	bigX, bigY := 0, 0

	for n, v := range lines {
		split := utils.RegSplit(v, "[, ]+")
		i, err := strconv.Atoi(split[0])
		if err != nil {
			log.Fatal(err)
		}
		x[n] = i

		if bigX < i {
			bigX = i
		}

		j, err := strconv.Atoi(split[1])
		if err != nil {
			log.Fatal(err)
		}
		y[n] = j

		if bigY < j {
			bigY = j
		}
	}

	grid := make([][]int, bigX+1)
	for i := range grid {
		grid[i] = make([]int, bigY+1)
	}
	region := 0
	for i := 0; i < bigX+1; i++ {
		for j := 0; j < bigY+1; j++ {
			s := sum(i, j, x, y)
			//fmt.Println(c)
			grid[i][j] = s
			if s < 10000 {
				region++
			}
		}
	}

	fmt.Println(region)
}

func sum(x, y int, x1, y1 []int) int {
	distance := 0
	for n, _ := range y1 {
		man := manhattanDistance(x, y, x1[n], y1[n])
		//fmt.Println(x,y,":",x1[n],y1[n],"",man)
		distance = distance + man
	}
	return distance
}

func manhattanDistance(x1, y1, x2, y2 int) int {
	return abs(x1-x2) + abs(y1-y2)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
