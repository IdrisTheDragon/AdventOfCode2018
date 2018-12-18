package main

import (
	"fmt"
	"github.com/IdrisTheDragon/AdventOfCode2018/utils"
)

func main() {
	lines := utils.GetLines("../myInput.txt")
	//lines := utils.GetLines("../bigInput.txt")
	head := &Star{}
	tail := head

	

	for _,line := range lines {
		split := utils.RegSplit(line,"[=< ,>]+")
		star := &Star{ 
			x: utils.Str2Int(split[1]),
			y: utils.Str2Int(split[2]),
			vX: utils.Str2Int(split[4]),
			vY: utils.Str2Int(split[5]),
		}
		tail.next = star
		tail = star
	}

	smallestT := 0 
	smallestArea := int(^uint(0) >> 1)
	for t := 1; t < 100000; t++ {
		maxX := 0
		maxY := 0
		minX := 0
		minY := 0

		for temp := head.next; temp.next != nil; temp = temp.next {
			x := temp.x + temp.vX * t
			if maxX < x {
				maxX = x
			} else if minX > x{
				minX = x
			}
			y := temp.y + temp.vY * t
			if maxY < y {
				maxY = y
			} else if minY > y{
				minY = y
			}
		}

		lenX := maxX - minY + 1
		lenY := maxY - minY + 1
		area := lenX + lenY

		if(smallestArea > area) {
			smallestArea = area
			smallestT = t
		}		
	}
	fmt.Println(smallestT)


	 t := smallestT

	 	maxX := 0
		maxY := 0
		minX := 0
		minY := 0

		for temp := head.next; temp.next != nil; temp = temp.next {
			temp.x = temp.x + temp.vX * t
			if maxX < temp.x {
				maxX = temp.x
			} else if minX > temp.x{
				minX = temp.x
			}
			temp.y = temp.y + temp.vY * t
			if maxY < temp.y {
				maxY = temp.y
			} else if minY > temp.y{
				minY = temp.y
			}
		}

		mapper:= make ([][]bool,maxY-minY+1)

		for i:=0 ; i<len(mapper); i++ {
			mapper[i] = make([]bool,maxX-minX+1)
		}

		for temp := head.next; temp.next != nil; temp = temp.next {
			mapper[temp.y][temp.x] = true
		}

		for i :=  0;i < len(mapper); i++ {
			for j := 0;j < len(mapper[0]); j++ {
				if mapper[i][j] {
					fmt.Print("#")
				} else {
					fmt.Print(" ")
				}
			}
			fmt.Println()
		}


}

type Star struct {
	x int
	y int
	vX int
	vY int
	next *Star
}


