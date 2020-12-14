package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Node struct {
	parents  map[*Node]int64
	children map[*Node]int64
	name     string
}

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var bagMap = make(map[string]*Node, 0)

	for scanner.Scan() {
		var line = scanner.Text()
		var bagTypeRegex = regexp.MustCompile(`^(\w+ \w+) bags contain`)
		var otherBagsRegex = regexp.MustCompile(`(\d) (\w+ \w+)`)

		var bagType = bagTypeRegex.FindStringSubmatch(line)
		var childrenBags = otherBagsRegex.FindAllStringSubmatch(line, -1)
		var bagName = strings.TrimSpace(bagType[1])
		// fmt.Println(bagName)
		var node = getNode(bagMap, bagName)

		if len(childrenBags) != 0 {

			for _, x := range childrenBags {
				count, _ := strconv.ParseInt(x[1], 10, 64)
				name := strings.TrimSpace(x[2])
				// fmt.Println(count, name)
				childNode := getNode(bagMap, name)
				childNode.parents[node] = count

				node.children[childNode] = count
				bagMap[name] = childNode

			}
		}
		bagMap[bagName] = node
	}

	shinyGold := getNode(bagMap, "shiny gold")

	// validParents := map[string]int{}
	// traverseParents(shinyGold, validParents)
	// fmt.Println(len(validParents))

	fmt.Println(traverseChildrenCounts(shinyGold) - 1)

}

func traverseParents(node *Node, parents map[string]int) {
	for parent, _ := range node.parents {
		parents[parent.name] = 1
		traverseParents(parent, parents)
	}
}

func traverseChildrenCounts(node *Node) int64 {
	var extra int64 = 1
	for child, val := range node.children {
		extra += val * traverseChildrenCounts(child)
	}
	return extra
}

func getNode(bagMap map[string]*Node, name string) *Node {
	if node, exists := bagMap[name]; exists {
		return node
	} else {
		node := newNode(name)
		return node
	}
}

func newNode(name string) *Node {
	return &Node{
		parents:  map[*Node]int64{},
		children: map[*Node]int64{},
		name:     name,
	}
}
