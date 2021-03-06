package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	sum := 0
	var sums []int

	//Thank Jassler for the improvement :)
	file, err := os.Open("../myInput.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	for {
		// jump to beggining of the file
		file.Seek(0, 0)
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			sums = append(sums, sum)

			line := scanner.Text()
			i, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal(err)
			}
			sum += i
			//fmt.Println(sums)
			for _, v := range sums {
				if v == sum {
					fmt.Println(sum)
					return
				}
			}
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
}
