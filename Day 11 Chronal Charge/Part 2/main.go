package main

import (
	"fmt"
	//"math"
)

const SERIAL int = 1723
//const SERIAL int = 18
const GRIDSIZE int = 300

func main() {
	grid := make([][]int,GRIDSIZE)
	for i:=0; i<len(grid); i++ {
		grid[i] = make([]int,GRIDSIZE)
	}

	for i:=0; i<len(grid); i++ {
		for j:=0; j<len(grid); j++ {
			grid[i][j] = powerLevel(i,j)
		}
	}

	largestPower := 0
	largestPowerX := 0
	largestPowerY := 0
	largestSquareSize := 0
	for squareSize := 1; squareSize < GRIDSIZE; squareSize++ {
	for i:=0; i<len(grid)-squareSize+1; i++ {
		for j:=0; j<len(grid)-squareSize+1; j++ {
			power := 0
			for x:=0; x <squareSize; x++ {
				for y:=0;y<squareSize;y++{
					power = power + grid[i+x][j+y]
				}
			}
			if(power > largestPower){
				largestPower = power
				largestPowerX = i
				largestPowerY = j
				largestSquareSize = squareSize
			}

		}
	}
}
	fmt.Println(largestPower,largestPowerX,largestPowerY,largestSquareSize)

}




func powerLevel(x,y int) int{
	//Find the fuel cell's rack ID, which is its X coordinate plus 10.
	rackId := x + 10
	//Begin with a power level of the rack ID times the Y coordinate.
	powerLevel := rackId * y
	//Increase the power level by the value of the grid serial number (your puzzle input).
	powerLevel = powerLevel + SERIAL
	//Set the power level to itself multiplied by the rack ID.
	powerLevel = powerLevel * rackId
	//Keep only the hundreds digit of the power level (so 12345 becomes 3; numbers with no hundreds digit become 0).
	powerLevel = (powerLevel/100)%10
	//Subtract 5 from the power level.
	powerLevel = powerLevel - 5  
	return powerLevel
}

