package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// what ends up on top
	graph, destination := loadGraph()

}

type Node struct {
	val int
	to  []int
}

type Graph struct {
	head  *Node
	nodes []Node
}

func loadGraph() (graph Graph, destination int) {
	file, err := os.ReadFile("./day12/input.txt")
	if err != nil {
		fmt.Println("error:", err)
	}
	fileContents := strings.Split(string(file), "\n")


	for i, line := range fileContents {
		// i = row
		// j = col
		for j, char := range line {

			node := Node{
				val: int(char),
			}
			if char == 'S' {
				node.val = int('a')
				graph.head = &node
			}
			if char == 'E' {
				node.val = int('z')
				destination = i + j
			}
			// get the elements it points to
			
			node.to = getPointers(i, j, &fileContents)

			// add it to the node
			graph.nodes = append(graph.nodes, node)
		}

	}
	

	return graph, destination 
	// iterate through each char
	// convert to int
	// get all the index of the elements next to it that we can walk to
	// create a node
	// return the head & what we are searching for, index of
}


func getPointers (row int, col int, matrix *[]string) (pointers []int) {

	canMoveUpTo := (*matrix)[row][col] + 1 
	
	if row - 1 >= 0 && (*matrix)[row -1][col] <= canMoveUpTo{
		pointers = append(pointers, matrixToGraphPos(row -1, col) )
	}
	
	if row + 1 < len(*matrix) && (*matrix)[row + 1][col] <= canMoveUpTo{
		pointers = append(pointers, matrixToGraphPos(row + 1, col))
	}
	
	if col - 1 >=  0 && (*matrix)[row][col - 1] <= canMoveUpTo{
		pointers = append(pointers, matrixToGraphPos(row, col -1))
	}

	if col + 1 < len((*matrix)[row]) && (*matrix)[row][col + 1] <= canMoveUpTo{
		pointers = append(pointers, matrixToGraphPos(row, col + 1))
	}
	
	return pointers
}

func matrixToGraphPos (row int, col int, totalCols int) int {
	return row * totalCols + col
}