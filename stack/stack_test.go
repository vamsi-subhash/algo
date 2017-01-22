package stack

import "testing"

func TestStackPushPop(t *testing.T) {
	stack := NewStack(5)

	elements := []int{1, 2, 3, 4, 5}
	for _, elem := range elements {
		err := stack.push(elem)
		if err != nil {
			t.Fatal(err)
		}

		poppedElem, err := stack.pop()
		if err != nil {
			t.Fatal(err)
		}
		if poppedElem != elem {
			t.Fatalf("Popping stack after pushing %d should return %d instead of %d\n", elem, elem, poppedElem)
		}
	}

	stack.display()
	if stack.size() != 0 {
		t.Fatalf("Stack should be empty after popping all elements\n")
	}
}
