package main

import "fmt"

// BinaryTreeNode to store node in tree
type BinaryTreeNode struct {
	Left    *BinaryTreeNode
	Right   *BinaryTreeNode
	Element int
}

// BinarySearchTree to store integers in a BST
type BinarySearchTree struct {
	root     *BinaryTreeNode
	numItems int
}

func (bst *BinarySearchTree) insert(elems ...int) {
	for _, v := range elems {
		node := &BinaryTreeNode{nil, nil, v}
		if bst.root == nil {
			bst.root = node
		} else {
			bst.insertHelper(bst.root, node)
		}
		bst.numItems++
	}
}

func (bst *BinarySearchTree) insertHelper(subRoot *BinaryTreeNode, node *BinaryTreeNode) {
	if node.Element <= subRoot.Element {
		if subRoot.Left == nil {
			subRoot.Left = node
		} else {
			bst.insertHelper(subRoot.Left, node)
		}
	} else {
		if subRoot.Right == nil {
			subRoot.Right = node
		} else {
			bst.insertHelper(subRoot.Right, node)
		}
	}
}

func (bst *BinarySearchTree) traversal(t string) []int {
	results := make([]int, 0, bst.numItems)
	switch t {
	case "in":
		bst.inOrderTraversalHelper(bst.root, &results)
	case "pre":
		bst.preOrderTraversalHelper(bst.root, &results)
	case "post":
		bst.postOrderTraversalHelper(bst.root, &results)
	}
	return results
}

func (bst *BinarySearchTree) inOrderTraversal() []int {
	return bst.traversal("in")
}

func (bst *BinarySearchTree) preOrderTraversal() []int {
	return bst.traversal("pre")
}

func (bst *BinarySearchTree) postOrderTraversal() []int {
	return bst.traversal("post")
}

func (bst *BinarySearchTree) inOrderTraversalHelper(subRoot *BinaryTreeNode, results *[]int) {
	if subRoot == nil {
		return
	}
	bst.inOrderTraversalHelper(subRoot.Left, results)
	*results = append(*results, subRoot.Element)
	bst.inOrderTraversalHelper(subRoot.Right, results)
}

func (bst *BinarySearchTree) preOrderTraversalHelper(subRoot *BinaryTreeNode, results *[]int) {
	if subRoot == nil {
		return
	}
	*results = append(*results, subRoot.Element)
	bst.preOrderTraversalHelper(subRoot.Left, results)
	bst.preOrderTraversalHelper(subRoot.Right, results)
}

func (bst *BinarySearchTree) postOrderTraversalHelper(subRoot *BinaryTreeNode, results *[]int) {
	if subRoot == nil {
		return
	}
	bst.postOrderTraversalHelper(subRoot.Left, results)
	bst.postOrderTraversalHelper(subRoot.Right, results)
	*results = append(*results, subRoot.Element)
}

func (bst *BinarySearchTree) getMax() int {
	return bst.getMaxHelper(bst.root)
}

func (bst *BinarySearchTree) getMaxHelper(subRoot *BinaryTreeNode) int {
	if subRoot.Right == nil {
		return subRoot.Element
	}
	return bst.getMaxHelper(subRoot.Right)
}

func (bst *BinarySearchTree) getMin() int {
	return bst.getMinHelper(bst.root)
}

func (bst *BinarySearchTree) getMinHelper(subRoot *BinaryTreeNode) int {
	if subRoot.Left == nil {
		return subRoot.Element
	}
	return bst.getMinHelper(subRoot.Left)
}

func (bst *BinarySearchTree) contains(num int) bool {
	return bst.containsHelper(bst.root, num)
}

func (bst *BinarySearchTree) containsHelper(subRoot *BinaryTreeNode, num int) bool {
	if subRoot == nil {
		return false
	} else if subRoot.Element == num {
		return true
	} else if subRoot.Element < num {
		return bst.containsHelper(subRoot.Right, num)
	} else {
		return bst.containsHelper(subRoot.Left, num)
	}
}

func main() {
	bst := &BinarySearchTree{}
	numsToSend := []int{8, 3, 4, 15, 12, 7, 13, 9}
	bst.insert(numsToSend...)
	fmt.Printf("Inserted the array: %v\n", numsToSend)
	fmt.Printf("In-Order Traversal: %v\n", bst.traversal("in"))
	fmt.Printf("Pre-Order Traversal: %v\n", bst.traversal("pre"))
	fmt.Printf("Post-Order Traversal: %v\n", bst.traversal("post"))
	fmt.Printf("Max Number: %d\n", bst.getMax())
	fmt.Printf("Min Number: %d\n", bst.getMin())
	fmt.Printf("12 is in BST: %v\n", bst.contains(12))
	fmt.Printf("14 is in BST: %v", bst.contains(14))
}
