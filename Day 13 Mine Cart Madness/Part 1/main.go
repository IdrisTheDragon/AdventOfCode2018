package main

import (
	"fmt"
	"github.com/IdrisTheDragon/AdventOfCode2018/utils"
)

func main() {
	track := make([][]rune,0)
	
	carts := make([]Cart,0)

	//Load in carts and track into useable structure
	lines := utils.GetLines("../myInput.txt")
	//lines := utils.GetLines("../test.txt")
	for i,line := range lines {
		track = append(track,make([]rune, 0))
		for j,s := range line {
			switch (s) {
			case '>':
				track[i] = append(track[i],'-')
				cart := Cart{x:j,y:i,dir:'>'}
				carts = append(carts,cart)
				break
			case '<':
				track[i] = append(track[i],'-')
				cart := Cart{x:j,y:i,dir:'<'}
				carts = append(carts,cart)
				break
			case '^':
				track[i] = append(track[i],'|')
				cart := Cart{x:j,y:i,dir:'^'}
				carts = append(carts,cart)
				break
			case 'v':
				track[i] = append(track[i],'|')
				cart := Cart{x:j,y:i,dir:'v'}
				carts = append(carts,cart)
				break
			default:
				track[i] = append(track[i],s)
				break
			}
		}
	}

	//loop to crash
	collision := false
	for !collision {
		//update carts
		//fmt.Println(i,carts)
		//printTrack(track,carts)
		for i:=0; i<len(carts); i++ {
			switch (carts[i].dir) {
			case '>':
				carts[i] = MovingRight(track, carts[i])
				break
			case '<':
				carts[i] = MovingLeft(track,carts[i])
				break
			case '^':
				carts[i] = MovingUp(track,carts[i])
				break
			case 'v':
				carts[i] = MovingDown(track,carts[i])
				break
			default:
				fmt.Println("error not valid cart")
				break
			}
		}

		//check collisions
		for i,cart := range carts {
			for j:=i+1; j < len(carts); j++ {
				if(cart.x == carts[j].x && cart.y == carts[j].y) {
						collision = true
						fmt.Println("Collision at :", cart.x, cart.y)
				}
			}
		}
	}
}

func printTrack(track [][]rune, carts []Cart){
	h := make([][]rune,0)

	for i,_ := range track {
		h = append(h,make([]rune,len(track[i])))
		copy(h[i], track[i])
	}

	for _,cart := range carts {
		h[cart.y][cart.x] = cart.dir
	}

	for _, row := range h{
		for _,s := range row {
			fmt.Print(string(s))
		}
		fmt.Println()
	}
}

func MovingDown(track [][]rune, cart Cart) Cart {
	switch(track[cart.y + 1][cart.x]){
	case '/':
		cart.dir = '<'
		break
	case '\\':
		cart.dir = '>'
		break
	case '+':
		if cart.turn == 0 {
			//left
			cart.dir = '>'
			cart.turn = 1
		} else if cart.turn == 1 {
			//straight
			cart.turn = 2
		} else if cart.turn == 2 {
			//right
			cart.dir = '<'
			cart.turn = 0
		}
		break;
	case '|':
		break
	default:
		fmt.Println("Error on track cart can't move :", cart.x , cart.y - 1, track[cart.y - 1][cart.x])
	}
	cart.y = cart.y + 1
	return cart
}

func MovingUp(track [][]rune, cart Cart) Cart {
	switch(track[cart.y - 1][cart.x]){
	case '/':
		cart.dir = '>'
		break
	case '\\':
		cart.dir = '<'
		break
	case '+':
		if cart.turn == 0 {
			//left
			cart.dir = '<'
			cart.turn = 1
		} else if cart.turn == 1 {
			//straight
			cart.turn = 2
		} else if cart.turn == 2 {
			//right
			cart.dir = '>'
			cart.turn = 0
		}
		break;
	case '|':
		break
	default:
		fmt.Println("Error on track cart can't move :", cart.x , cart.y - 1, track[cart.y - 1][cart.x])
	}
	cart.y = cart.y - 1
	return cart
}

func MovingLeft(track [][]rune, cart Cart) Cart{
	switch(track[cart.y][cart.x - 1]){
	case '/':
		cart.dir = 'v'
		break
	case '\\':
		cart.dir = '^'
		break
	case '+':
		if cart.turn == 0{
			//left
			cart.dir = 'v'
			cart.turn = 1
		} else if cart.turn == 1 {
			//straight
			cart.turn = 2
		} else if cart.turn == 2 {
			//right
			cart.dir = '^'
			cart.turn = 0
		}
		break;
	case '-':
		break
	default:
		fmt.Println("Error on track cart can't move :", cart.x - 1, cart.y, track[cart.y][cart.x-1])
	}
	cart.x = cart.x - 1
	return cart
}

func MovingRight(track [][]rune, cart Cart) Cart{
	switch(track[cart.y][cart.x + 1]){
	case '\\':
		cart.dir = 'v'
		break
	case '/':
		cart.dir = '^'
		break
	case '+':
		if cart.turn == 0 {
			//left
			cart.dir = '^'
			cart.turn = 1
		} else if cart.turn == 1 {
			//straight	
			cart.turn = 2
		} else if cart.turn == 2 {
			//right
			cart.dir = 'v'
			cart.turn = 0
		}
		break;
	case '-':
		break
	default:
		fmt.Println("Error on track cart can't move :", cart.x + 1, cart.y, track[cart.y][cart.x+1])
	}
	cart.x = cart.x + 1
	return cart
}

type Cart struct {
	x int
	y int
	dir rune
	turn int
}

