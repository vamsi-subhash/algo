package tree

import (
	"reflect"
	"testing"
)

func TestBinarySearchTree_InsertAndInorderTraversal(t *testing.T) {
	rootNode := IntTreeNode{data: 4}
	bstree := NewBinarySearchTree(&rootNode)

	elements := []int{4, 2, 6, 1, 3, 5, 7}
	for _, elem := range elements[1:] {
		node := &IntTreeNode{data: elem}
		err := bstree.add(node)
		if err != nil {
			t.Fatal(err)
		}
	}

	// if bstree.size() != len(elements) {
	// 	t.Fatalf("Size of bstree should be same (%d) as the no of elements inserted: %d\n", bstree.size(), len(elements))
	// }

	inOrderElements := bstree.inOrder()
	if !reflect.DeepEqual(elements, inOrderElements) {
		t.Fatalf("Inorder traversal of bstree: %+v is not the same as sorted elements: %+v\n", inOrderElements, elements)
	}

}
