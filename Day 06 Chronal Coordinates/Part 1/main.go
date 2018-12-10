package main

import (
	"fmt"
	"github.com/IdrisTheDragon/AdventOfCode2018/utils"
	"log"
	"strconv"
	"math"
)

func main() {
	lines := utils.GetLines("../myInput.txt")

	x := make([] int, len(lines))
	y := make([] int, len(lines))

	bigX, bigY := 0,0

	for n,v := range lines {
		split := utils.RegSplit(v,"[, ]+")
		i, err := strconv.Atoi(split[0])
		if err != nil {
			log.Fatal(err)
		}
		x[n] = i

		if(bigX < i){
			bigX = i
		}

		j, err := strconv.Atoi(split[1])
		if err != nil {
			log.Fatal(err)
		}
		y[n] = j

		if(bigY < j){
			bigY = j
		}
	}

	grid := make([][]int,bigX+1)
	for i := range grid {
		grid[i] = make ([]int, bigY+1)
	}
	counted := make([]int,len(x))
	for i := 0; i < bigX+1; i++ {
		for j := 0; j < bigY+1; j++ {
			c := closest(i,j,x,y)
			//fmt.Println(c)
			grid[i][j] = c
			if c!= -1 && counted[c] != -1 {
				if i == bigX || i == 0 || j == 0  || j == bigY {
					counted[c] = -1
				} else {
					counted[c]++
				}
			}
		}
	}
	fmt.Println(counted)
	
	value := 0
	for i := 0 ; i < len(counted); i++ {
		if counted[i] > value {
			value = counted[i]
		}
	}

	fmt.Println(value)
}

func closest(x,y int , x1 , y1 [] int) int {
	closest, distance := 0,math.MaxInt32
	equal := false
	for n,_ := range y1 {
		man := manhattanDistance(x,y,x1[n],y1[n])
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

func manhattanDistance(x1,y1,x2,y2 int) int {
	return abs(x1-x2) + abs(y1-y2)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

