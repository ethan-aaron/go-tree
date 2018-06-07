package main

import "fmt"

type node struct {
	Next  *node
	Child *node
	Data  string
}

func main() {
	n := node{Child: &node{Child: &node{Child: &node{Data: "inner"}, Data: "mid 2"}, Data: "mid 1"}, Data: "outer"}
	printTree(&n)
}

func printTreeR(node *node, depth int) {
	for node != nil {
		if node.Child != nil {
			for i := 0; i < (depth * 3); i++ {
				fmt.Printf(" ")
			}
			fmt.Printf("{\n")
			printTreeR(node.Child, depth+1)
			for i := 0; i < (depth * 3); i++ {
				fmt.Printf(" ")
			}
			fmt.Printf("{\n")

			for i := 0; i < (depth * 3); i++ {
				fmt.Printf(" ")
			}
			fmt.Printf("%s\n", node.Data)

			node = node.Next
		} else {
			for i := 0; i < (depth * 3); i++ {
				fmt.Printf(" ")
			}
			fmt.Printf("%s\n", node.Data)
			node = node.Next
		}
	}
}

func printTree(root *node) {
	printTreeR(root, 0)
}
