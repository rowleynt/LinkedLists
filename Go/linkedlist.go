// Simple linked list using Go

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

func (list *LinkedList[T]) append (value T) {
    newNode := &Node[T]{ value: value }
    if list.tail == nil {
        list.tail = list.head
    }
    if list.head == nil {
        list.head = newNode
    } else {
        list.tail.next = newNode
    }
    list.tail = newNode
}

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
    if list.length() <= 0 {/* do nothing */}
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
    currNode := list.head
    outStr := ""
    for currNode != nil {
        outStr += fmt.Sprintf("[%v] -> ", currNode.value)
        currNode = currNode.next
    }
    return outStr[:(len(outStr) - 4)]
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

func main() {
    mylist := new(LinkedList[int])
    mylist.append(4)
    mylist.append(5)
    mylist.append(6)
    mylist.append(8)
    
    mylist.prepend(3)
    mylist.prepend(2)
    mylist.prepend(1)
    mylist.prepend(0)
    
    mylist.insert(7, 7)
    
    mylist.pop(mylist.length() - 1)
    mylist.pop(0)
    mylist.pop(5)
    
    x := mylist.get(1)
    
    fmt.Println(x)
    fmt.Println(mylist)
    fmt.Println(mylist.length())
    
    // final TODO: somehow make the list iterable
}
