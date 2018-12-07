package main

import (
	"sort"
	"fmt"
	"github.com/IdrisTheDragon/AdventOfCode2018/utils"
)

func main() {
	lines := utils.GetLines("../myInput.txt")

	instructions := make(map[rune] []rune)
	parents := make(map[rune] int)

	for _,k := range lines {
		key := rune(k[5])
		value := rune(k[36])
		instructions[key] = append(instructions[key],value)
		parents[value] = parents[value] + 1
	}

	done := make([]rune,0)
	for k,_ := range instructions {
		if parents[k] == 0 {
			done = append(done, k)
		}
	}

	answer := ""
	for ; len(done) > 0 ; {
		temp := make([]rune,len(done))
		copy(temp,done)
		sort.Sort(runes(temp))
		x := temp[0]
		for i := 0; i < len(done); i++ {
			if(done[i] == x){
				done = append(done[:i], done[i+1:]...)
			}
		}
		answer = answer + string(x)
		for _,v := range instructions[x] {
			parents[v] = parents[v] - 1
			if(parents[v] == 0){
				done = append(done,v)
			}
		}
	}

	fmt.Println(answer)
	
	
}

type runes []rune

func (s runes) Len() int {
	return len(s) 
}

func (s runes) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}
func (s runes) Less(i, j int) bool {
    return s[i] < s[j]
}





