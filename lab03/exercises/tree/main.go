package main

import "fmt"

type BinaryTree struct {
	// try to remember what a binary tree needed to contain
	// you can use int values
}

func newBinaryTree(root int) *BinaryTree {
 	return &BinaryTree{}
}

/* A tree should look like this
   2
  / \
 1   3
*/

func main() {
	t := newBinaryTree(3)
	// TODO uncomment these lines after implementing insert
	// t.insert(2)
	// t.insert(3)
	fmt.Println(t)
}
