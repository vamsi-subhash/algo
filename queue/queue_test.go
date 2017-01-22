package queue

import "testing"

func TestAddAndRemoveToQueue(t *testing.T) {
	q := NewQueue(5)
	q.display()
	elements := []int{1, 2, 3, 4, 5}
	for _, elem := range elements {
		err := q.add(elem)
		if err != nil {
			t.Fatalf("Error in adding elements to queue", err)
		}

		frontVal, err := q.remove()
		if err != nil || frontVal != elem {
			t.Fatalf("Inserting %d into a queue returned %d instead of %d\n", elem, frontVal, elem)
		}
	}

	q.display()
	if q.size() != 0 {
		t.Fatalf("Queue is not empty after removing everything. Size: %d\n", q.size())
	}
}
