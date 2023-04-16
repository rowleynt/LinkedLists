// Simple linked list with JavaScript

class Node {
  constructor(value) {
    this.value = value
    this.nextNode = null
  }
}

class LinkedList {
  constructor() {
    this.head = new Node(null)
    this.tail = this.head
  }
  append(value) {
    let nextNode = new Node(value)
    let temp = this.tail
    temp.nextNode = nextNode
    this.tail = nextNode
  }
  prepend(value) {
    let newFirst = new Node(value)
    let temp = this.head.nextNode
    this.head.nextNode = newFirst
    newFirst.nextNode = temp
  }
  insert(value, index) {
    let newNode = new Node(value)
    let currNode = this.head.nextNode
    let prevNode = null
    let count = 0
    while (count < index) {
      prevNode = currNode
      currNode = currNode.nextNode
      count += 1
    }
    prevNode.nextNode = newNode
    newNode.nextNode = currNode
  }
  get(index) {
    let count = 0
    let currNode = this.head.nextNode
    while (count < index) {
      currNode = currNode.nextNode
      count += 1
    }
    return currNode.value
  }
  pop(index=-1) {
    if (index == 0) {
      let currNode = this.head.nextNode
      this.head.nextNode = currNode.nextNode
    }
    else if (index > 0) {
      let count = 0
      let currNode = this.head.nextNode
      let prevNode = null
      let nextNode = null
      while (count < index) {
        prevNode = currNode
        currNode = currNode.nextNode
        nextNode = currNode.nextNode
        count += 1
      }
      prevNode.nextNode = nextNode
    }
    else if (index == -1) {
      let currNode = this.head
      let prevNode = null
      while (currNode.nextNode) {
        prevNode = currNode
        currNode = currNode.nextNode
      }
      prevNode.nextNode = null
    }
    else { void(0) }
  }
  length() {
    let len = 0
    let currNode = this.head
    while (currNode.nextNode) {
      len += 1
      currNode = currNode.nextNode
    }
    return len
  }
  toString() {
    let currNode = this.head.nextNode
    let out = ""
    while (currNode) {
      out += "[" + String(currNode.value) + "] -> "
      currNode = currNode.nextNode
    }
    return out.slice(0, -4) // drop trailing arrow
  }
  [Symbol.iterator]() {
    let currNode = this.head
    return {
      next() {
        if (currNode.nextNode) {
          currNode = currNode.nextNode
          return { value: currNode.value, done: false }
        } else { 
          return { done: true }
        }
      }
    }
  }
}

mylist = new LinkedList()

mylist.append(4)
mylist.append(5)
mylist.append(6)
mylist.append(8)

mylist.prepend(3)
mylist.prepend(2)
mylist.prepend(1)
mylist.prepend(0)

mylist.insert(7, 7)

mylist.pop()
mylist.pop(0)
mylist.pop(5)

x = mylist.get(1)
console.log(x)

console.log(mylist.toString())
console.log(mylist.length())

for (let item of mylist) {
  console.log(item)
}