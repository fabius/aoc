package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var day = 7

type Node struct {
	Name     string
	Size     int
	Children map[string]*Node
	Parent   *Node
}

func (n *Node) updateDirSize(size int) {
	n.Size += size
	if n.Parent != nil {
		n.Parent.updateDirSize(size)
	}
}

func parseInput(f string) *Node {
	input, err := os.ReadFile(f)
	if err != nil {
		log.Fatalln(err)
	}
	tree := Node{Name: "/"}
	currentDir := &tree
	for _, l := range strings.Split(string(input), "\n") {
		if l == "" || l == "$ ls" {
			continue
		}
		if l == "$ cd /" {
			currentDir = &tree
		} else if l == "$ cd .." {
			currentDir = currentDir.Parent
		} else if strings.HasPrefix(l, "$ cd") {
			name := l[5:]
			currentDir = currentDir.Children[name]
		} else {
			nodeLine := strings.Split(l, " ")
			name := nodeLine[1]
			size, _ := strconv.Atoi(nodeLine[0])
			currentDir.updateDirSize(size)
			if currentDir.Children == nil {
				currentDir.Children = map[string]*Node{}
			}
			currentDir.Children[name] = &Node{
				Name:   name,
				Size:   size,
				Parent: currentDir,
			}
		}
	}
	return &tree
}

func (n *Node) findDirsSmallerThan(size int) []Node {
	nodes := []Node{}
	for _, node := range n.Children {
		if node.Children == nil {
			continue
		}
		if node.Size <= size {
			nodes = append(nodes, *node)
		}
		recNodes := node.findDirsSmallerThan(size)
		for _, r := range recNodes {
			nodes = append(nodes, r)
		}
	}
	return nodes
}

func (n *Node) findPossibleDirsToFree(nodes *[]Node, minSize int) {
	if n.Size >= minSize {
		*nodes = append(*nodes, *n)
	}
	for _, c := range n.Children {
		if c.Children == nil {
			continue
		}
		c.findPossibleDirsToFree(nodes, minSize)
	}
}

func (n *Node) smallestDirToFree(minSize int) int {
	possibleNodes := []Node{}
	n.findPossibleDirsToFree(&possibleNodes, minSize)
	sort.Slice(possibleNodes, func(i, j int) bool {
		return possibleNodes[i].Size < possibleNodes[j].Size
	})
	return possibleNodes[0].Size
}

// sum of dirs which are at most 100000
func solveA(f string) int {
	tree := parseInput(f)
	nodes := tree.findDirsSmallerThan(100000)
	sum := 0
	for _, n := range nodes {
		sum += n.Size
	}
	return sum
}

func solveB(f string) int {
	avail := 70000000
	shouldBeFree := 30000000
	tree := parseInput(f)
	unused := avail - tree.Size
	minSize := shouldBeFree - unused
	return tree.smallestDirToFree(minSize)
}

func main() {
	in := fmt.Sprintf("data/%d.txt", day)
	fmt.Printf("A: %d\n", solveA(in))
	fmt.Printf("B: %d\n", solveB(in))
}
