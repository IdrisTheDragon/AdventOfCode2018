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

	for {
		file, err := os.Open("../myInput.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			sums = append(sums, sum)

			line := scanner.Text()
			i, err := strconv.Atoi(line[1:])
			if err != nil {
				log.Fatal(err)
			}
			if line[0] == '+' {
				sum = sum + i
			} else if line[0] == '-' {
				sum = sum - i
			} else {
				fmt.Println("error")
			}
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
