package main

import "fmt"

type node struct {
	Next  *node
	Child *node
	Data  string
}

func main() {
	n := node{Next: &node{}, Data: "hi"}
	fmt.Println(n.Child)
}

func printTreeR(node *node, depth int) {
	// var i int
	for node != nil {
		if node.Child != nil {
			//
		}
	}
}

func printTree(root *node) {
	printTreeR(root, 0)
}
