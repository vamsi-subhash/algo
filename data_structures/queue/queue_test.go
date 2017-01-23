package queue

import (
	"fmt"
	"testing"
)

func TestAddAndRemoveToQueue(t *testing.T) {
	q := NewQueue(5)
	q.Display()
	elements := []int{1, 2, 3, 4, 5}
	for _, elem := range elements {
		err := q.Add(elem)
		if err != nil {
			t.Fatalf("Error in adding elements to queue", err)
		}
	}

	for _, elem := range elements {
		frontVal, err := q.Remove()
		if err != nil || frontVal != elem {
			t.Fatalf("Inserting %d into a queue returned %d instead of %d\n", elem, frontVal, elem)
		}
		fmt.Printf("Dequeued %d\n", frontVal)

	}

	q.Display()
	if q.Size() != 0 {
		t.Fatalf("Queue is not empty after removing everything. Size: %d\n", q.Size())
	}
}
