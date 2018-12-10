package main

import (
	"fmt"
)

func main() {
	nextMarble := 0
	lastMarble := 70950
	players := make([]int, 431)
	
	//Example 1
	//lastMarble := 25
	//players := make([]int, 9)

	//lastMarble := 1104
	//players := make([]int, 17)

	result := make([]int, 0)

	//add the 0
	result = append(result, nextMarble)
	nextMarble++
	currentMarbleIndex := 0
	//fmt.Println("-1", currentMarbleIndex, result)

	for ; nextMarble <= lastMarble; nextMarble++ {
		playerIndex := (nextMarble - 1) % len(players)

		if nextMarble%23 == 0 {
			//step 1 keep mrble that would have been placed
			players[playerIndex] = players[playerIndex] + nextMarble

			//step 2 7 marbles back is removed and added to score
			removeMarblesIndex := currentMarbleIndex - 7
			if removeMarblesIndex < 0 {removeMarblesIndex = removeMarblesIndex + len(result)}
			players[playerIndex] = players[playerIndex] + result[removeMarblesIndex]

			result = append(result[:removeMarblesIndex], result[removeMarblesIndex+1:]...)

			//step 3 the marble immediately clockwise becomes current marble
			currentMarbleIndex = removeMarblesIndex

		} else {
			nextMarblesIndex := currentMarbleIndex + 2
			if nextMarblesIndex < len(result) {
				result = addValueAtIndex(result, nextMarble, nextMarblesIndex)
			} else if nextMarblesIndex == len(result) {
				result = append(result, nextMarble)
			} else { //(nextMarblesIndex > len(result))
				nextMarblesIndex = 1
				result = addValueAtIndex(result, nextMarble, nextMarblesIndex)
			}
			currentMarbleIndex = nextMarblesIndex
		}

		//fmt.Println(playerIndex+1, currentMarbleIndex, result)
	}
	fmt.Println(max(players))

}

func addValueAtIndex(result []int, value, index int) []int {
	result = append(result, 0)
	copy(result[index+1:], result[index:])
	result[index] = value
	return result
}

func max(x []int) int {
	max := 0
	for _, v := range x {
		if v > max {
			max = v
		}
	}
	return max
}
