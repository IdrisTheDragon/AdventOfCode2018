package main

import (
	"log"
	"math"
	"os"
	"strconv"

	"github.com/IdrisTheDragon/AdventOfCode2018/utils"
)

func main() {
	downscaleX, downscaleY := 8,3

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

	f, err := os.Create("../out.txt")
	check(err)
	defer f.Close()

	for i := 0; i < bigX+1; i++ {
		for j := 0; j < bigY+1; j++ {
			c := closest(i, j, x, y)
			if i%downscaleX == 0 && j%downscaleY == 0 {
				_, err := f.WriteString(string(c + 65))
				check(err)
			}
		}
		if i%downscaleX == 0 {
			_, err := f.WriteString("\n")
			check(err)
		}
	}
}

func closest(x, y int, x1, y1 []int) int {
	closest, distance := 0, math.MaxInt32
	equal := false
	for n := range y1 {
		man := manhattanDistance(x, y, x1[n], y1[n])
		//fmt.Println(x,y,":",x1[n],y1[n],"",man)
		if man < distance {
			distance = man
			closest = n
			equal = false
		} else if man == distance {
			equal = true
		}
	}
	if equal {
		return -1
	} else {
		return closest
	}
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

func check(e error) {
	if e != nil {
		panic(e)
	}
}
