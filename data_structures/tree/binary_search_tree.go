package tree

import (
	"fmt"

	"github.com/vamsi-subhash/algo/queue"
)

type TreeNode struct {
	left  *TreeNode
	Data  int
	right *TreeNode
}

type BinarySearchTree struct {
	Root *TreeNode
}

func NewBinarySearchTree(rootData int) *BinarySearchTree {
	root := TreeNode{Data: rootData}
	return &BinarySearchTree{Root: &root}
}

func (this *BinarySearchTree) add(elem int) error {
	fmt.Printf("Adding %v to queue\n", elem)

	newNode := TreeNode{Data: elem}
	curr := this.Root

	for curr != nil {
		if elem < curr.Data {
			if curr.left == nil {
				curr.left = &newNode
				return nil
			}
			curr = curr.left
		} else {
			if curr.right == nil {
				curr.right = &newNode
				return nil
			}
			curr = curr.right
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

		node, ok := elem.(TreeNode)
		if ok {
			fmt.Printf("Dequeued %+v\n", node.Data)
			inOrderSlice = append(inOrderSlice, node.Data)

			if node.left != nil {
				q.Add(*node.left)
			}
			if node.right != nil {
				q.Add(*node.right)
			}
		} else {
			fmt.Println("Error in typecasting elem to TreeNode type")
			break
		}

	}
	return inOrderSlice
}
