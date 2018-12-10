package main

import (
	"fmt"
	"math"
	"github.com/IdrisTheDragon/AdventOfCode2018/utils"
)

func main() {
	lines := utils.GetLines("../myInput.txt")
	line := lines[0]
	fmt.Println(len(line))

	alpha := make([]int, 26)
	for j := 65; j < 65+26; j++ {

		fmt.Print(string(j), " ")
		line = lines[0]
		for i := 0; i < len(line); i++ {
			if line[i] == byte(j) || line[i] == byte(j+32) {
				line = line[:i] + line[i+1:]
				i--
			}
		}
		fmt.Print(len(line), " ")

		found := true

		for found {
			found = false
			for i := 0; i < len(line)-1; i++ {
				if line[i] == line[i+1]+32 || line[i] == line[i+1]-32 {
					line = line[:i] + line[i+2:]
					found = true
				}
			}
		}

		alpha[j-65] = len(line)
		fmt.Println(len(line))
	}

	index := 0
	value := math.MaxInt32
	for n, v := range alpha {
		if(value > v){
			value = v
			index = n
		}
	}

	fmt.Println(value,string(index+65))
}
