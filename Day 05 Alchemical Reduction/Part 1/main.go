package main

import (
	"fmt"
	"github.com/IdrisTheDragon/AdventOfCode2018/utils"
)

func main() {
	lines := utils.GetLines("../myInput.txt")
	line := lines[0]
	fmt.Println(len(line))

	found := true

	for ; found ; {
		found = false
		for i := 0; i < len(line) - 1 ;i++ {
			if(line[i] == line[i+1] + 32 || line[i] == line[i+1] - 32 ){
				line = line[:i] + line[i+2:]
				found = true
			}
		}
	}


	fmt.Println(len(line))
}


