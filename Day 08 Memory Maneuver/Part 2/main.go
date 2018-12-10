package main

import (
	"fmt"
	"github.com/IdrisTheDragon/AdventOfCode2018/utils"
)

func main() {
	lines := utils.GetLines("../myInput.txt")
	line := lines[0]
	split := utils.RegSplit(line," ")

	node := getNode(0,split);


	fmt.Println(node)

	fmt.Println(sumMeta(node))
	fmt.Println(sumNodeValue(node))
}

func sumNodeValue(node Node) int {
	sum := 0
	if(node.numChildNodes == 0){
		sum = getSum(node.metaData)
	} else {
		for _,v:= range node.metaData {
			if(v-1 < node.numChildNodes && v > 0){
				sum = sum + sumNodeValue(node.childNodes[v-1])
			}
		}
	}
 	return sum
}

func sumMeta(node Node) int {
	sum := 0
	for _,v := range node.childNodes {
		sum = sum + sumMeta(v)
	}
	sum = sum + getSum(node.metaData)
	return sum
}

func getSum(list []int) int {
	sum := 0
	for _,v := range list {
		sum = sum + v
	}
	return sum
}

func getNode(index int, split []string) Node {
	node := Node{index: index, numChildNodes: utils.Str2Int(split[index]) , numMetaData : utils.Str2Int(split[index+1])}
	//fmt.Println(node)
	offset := node.index + 2 

	for i := 0; i < node.numChildNodes ; i++ {
		childNode := getNode( offset,split)
		node.childNodes = append(node.childNodes, childNode)
		offset = offset + getLength(childNode)
	}

	for i := 0; i < node.numMetaData ; i++ {
		node.metaData = append(node.metaData,utils.Str2Int(split[offset + i]))
	}
	return node
}

func getLength(node Node) int {
	length := 2
	for i := 0; i < node.numChildNodes ; i++ {
		length = length + getLength(node.childNodes[i])
	}
	length = length + node.numMetaData
	return length
}


type Node struct {
	index int
	numChildNodes int
	childNodes []Node
	numMetaData int
	metaData []int
}


