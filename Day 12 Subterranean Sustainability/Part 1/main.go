package main

import (
	"fmt"
	"github.com/IdrisTheDragon/AdventOfCode2018/utils"
)

const GENERATIONS int = 20

func main() {
	lines := utils.GetLines("../myInput.txt")

	pots := utils.RegSplit(lines[0],": ")[1]
	rules := make(map[string] string)

	for i:=2; i < len(lines); i++ {
		split := utils.RegSplit(lines[i],"[ =>]+")
		rules[split[0]] = split[1]
	}

	fmt.Println(pots)
	arrayExpansion := 0
	for i:= 0; i < GENERATIONS; i++ {
		nextPots := ""
		for j:=-1; j < len(pots) + 1 ; j++ {
			neighbours := getNeighbours(pots,j)
			if val, ok := rules[neighbours]; ok {
				nextPots = nextPots + val
				if j == -1 {
					arrayExpansion = arrayExpansion + 1
				}
			} else {
				if j != -1 && j != len(pots)+1 {
					nextPots = nextPots + "."
				}
			}
		}
		pots = nextPots
		//fmt.Println(nextPots)
	}
	sum := 0
	for i:= 0; i < len(pots); i++ {
		if pots[i] == '#' {
			sum = sum + i - arrayExpansion
		}
	}	
	fmt.Println(sum)
}

func getNeighbours(pots string, index int) string {
	if index == -1 {
		return "..." + string(pots[0]) + string(pots[1])
	} else if index == 0 {
		return ".." + string(pots[0]) + string(pots[1]) + string(pots[2])
	} else if index == 1 {
		return "." + string(pots[0]) + string(pots[1]) + string(pots[2]) + string(pots[3])
	} else if index == len(pots)-2 {
		return string(pots[index-2]) + string(pots[index-1]) + string(pots[index]) + string(pots[index + 1]) + "."
	} else if index == len(pots) - 1 {
		return string(pots[index-2]) + string(pots[index-1]) + string(pots[index]) + ".."
	} else if index == len(pots) {
		return string(pots[index-2]) + string(pots[index-1]) + "..."
	} else {
		return string(pots[index-2]) + string(pots[index-1]) + string(pots[index]) + string(pots[index + 1]) + string(pots[index + 2])
	}
}


