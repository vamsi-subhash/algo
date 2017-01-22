package queue

import (
	"errors"
	"fmt"
)

type Queue struct {
	Size  int
	arr   []int
	front int
	rear  int
}

func NewQueue(size int) *Queue {
	q := Queue{Size: size, front: -1, rear: -1}
	q.arr = make([]int, size)
	return &q
}

func (this *Queue) add(item int) error {
	if this.isFull() {
		return errors.New("Queue is full")
	}
	this.rear += 1
	this.arr[this.rear] = item
	return nil
}

func (this *Queue) remove() (int, error) {
	if this.isEmpty() {
		return -1, errors.New("Queue is empty")
	}
	this.front += 1
	elem := this.arr[this.front]
	return elem, nil
}

func (this *Queue) isFull() bool {
	return this.rear == len(this.arr)
}

func (this *Queue) isEmpty() bool {
	return this.front == this.rear
}

func (this *Queue) size() int {
	return this.rear - this.front
}

func (this *Queue) display() {
	front := this.front
	rear := this.rear
	if front == -1 {
		front = 0
	}
	if rear == -1 {
		rear = 0
	}
	fmt.Printf("Queue: %+v\n", this.arr[front:rear])
}
