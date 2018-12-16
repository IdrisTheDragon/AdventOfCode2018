package main

import (
	"fmt"
	"github.com/IdrisTheDragon/AdventOfCode2018/utils"
)


func main() {
	lines := utils.GetLines("../myInput.txt")

	opcodes := []OP{
		{name:"addr",action:'+',a:'r',b:'r'},
		{name:"addi",action:'+',a:'r',b:'v'},
		{name:"mulr",action:'*',a:'r',b:'r'},
		{name:"muli",action:'*',a:'r',b:'v'},
		{name:"banr",action:'&',a:'r',b:'r'},
		{name:"bani",action:'&',a:'r',b:'v'},
		{name:"borr",action:'|',a:'r',b:'r'},
		{name:"bori",action:'|',a:'r',b:'v'},
		{name:"setr",action:'a',a:'r',b:'r'},
		{name:"seti",action:'a',a:'v',b:'r'},
		{name:"gtir",action:'>',a:'v',b:'r'},
		{name:"gtri",action:'>',a:'r',b:'v'},
		{name:"gtrr",action:'>',a:'r',b:'r'},
		{name:"eqir",action:'=',a:'v',b:'r'},
		{name:"eqri",action:'=',a:'r',b:'v'},
		{name:"eqir",action:'=',a:'r',b:'r'},
	}

	sum:=0
	var lineCount int
	for lineCount< len(lines) {
		if len(lines[lineCount]) > 0 && lines[lineCount][0] == 'B'{
			split :=  utils.RegSplit(lines[lineCount],"[^0-9]+")
			registers := []int{
				utils.Str2Int(split[1]),
				utils.Str2Int(split[2]),
				utils.Str2Int(split[3]),
				utils.Str2Int(split[4]),
			}
			split =  utils.RegSplit(lines[lineCount+1],"[^0-9]+")
			instruction := []byte{
				byte(utils.Str2Int(split[0])),
				byte(utils.Str2Int(split[1])),
				byte(utils.Str2Int(split[2])),
				byte(utils.Str2Int(split[3])),
			}
			split =  utils.RegSplit(lines[lineCount+2],"[^0-9]+")
			result := []int{
				utils.Str2Int(split[1]),
				utils.Str2Int(split[2]),
				utils.Str2Int(split[3]),
				utils.Str2Int(split[4]),
			}
			tempSum := testCode(registers,result, instruction, opcodes)

			if tempSum >= 3 {
				sum++
			}

			lineCount = lineCount + 4 
		} else {
			break
		}
	}

	fmt.Println(sum)
	//fmt.Println(opcodes)

	//part 2 below here

	orderedOpCodes := make(map[byte]*OP,0)

	for len(orderedOpCodes) < 16 {
		for i := range opcodes {
			
			if len(opcodes[i].matchCount) == 1 {
				c := opcodes[i].matchCount[0]
				orderedOpCodes[c] = &opcodes[i]
				for j := range opcodes {
					remove(&opcodes[j],c)
				}
			}
		}
	}

	//for k,v := range orderedOpCodes {
	//	fmt.Println(k,*v)
	//}
	
	lineCount = lineCount + 2

	r := make([]int,4)

	for ; lineCount < len(lines); lineCount++ {
		split :=  utils.RegSplit(lines[lineCount],"[^0-9]+")
		instruction := []byte{
			byte(utils.Str2Int(split[0])),
			byte(utils.Str2Int(split[1])),
			byte(utils.Str2Int(split[2])),
			byte(utils.Str2Int(split[3])),
		}

		r = runOp(*orderedOpCodes[instruction[0]],r,instruction)
	}

	fmt.Println(r)
	
}

func remove(op *OP, c byte) {
	i:= -1
	for j,v := range op.matchCount {
		if c == v {
			i = j
		}
	}
	if i != -1 {
		op.matchCount = append(op.matchCount[:i],op.matchCount[i+1:]...)
	}
}

func add(op *OP, c byte){
	for _,v := range op.matchCount {
		if v == c {
			return
		}
	}
	op.matchCount = append(op.matchCount,c)
}

func testCode( registers,result []int, instruction []byte, opcodes []OP ) int {
	sum := 0
	for i := range opcodes {
		if match(result,runOp(opcodes[i],registers,instruction)) {
			add(&opcodes[i],instruction[0])
			sum++
		}
	}
	return sum
}

func match(r, c[]int) bool {
	if len(r) != len(c) {
		return false
	}
	for i := range r {
		if r[i] != c[i] {
			return false
		}
	}
	return true
}

func runOp(op OP,registers []int, instruction []byte) []int {
	registerCP := make([]int,4)
	copy(registerCP,registers)
	var A,B int
	if(op.a == 'r'){
		A = registerCP[instruction[1]]
	} else {
		A = int(instruction[1])
	}
	if(op.b == 'r'){
		B = registerCP[instruction[2]]
	} else {
		B = int(instruction[2])
	}
	switch(op.action){
	case '+':
		registerCP[instruction[3]] = A + B
		break
	case '*':
		registerCP[instruction[3]] = A * B
		break
	case '&':
		registerCP[instruction[3]] = A & B
		break
	case '|':
		registerCP[instruction[3]] = A | B
		break
	case 'a':
		registerCP[instruction[3]] = A
		break
	case '>':
		if A > B {
			registerCP[instruction[3]] = 1
		} else {
			registerCP[instruction[3]] = 0
		}
		break
	case '=':
		if A == B {
			registerCP[instruction[3]] = 1
		} else {
			registerCP[instruction[3]] = 0
		}
		break
	default:
		fmt.Println("not valid instruction")
	}
	return registerCP
}


type OP struct {
	a rune
	b rune
	action rune
	name string
	matchCount []byte
}


