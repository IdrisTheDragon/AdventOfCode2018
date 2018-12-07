package utils

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

//Str2Int converts a string to an int and handles errors
func Str2Int(str string) int {
	i, err := strconv.Atoi(str)
		if err != nil {
			log.Fatal(err)
		}
		return i;
}


//GetLines gets the lines from a file
func GetLines(fileName string) []string {

	file, err := os.Open(fileName)
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

	return lines
}

//RegSplit A simple regex splitter
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
