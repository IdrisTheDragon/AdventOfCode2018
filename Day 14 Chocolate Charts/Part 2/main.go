package main

import (
	"fmt"
	"strconv"
)



func main() {
	recipies := []int{3,7}
	SEQUENCE := []int{7,9,3,0,3,1}

	elf1 := 0
	elf2 := 1
	foundSequence := false
	for !foundSequence {
		newRecipies := recipies[elf1] + recipies[elf2]
		temp := strconv.Itoa(newRecipies)
		for _, v := range temp {
			recipies = append(recipies, int(v - 48))
		}
		

		elf1 = (elf1 + 1 + recipies[elf1]) % len(recipies)
		elf2 = (elf2 + 1 + recipies[elf2]) % len(recipies)

		if len(recipies) - len(SEQUENCE) > 0 {
		 if testEq(recipies[ len(recipies) - len(SEQUENCE):], SEQUENCE) {
			fmt.Println(len(recipies)-len(SEQUENCE))
			foundSequence = true
		 } else if testEq(recipies[ len(recipies) - len(SEQUENCE) - 1 :len(recipies)-1], SEQUENCE) {
			fmt.Println(len(recipies) - len(SEQUENCE) - 1)
			foundSequence = true
		 }
		}
	}	
}

func testEq(a, b []int) bool {

    // If one is nil, the other must also be nil.
    if (a == nil) != (b == nil) { 
        return false; 
    }

    if len(a) != len(b) {
        return false
    }

    for i := range a {
        if a[i] != b[i] {
            return false
        }
    }

    return true
}


