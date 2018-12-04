package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("../myInput.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	countThree := 0
	countFour := 0

	for scanner.Scan() {
		m := make(map[rune] int)
		line := scanner.Text()
		if err != nil {
			log.Fatal(err)
		}
		
		for _, char := range line {
			m[char] = m[char] + 1
		}
		notThree := true
		notFour := true
		for _, count := range m {
			if count == 2 && notThree {
				notThree = false
				countThree++
			} else if count == 3 && notFour {
				countFour++
				notFour = false
			}
		}
	}
	checksum := countFour * countThree
	fmt.Println(countFour)
	fmt.Println(countThree)
	fmt.Println(checksum)
	

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
