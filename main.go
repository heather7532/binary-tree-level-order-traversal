package main

import (
	"fmt"
)

// struct for binary tree node
type Node struct {
	Left  *Node
	Right *Node
	Value int
}

func main() {

	// create test data
	root := &Node{Value: 1}
	root.Left = &Node{Value: 2}
	root.Right = &Node{Value: 3}
	root.Left.Left = &Node{Value: 4}
	root.Right.Left = &Node{Value: 5}
	root.Right.Right = &Node{Value: 6}

	fmt.Println("Level Order Traversal of Binary Tree:")
	levelOrderTraversal(root)
}

func levelOrderTraversal(root *Node) {
	if root == nil {
		return
	}

	currentLevel := []*Node{root}
	for len(currentLevel) > 0 {
		nextLevel := []*Node{}
		for _, node := range currentLevel {
			fmt.Print(node.Value, " ")
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}
		fmt.Println()
		currentLevel = nextLevel
	}
}
