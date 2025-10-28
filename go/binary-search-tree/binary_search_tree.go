package binarysearchtree

type BinarySearchTree struct {
	left  *BinarySearchTree
	data  int
	right *BinarySearchTree
}

// NewBst creates and returns a new BinarySearchTree.
func NewBst(i int) *BinarySearchTree {
	return &BinarySearchTree{data: i}
}

// Insert inserts an int into the BinarySearchTree.
// Inserts happen based on the rules of a binary search tree
func (bst *BinarySearchTree) Insert(i int) {
	if bst == nil {
		return
	}

	if i <= bst.data {
		if bst.left == nil {
			bst.left = NewBst(i)
		} else {
			bst.left.Insert(i)
		}
	} else {
		if bst.right == nil {
			bst.right = NewBst(i)
		} else {
			bst.right.Insert(i)
		}
	}
}

// SortedData returns the ordered contents of BinarySearchTree as an []int.
// The values are in increasing order starting with the lowest int value.
// A BinarySearchTree that has the numbers [1,3,7,5] added will return the
// []int [1,3,5,7].
func (bst *BinarySearchTree) SortedData() []int {
	var result []int
	stack := []*BinarySearchTree{}
	current := bst

	for current != nil || len(stack) > 0 {
		for current != nil {
			stack = append(stack, current)
			current = current.left
		}

		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		result = append(result, current.data)

		current = current.right
	}
	return result
}

// // recursive version
// func (bst *BinarySearchTree) SortedData() []int {
// 	var sorted []int
//
// 	// Linken Teilbaum verarbeiten
// 	if bst.left != nil {
// 		sorted = append(sorted, bst.left.SortedData()...)
// 	}
//
// 	// Aktuellen Knoten hinzufügen
// 	sorted = append(sorted, bst.data)
//
// 	// Rechten Teilbaum verarbeiten
// 	if bst.right != nil {
// 		sorted = append(sorted, bst.right.SortedData()...)
// 	}
//
// 	return sorted
// }
