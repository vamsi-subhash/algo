package tree

import (
	"fmt"

	"github.com/vamsi-subhash/algo/data_structures/queue"
)

// type TreeNode interface {
// 	GetLeft() *TreeNode
// 	GetRight() *TreeNode
// 	SetLeft(child *TreeNode)
// 	SetRight(child *TreeNode)
// 	LessThan(other *TreeNode) bool
// }

type IntTreeNode struct {
	left  *IntTreeNode
	data  int
	right *IntTreeNode
}

func (node *IntTreeNode) GetLeft() *IntTreeNode {
	return node.left
}

func (node *IntTreeNode) GetRight() *IntTreeNode {
	return node.right
}

func (node *IntTreeNode) SetLeft(child *IntTreeNode) {
	node.left = child
}

func (node *IntTreeNode) SetRight(child *IntTreeNode) {
	node.right = child
}

func (node *IntTreeNode) LessThan(other *IntTreeNode) bool {
	// otherNode := other.(*IntTreeNode)
	return node.data < other.data
}

type BinarySearchTree struct {
	Root *IntTreeNode
}

func NewBinarySearchTree(root *IntTreeNode) *BinarySearchTree {
	return &BinarySearchTree{Root: root}
}

func (this *BinarySearchTree) add(newNode *IntTreeNode) error {
	fmt.Printf("Adding %v to queue\n", newNode)

	curr := this.Root
	for curr != nil {
		if newNode.LessThan(curr) {
			if curr.GetLeft() == nil {
				curr.SetLeft(newNode)
				return nil
			}
			curr = curr.GetLeft()
		} else {
			if curr.right == nil {
				curr.SetRight(newNode)
				return nil
			}
			curr = curr.GetRight()
		}
	}

	return nil
}

func (this *BinarySearchTree) remove(elem int) error {
	return nil
}

func (this *BinarySearchTree) size() int {
	return 0
}

func (this *BinarySearchTree) inOrder() []int {
	inOrderSlice := make([]int, 0, 10000)

	fmt.Printf("Root: %+v\n", this.Root)
	q := queue.NewQueue(100)
	q.Add(*this.Root)

	for {
		elem, err := q.Remove()
		if err != nil {
			break
		}

		node, ok := elem.(IntTreeNode)
		if ok {
			fmt.Printf("Dequeued %+v\n", node)
			inOrderSlice = append(inOrderSlice, node.data)

			if node.left != nil {
				q.Add(*node.GetLeft())
			}
			if node.right != nil {
				q.Add(*node.GetRight())
			}
		} else {
			fmt.Println("Error in typecasting elem to TreeNode type")
			break
		}

	}
	return inOrderSlice
}
