package main

import (
	"image"
	"image/color"
	"image/png"
	"fmt"
	"os"
	"github.com/IdrisTheDragon/AdventOfCode2018/utils"
)

func main() {
	lines := utils.GetLines("../myInput.txt")
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
	smallestT = 10639


	 	maxX := 0
		maxY := 0
		minX := 0
		minY := 0

		for temp := head.next; temp.next != nil; temp = temp.next {
			temp.x = temp.x + temp.vX * smallestT
			if maxX < temp.x {
				maxX = temp.x
			} else if minX > temp.x{
				minX = temp.x
			}
			temp.y = temp.y + temp.vY * smallestT
			if maxY < temp.y {
				maxY = temp.y
			} else if minY > temp.y{
				minY = temp.y
			}
		}

		lengthX := maxX-minX+1
		lengthY := maxY-minY+1

		createImage(lengthX,lengthY,smallestT,head)
}

func createImage(lengthX,lengthY,t int, head *Star){
	
		img := image.NewRGBA(image.Rectangle{image.Point{0,0}, image.Point{lengthX-140,lengthY-100}})
		cyan := color.RGBA{100, 200, 200, 0xff}

		for temp := head.next; temp.next != nil; temp = temp.next {
			if temp.x < lengthX && temp.y <= lengthY {
				img.Set(temp.x-140,temp.y-100,cyan)
			}
		}
		

		f, _ := os.Create(fmt.Sprintf("image%d.png",t))
		png.Encode(f, img)
}

type Star struct {
	x int
	y int
	vX int
	vY int
	next *Star
}


