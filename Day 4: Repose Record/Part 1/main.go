package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("../myInput.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := []string{}

	for scanner.Scan() {
		line := scanner.Text()

		lines = append(lines, line)

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Strings(lines)

	guards := make(map[int] []int)
	guardsSum := make(map[int] int)

	id := 0
	start := 0

	for _, line := range lines {
		split := RegSplit(line, "[# :\\]]+")

		minute, err := strconv.Atoi(split[2])
		if err != nil {
			log.Fatal(err)
		}
		action := split[3][0]
		if action == 'G' {
			tempId, err := strconv.Atoi(split[4])
			if err != nil {
				log.Fatal(err)
			}
			id = tempId

			if _, ok := guards[id]; !ok {
				guards[id] = make([]int, 60)
			}

		} else if action == 'f' {
			start = minute
		} else if action == 'w' {
			end := minute
			temp := guards[id]
			for i := start; i < end; i++ {
				 temp[i] = temp[i] + 1
				 guardsSum[id] ++
			} 
			guards[id] = temp
		}
	}

	Key := 0
	biggestValue := 0
	
	for k,v := range guardsSum {
		if v > biggestValue {
			biggestValue = v
			Key = k
		}
	}

	sleepiest := guards[Key]

	mostTimesAsleep :=0
	minuteAsleep := 0

	for i,v := range sleepiest {
		if v > mostTimesAsleep {
			mostTimesAsleep = v
			minuteAsleep = i
		}
	}

	fmt.Println(Key,minuteAsleep, Key * minuteAsleep)
	

	
}

func RegSplit(text string, delimeter string) []string {
	reg := regexp.MustCompile(delimeter)
	indexes := reg.FindAllStringIndex(text, -1)
	laststart := 0
	result := make([]string, len(indexes)+1)
	for i, element := range indexes {
		result[i] = text[laststart:element[0]]
		laststart = element[1]
	}
	result[len(indexes)] = text[laststart:len(text)]
	return result
}
