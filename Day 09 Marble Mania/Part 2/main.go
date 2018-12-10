package main

import (
	"fmt"
)

func main() {
	nextMarble := 0

	//myInput.txt
	lastMarble := 7095000 //70950 * 100
	players := make([]int, 431)
	
	//Example 1
	//lastMarble := 25
	//players := make([]int, 9)

	//Example 4
	//lastMarble := 1104
	//players := make([]int, 17)

	/*
This is basically a complete re-write of my part 1 
As when we are dealing with large arrays/slices it takes far too long to calculate. 
first implementain took about 3.5hours on my 2.7-3Ghz processor, this implementation in comparison takes seconds.
So using a circular double linked list I can add and remove from it without time consuming computations to shift all the elements.
I also don't need to keep track of the index of the current marble as I have a pointer straight to the current node.
Just need to keep track of the next and previous pointers when updating an element
	*/
	

	//add marble 0
	currentMarble := &Marble{ value: nextMarble}
	nextMarble++

	//initiate a circlular linked list with marble 1
	currentMarble.next = &Marble{ value: nextMarble, next: currentMarble, previous:currentMarble}
	currentMarble.previous = currentMarble.next
	nextMarble++

	for ; nextMarble <= lastMarble; nextMarble++ {
		if nextMarble%23 == 0 { //multiple of 23
			//find out the player
			playerIndex := (nextMarble - 1) % len(players)

			//step 1 keep marble that would have been placed
			players[playerIndex] = players[playerIndex] + nextMarble

			//step 2 7 marbles back is removed and added to score
			for i:=0; i < 7; i++ {
				currentMarble = currentMarble.previous
			}
			marbleForRemoval := currentMarble
			marbleForRemoval.next.previous = marbleForRemoval.previous
			marbleForRemoval.previous.next = marbleForRemoval.next
			
			players[playerIndex] = players[playerIndex] + currentMarble.value

			//step 3 the marble immediately clockwise becomes current marble
			currentMarble = marbleForRemoval.next

		} else {

			//add an marble skipping the one immediately clockwise to the currentMarble
			newMarble := &Marble{
				value: nextMarble,
				next: currentMarble.next.next,
				previous: currentMarble.next,
			}
			newMarble.previous.next = newMarble
			newMarble.next.previous = newMarble
			currentMarble = newMarble
		}
	}
	fmt.Println(Max(players))

}

type Marble struct {
	value int
	next *Marble
	previous *Marble
}

func Max(x []int) int {
	max := 0
	for _, v := range x {
		if v > max {
			max = v
		}
	}
	return max
}