package main

import (
	"fmt"

	"github.com/IdrisTheDragon/AdventOfCode2018/utils"
)

func main() {
	lines := utils.GetLines("../myInput.txt")
	//lines := utils.GetLines("../test.txt")
	worldOdd := make([][]rune, len(lines))
	for j := range worldOdd {
		worldOdd[j] = make([]rune, len(lines[j]))
	}
	worldEven := make([][]rune, len(lines))
	for j := range worldOdd {
		worldEven[j] = make([]rune, len(lines[j]))
	}

	for i, line := range lines {
		for j, c := range line {
			worldEven[i][j] = c
		}
	}

	var currentState *[][]rune

	fmt.Println("minute:", 0)
	for i := range worldEven {
		for j := range worldEven[i] {
			fmt.Print(string(worldEven[i][j]))
		}
		fmt.Println()
	}

	for minute := 1; minute <= 10; minute++ {
		var prevWorld, nextWorld *[][]rune
		if minute%2 == 1 {
			prevWorld = &worldEven
			nextWorld = &worldOdd
		} else {
			prevWorld = &worldOdd
			nextWorld = &worldEven
		}

		for i := range *prevWorld {
			for j := range (*prevWorld)[i] {
				switch (*prevWorld)[i][j] {
				case '.':
					neighbouringTrees := 0
					for _, offset := range [][]int{
						{-1, 0}, {1, 0}, {0, -1}, {0, 1},
						{1, 1}, {1, -1}, {-1, 1}, {-1, -1},
					} {
						if i+offset[0] >= 0 && i+offset[0] < len((*prevWorld)) &&
							j+offset[1] >= 0 && j+offset[1] < len((*prevWorld)[i+offset[0]]) &&
							(*prevWorld)[i+offset[0]][j+offset[1]] == '|' {
							neighbouringTrees++
						}
					}
					if neighbouringTrees >= 3 {
						(*nextWorld)[i][j] = '|'
					} else {
						(*nextWorld)[i][j] = '.'
					}
					break
				case '|':
					neighbouringLumber := 0
					for _, offset := range [][]int{
						{-1, 0}, {1, 0}, {0, -1}, {0, 1},
						{1, 1}, {1, -1}, {-1, 1}, {-1, -1},
					} {
						if i+offset[0] >= 0 && i+offset[0] < len((*prevWorld)) &&
							j+offset[1] >= 0 && j+offset[1] < len((*prevWorld)[i+offset[0]]) &&
							(*prevWorld)[i+offset[0]][j+offset[1]] == '#' {
							neighbouringLumber++
						}
					}
					if neighbouringLumber >= 3 {
						(*nextWorld)[i][j] = '#'
					} else {
						(*nextWorld)[i][j] = '|'
					}
					break
				case '#':
					neighbouringLumber := 0
					neighbouringTrees := 0
					for _, offset := range [][]int{
						{-1, 0}, {1, 0}, {0, -1}, {0, 1},
						{1, 1}, {1, -1}, {-1, 1}, {-1, -1},
					} {
						if i+offset[0] >= 0 && i+offset[0] < len((*prevWorld)) &&
							j+offset[1] >= 0 && j+offset[1] < len((*prevWorld)[i+offset[0]]) &&
							(*prevWorld)[i+offset[0]][j+offset[1]] == '#' {
							neighbouringLumber++
						} else if i+offset[0] >= 0 && i+offset[0] < len((*prevWorld)) &&
							j+offset[1] >= 0 && j+offset[1] < len((*prevWorld)[i+offset[0]]) &&
							(*prevWorld)[i+offset[0]][j+offset[1]] == '|' {
							neighbouringTrees++
						}
					}
					if neighbouringLumber >= 1 && neighbouringTrees >= 1 {
						(*nextWorld)[i][j] = '#'
					} else {
						(*nextWorld)[i][j] = '.'
					}

				}
			}
		}

		currentState = nextWorld

		fmt.Println("minute:", minute)
		for i := range *currentState {
			for j := range (*currentState)[i] {
				fmt.Print(string((*currentState)[i][j]))
			}
			fmt.Println()
		}

	}

	treeCount := 0
	lumberCount := 0

	for i := range *currentState {
		for j := range (*currentState)[i] {
			if (*currentState)[i][j] == '|' {
				treeCount++
			} else if (*currentState)[i][j] == '#' {
				lumberCount++
			}
		}
	}
	fmt.Println("trees:", treeCount, "lumber:", lumberCount, "result:", treeCount*lumberCount)

}
