package stack

import (
	"errors"
	"fmt"
)

type Stack struct {
	Size int
	arr  []int
	top  int
}

func NewStack(size int) *Stack {
	stack := Stack{Size: size, top: -1}
	stack.arr = make([]int, size)
	return &stack
}

func (this *Stack) push(elem int) error {
	if this.isFull() {
		return errors.New("Stack is full")
	}
	this.top += 1
	this.arr[this.top] = elem
	return nil
}

func (this *Stack) pop() (int, error) {
	if this.isEmpty() {
		return -1, errors.New("Stack is empty")
	}
	elem := this.arr[this.top]
	this.top -= 1
	return elem, nil
}

func (this *Stack) isFull() bool {
	return this.top == (this.Size - 1)
}

func (this *Stack) isEmpty() bool {
	return this.top == -1
}

func (this *Stack) size() int {
	return this.top + 1
}

func (this *Stack) display() {
	top := this.top
	if top == -1 {
		top = 0
	}
	fmt.Printf("Stack: %+v\n", this.arr[0:top])
}
