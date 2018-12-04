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

	var lines []string

	for scanner.Scan() {
		line := scanner.Text()

		for _, line1 := range lines {
			result := compare(line, line1)
			if len(result) == (len(line) - 1) {
				fmt.Println(line)
				fmt.Println(line1)
				fmt.Println(result)
			}
		}

		lines = append(lines,line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func compare(line, line1 string) string {
	count := ""
	for num, _ := range line {
		if(line[num] == line1[num]){
			count = count + string(line[num])
		}
	}
	return count
}
