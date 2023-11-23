package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const totalDiskSpace = 70000000
const diskSpaceRequiredForUpdate = 30000000

func main() {
	// keep reference to current path with a slice/stack push & pop from it
	// for each dir sum files
	// but I need to some how figure out what it contains...

	// use a list

	// make a tree

	file, err := os.Open("./day7/input.txt")
	if err != nil {
		fmt.Println("error:", err)
	}
	defer file.Close()
	tree := loadDirTree(file)

	totalSize, problemNodesSize, problemNodes := sumDirSizes(&tree)
	fmt.Println(totalSize, problemNodesSize, problemNodes)

	spaceLeft := totalDiskSpace - totalSize
	needToFree := spaceLeft - diskSpaceRequiredForUpdate

	if needToFree < 0 {
		// fmt.Println(needToFree * -1)
		sizeOfDir := findSmallestDirGte(&tree, needToFree*-1)
		fmt.Println(sizeOfDir)

	}
}

type FileNode struct {
	name string
	size int
}

type DirNode struct {
	name       string
	parentNode *DirNode
	childFiles []*FileNode
	childDirs  []*DirNode
}

type DirTree struct {
	root *DirNode
}

func loadDirTree(file *os.File) DirTree {
	var tree DirTree
	var curNode *DirNode
	// keep pointer to current dir! & return a pointer to the head/root
	// line :=
	// line :=
	// lines := []string{
	// 	"$ cd /",
	// 	"187585 dgflmqwt.srm",
	// 	"dir gnpd",
	// 	"$ cd gnpd",
	// 	"$ cd ..",
	// 	"$ cd plj",
	// }
	scanner := bufio.NewScanner(file)
	// for _, l := range lines {
	for scanner.Scan() {
		// line := l
		line := scanner.Text()
		// fmt.Println("LINE: ", l)
		if strings.HasPrefix(line, "$") {
			command := strings.Split(line, " ")

			switch command[1] {
			case "cd":
				if len(command) != 3 {
					log.Fatal("Invalid CD command")
				}
				switch command[2] {
				case "..":
					curNode = curNode.parentNode
				default:
					newDir := DirNode{
						name: command[2],
						// parentNode: curNode,
						childFiles: []*FileNode{},
						childDirs:  []*DirNode{},
					}

					ok := tree.addDirNode(&newDir, curNode)
					if ok {
						curNode = &newDir
						// fmt.Println("NEW NODE: ", curNode)
					} else {
						fmt.Println("Couldn't add node")
					}
				}
			case "ls":
				// Do nothing it doesn't matter actually
			}
		} else {
			// is tree node
			nodeData := strings.Split(line, " ")

			if len(nodeData) != 2 {
				fmt.Println("Invalid node ", nodeData)
			} else if nodeData[0] == "dir" {
				// we only load dirs when we cd into them so this does nothing
				// fmt.Println("Its a dir")
				// newDir := DirNode{
				// 	name: nodeData[1],
				// 	// parentNode: curNode,
				// 	childFiles: []*FileNode{},
				// 	childDirs:  []*DirNode{},
				// }

				// ok := tree.addDirNode(&newDir, curNode)
				// fmt.Println(ok)
				// fmt.Println(tree)
				// fmt.Println(tree.root)
			} else {
				size, err := strconv.Atoi(nodeData[0])
				if err != nil {
					fmt.Println("Invalid console result ", nodeData)
					// BREAK!
				}
				newFile := FileNode{
					name: nodeData[1],
					size: size,
				}
				// fmt.Println(newFile)
				tree.addFileNode(&newFile, curNode)
			}
		}
	}

	// tree.printTree()

	return tree
}

// WARNING. DOES NO DUPLICATE CHECKING WHICH CAN BE A BIG PROBLEM
func (tree *DirTree) addDirNode(node *DirNode, parentNode *DirNode) bool {
	if tree.root != nil && parentNode == nil {
		fmt.Println("ERROR LOADING DIR TO TREE")
		return false
	}
	// else add node to parentNode
	// warning set parentNode property appropriately
	if tree.root == nil {
		tree.root = node
	} else {
		node.parentNode = parentNode
		parentNode.childDirs = append(parentNode.childDirs, node)
	}
	return true
}

func (tree *DirTree) addFileNode(node *FileNode, parentNode *DirNode) bool {
	if parentNode == nil {
		fmt.Println("ERROR LOADING FILE TO TREE")
		return false
	}

	parentNode.childFiles = append(parentNode.childFiles, node)

	return true

}

func findSmallestDirGte(tree *DirTree, minReqSize int) (dirSize int) {
	recSumDirs(tree.root, checkDirIsSmallestFactory(&dirSize, minReqSize))

	return dirSize
}

func checkDirIsSmallestFactory(size *int, minRequiredSize int) func(node *DirNode, dirSize int) {

	return func(node *DirNode, dirSize int) {
		if dirSize > minRequiredSize {
			if *size == 0 || dirSize < *size {
				*size = dirSize
			}
		}
	}
}

func sumDirSizes(tree *DirTree) (totalSize int, problemNodesSize int, problemNodes []string) {
	return recSumDirs(tree.root, problemNodeFactoryFinder(&problemNodesSize, &problemNodes)), problemNodesSize, problemNodes
}

func problemNodeFactoryFinder(problemNodesSize *int, problemNodes *[]string) func(node *DirNode, dirSize int) {
	return func(node *DirNode, dirSize int) {
		if dirSize < 100000 {
			// if dirSize < 4 {
			*problemNodes = append(*problemNodes, node.name)
			*problemNodesSize += dirSize
		}
	}
}

func recSumDirs(node *DirNode, callBack func(node *DirNode, dirSize int)) int {
	// base case
	if node == nil {
		return 0
	}
	// pre
	dirSize := 0

	for _, file := range node.childFiles {
		dirSize += file.size
	}
	// recurse in loop while summing
	for _, dir := range node.childDirs {
		dirSize += recSumDirs(dir, callBack)
	}
	callBack(node, dirSize)
	// post
	return dirSize
}

// func (tree *DirTree) traverseAndDo(callBack func(node *DirNode)) {

// 	traverserWithFunction(tree.root, callBack)
// }

// func traverserWithFunction(node *DirNode, do func(node *DirNode)) {
// 	if node == nil {
// 		return
// 	}

// 	do(node)

// 	for _, node := range node.childDirs {
// 		traverserWithFunction(node, do)
// 	}
// }

func (tree *DirTree) printTree() {
	printNodeRec(tree.root)
}

func printNodeRec(node *DirNode) {

	if node == nil {
		return
	}

	// use the height to print spaces?
	fmt.Println("-----")
	fmt.Println("DIR: ", node.name)
	fmt.Println("Parent: ", node.parentNode)

	for _, file := range node.childFiles {
		fmt.Println("file: ", file.name)
	}

	fmt.Println("SUBDIRS: ", node.childDirs)
	for _, node := range node.childDirs {
		printNodeRec(node)
	}
}
