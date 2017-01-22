package tree

import (
	"reflect"
	"sort"
	"testing"
)

func TestBinarySearchTree_InsertAndInorderTraversal(t *testing.T) {
	bstree := NewBinarySearchTree(4)

	elements := []int{4, 2, 7, 1, 3, 5, 8}
	for _, elem := range elements[1:] {
		err := bstree.add(elem)
		if err != nil {
			t.Fatal(err)
		}
	}

	if bstree.size() != len(elements) {
		t.Fatalf("Size of bstree should be same (%d) as the no of elements inserted: %d\n", bstree.size(), len(elements))
	}

	inOrderElements := bstree.inOrder()
	sort.Ints(elements)
	if reflect.DeepEqual(elements, inOrderElements) {
		t.Fatalf("Inorder traversal of bstree: %+v is not the same as sorted elements: %+v\n", inOrderElements, elements)
	}

}
