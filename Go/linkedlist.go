// Simple linked list using Go

// Very WIP

package main

import (
    "fmt"
)

type Node[T any] struct {
    value T
    nextNode *Node[T]
}

type LinkedList[T any] struct {
    head, tail *Node[T]
}

func (list *LinkedList[T]) prepend (value T) {
    // head node is no longer a dummy node
    // instead, head node is first node in list
    newNode := &Node[T]{ value: value }
    if list.head == nil {
        list.head = newNode
    } else {
        currNode := list.head
        list.head = newNode
        newNode.nextNode = currNode
    }
}

// this is borked, sleep on it :)
func (list *LinkedList[T]) append (value T) {
    newNode := &Node[T]{ value: value }
    currNode := list.head
    for (currNode.nextNode != nil) {
        fmt.Println("node")
        currNode = currNode.nextNode
    }
    currNode.nextNode = newNode
}

// append: tail isn't used anymore, need new way to add to end

func main() {
    mylist := new(LinkedList[int])
    mylist.prepend(10)
    mylist.prepend(11)
    mylist.append(15)
    fmt.Println(mylist.head.nextNode.value)
    currNode := mylist.head
    for (currNode.nextNode != nil) {
        fmt.Println(currNode.value)
        currNode = currNode.nextNode
    }
}
