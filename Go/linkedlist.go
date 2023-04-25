// Simple linked list using Go

// Very WIP

package main

import (
    "fmt"
    "math"
)

type Node[T any] struct {
    value T
    next *Node[T]
}

type LinkedList[T any] struct {
    head, tail *Node[T]
}

// works (?)
func (list *LinkedList[T]) append (value T) {
    newNode := &Node[T]{ value: value }
    if list.tail == nil {
        list.tail = list.head
    } else {
        list.tail.next = newNode
        list.tail = newNode
    }
}

// works
func (list *LinkedList[T]) prepend (value T) {
    newNode := &Node[T]{ value: value }
    if list.head == nil {
        list.head = newNode
    } else {
        tempNode := list.head
        list.head = newNode
        list.head.next = tempNode
    }
    if list.tail == nil {
        list.tail = list.head
    }
}

func (list *LinkedList[T]) insert (value T, index int) {
    prevNode, nextNode, _ := list.walk(index)
    newNode := &Node[T]{ value: value }
    prevNode.next = newNode
    newNode.next = nextNode
}

func (list *LinkedList[T]) pop (index int) {
    if list.length() <= 0 {
        // do nothing
    }
    if index == 0 {
        list.head = list.head.next
    } else {
        prevNode, currNode, _ := list.walk(index)
        prevNode.next = currNode.next
    }
}

func (list *LinkedList[T]) get (index int) T {
    _, currNode, _ := list.walk(index)
    return currNode.value
}

func (list *LinkedList[T]) String() string {
    return "fix me"
}

func (list *LinkedList[T]) length() int {
    _, _, len := list.walk(-1)
    return len
}

func (list *LinkedList[T]) walk (stopIndex int) (*Node[T], *Node[T], int) {
    currNode := list.head
    prevNode := new(Node[T])
    count := 0
    if stopIndex == -1 {
        stopIndex = (int(math.Inf(1)) + 1) * -1
    }
    for currNode != nil && count < stopIndex {
        prevNode = currNode
        currNode = currNode.next
        count++
    }
    return prevNode, currNode, count
}

func (list *LinkedList[T]) output () {
    currNode := list.head
    for currNode != nil {
        fmt.Println(currNode.value)
        currNode = currNode.next
    }
}

func main() {
    mylist := new(LinkedList[int])
    mylist.prepend(6)
    mylist.prepend(5)
    mylist.prepend(4)
    mylist.append(7)
    mylist.insert(8, 2)
    mylist.pop(mylist.length() - 1)
    fmt.Println(mylist.get(0))
    //mylist.output()
}
