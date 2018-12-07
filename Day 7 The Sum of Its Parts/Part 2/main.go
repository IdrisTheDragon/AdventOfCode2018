package main

import (
	"fmt"
	"sort"

	"github.com/IdrisTheDragon/AdventOfCode2018/utils"
)

func main() {
	lines := utils.GetLines("../myInput.txt")

	instructions := make(map[rune][]rune)
	parents := make(map[rune]int)

	for _, k := range lines {
		key := rune(k[5])
		value := rune(k[36])
		instructions[key] = append(instructions[key], value)
		parents[value] = parents[value] + 1
	}

	readyTasks := make([]rune, 0)
	for k, _ := range instructions {
		if parents[k] == 0 {
			readyTasks = append(readyTasks, k)
		}
	}

	finishedTasks := make([]rune, 0)
	workersTasks := []rune{'.', '.', '.', '.', '.'}
	workersTimeLeftOnTask := []int{0, 0, 0, 0, 0}
	//workersTasks := []rune{'.', '.'}
	//workersTimeLeftOnTask := []int{0, 0}
	t := 0
	working := 1
	for ; working > 0; t++ {
		working = 0

		for n, _ := range workersTimeLeftOnTask {
			//decrease time left
			if workersTimeLeftOnTask[n] != 0 {
				workersTimeLeftOnTask[n] = workersTimeLeftOnTask[n] - 1
				working = working + 1
			} else {
				//check if more work to do on task
				if workersTasks[n] != '.' {
					finishedTask := workersTasks[n]
					workersTasks[n] = '.'
					
					//check children
					for _, v := range instructions[finishedTask] {
						parents[v] = parents[v] - 1
						if parents[v] == 0 {
							readyTasks = append(readyTasks, v)
						}
					}
				}
			}
		}

		//try add new tasks to workersTimeLeftOnTask
		for ; len(readyTasks) > 0 && working < len(workersTimeLeftOnTask); {
			temp := make([]rune, len(readyTasks))
			copy(temp, readyTasks)
			sort.Sort(utils.Runes(temp))
			x := temp[0]
			for i := 0; i < len(readyTasks); i++ {
				if readyTasks[i] == x {
					readyTasks = append(readyTasks[:i], readyTasks[i+1:]...)
				}
			}
			finishedTasks = append(finishedTasks, x)
			for n, _ := range workersTimeLeftOnTask {
				if workersTasks[n] == '.' {
					workersTasks[n] = x
					workersTimeLeftOnTask[n] = int(x) - 5
					//workersTimeLeftOnTask[n] = int(x) - 'A'
					working = working + 1
					break;
				}
			}
		}
		
		fmt.Print(t," ")
		for n, _ := range workersTimeLeftOnTask {
			fmt.Print("{",workersTimeLeftOnTask[n]," ",string(workersTasks[n]),"} ")
		}
		fmt.Print(string(finishedTasks))
		fmt.Println()
	}
	
	fmt.Println(t-1)
}