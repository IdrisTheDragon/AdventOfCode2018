package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("../myInput.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	fabric := [1000][1000]int{}

	for scanner.Scan() {
		line := scanner.Text()

		split := RegSplit(line,"[#@ ,:x]+")

		fmt.Println(split)

		offsetX, err := strconv.Atoi(split[2])
		if err != nil {
			log.Fatal(err)
		}
		offsetY, err := strconv.Atoi(split[3])
		if err != nil {
			log.Fatal(err)
		}
		sizeX, err := strconv.Atoi(split[4])
		if err != nil {
			log.Fatal(err)
		}
		sizeY, err := strconv.Atoi(split[5])
		if err != nil {
			log.Fatal(err)
		}
		
		for i := 0; i < sizeX; i++ {
			x := offsetX + i
			for j := 0; j < sizeY; j++ {
				y := offsetY + j
				fabric[x][y] = fabric[x][y] + 1
			}

		}
	}
	
	sum := 0

	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if fabric[i][j] > 1 {
				sum++
			}
		}

	}
	fmt.Println(sum)


	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func RegSplit(text string, delimeter string) []string {
    reg := regexp.MustCompile(delimeter)
    indexes := reg.FindAllStringIndex(text, -1)
    laststart := 0
    result := make([]string, len(indexes) + 1)
    for i, element := range indexes {
            result[i] = text[laststart:element[0]]
            laststart = element[1]
    }
    result[len(indexes)] = text[laststart:len(text)]
    return result
}
