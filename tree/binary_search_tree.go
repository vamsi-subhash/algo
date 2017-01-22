package tree

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
	return []int{}
}
