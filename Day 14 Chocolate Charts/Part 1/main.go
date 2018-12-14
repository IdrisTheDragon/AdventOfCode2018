package main

import (
	"fmt"
	"strconv"
)

const ITERATIONS = 793031

func main() {
	recipies := []int{3,7}

	elf1 := 0
	elf2 := 1

	for i := 0; i < ITERATIONS + 10; i++ {
		newRecipies := recipies[elf1] + recipies[elf2]
		temp := strconv.Itoa(newRecipies)
		for _, v := range temp {
			recipies = append(recipies, int(v - 48))
		}
		

		elf1 = (elf1 + 1 + recipies[elf1]) % len(recipies)
		elf2 = (elf2 + 1 + recipies[elf2]) % len(recipies)

		//fmt.Println(recipies)
		//fmt.Println("elf1", elf1, recipies[elf1])
		//fmt.Println("elf2", elf2, recipies[elf2])
	}

	fmt.Println(recipies[ITERATIONS:ITERATIONS+10])
	

	
}


