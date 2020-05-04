package stack 

import (
	"sync"
	"errors"
	"fmt"
)

type (
	Stack struct {
		top *node
		length int
		lock sync.Mutex
	}
	node struct {
		value interface{}
		prev *node
	}
)

func NewStack() *Stack {
	return &Stack{nil,0,sync.Mutex{}}
}

func (stack *Stack) Pop() (value interface{},err error) {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	if stack.length < 1 {
		err = errors.New("The Stack is null")
	}
	stack.length = stack.length - 1
	value = stack.top.value
	err = nil
	stack.top = stack.top.prev
	return
}


func (stack *Stack) Push(value interface{}) {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	node := &node{value, stack.top}
	stack.top = node
	stack.length = stack.length + 1
}

func (stack *Stack) Set(idx int, value interface{}) (err error){
	stack.lock.Lock()
	defer stack.lock.Unlock()
	if idx > stack.length {
		return errors.New("Index is out of length of stack")
	}
	node := &node{nil, nil}
	for i:=0;i<=idx;i++ {
		node = stack.top
		stack.top = stack.top.prev
	}
	node.value = value
	return nil
}

func (stack *Stack) Get(idx int) (value interface{}, err error) {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	if idx > stack.length {
		return nil,errors.New("Index is out of length of stack")
	}

	node := &node{nil, nil}
	for i:=0;i<=idx;i++ {
		node = stack.top
		stack.top = stack.top.prev
	}
	return node.value, nil
}

func (stack *Stack) IsEmpty() bool { return stack.length == 0}

func (stack *Stack) Print() (err error) {
	if stack.IsEmpty() {
		return errors.New("The stack is empty, print nothing")
	}

	for i:=0;i<stack.length;i++ {
		fmt.Printf("Location of Stack in %d is %v\n", i, stack.top.value)
		stack.top = stack.top.prev
	}
	return nil
}

func (stack *Stack) Swap(other *Stack) (err error) {
	if stack.length != other.length {
		return errors.New("The stack`s length is  is not equal to the source stack`s length \n")
	}

	stack.lock.Lock()
	other.lock.Lock()
	defer stack.lock.Unlock()
	defer other.lock.Unlock()

	stack.top, other.top = other.top, stack.top
	return nil
}
