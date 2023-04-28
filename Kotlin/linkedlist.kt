// Simple linked list with Kotlin

@Suppress("DIVISION_BY_ZERO")
const val INF: Double = 1.0 / 0.0

class Node<T>(var value: T, var next: Node<T>? = null)

class LinkedList<T>(var head: Node<T>? = null, var tail: Node<T>? = null) {
  var length: Int = 0
  init {
    this.head = this.tail
  }
  
  fun prepend(value: T) {
    var newNode = Node(value = value)
    if (this.head == null) {
      this.head = newNode
      this.tail = newNode
      this.length++
    } else {
      var nextNode = this.head
      this.head = newNode
      newNode.next = nextNode
      this.length++
    }
  }
  
  fun append(value: T) {
    var newNode = Node(value = value)
    if (this.head == null) {
      this.head = newNode
    }
    this.tail?.next = newNode
    this.tail = newNode
    this.length++
  }
  
  fun insert(value: T, index: Int) {
    var newNode = Node(value = value)
    var (prevNode, nextNode) = this.walk(index)
    prevNode?.next = newNode
    newNode.next = nextNode
    this.length++
  }
  
  fun pop(index: Int) {
    var (prevNode, currNode) = walk(index)
    prevNode?.next = currNode?.next
    this.length--
  }
  
  fun get(index: Int): T? {
    var (_, currNode) = walk(index)
    return currNode?.value
  }
  
  override fun toString(): String {
    var outStr = ""
    var currNode = this.head
    while (currNode != null) {
      outStr += "[${currNode.value}] -> "
      currNode = currNode.next
    }
    return outStr.slice(0..outStr.length - 4)
  }
  
  fun walk(stopIndex: Int = INF.toInt()): Triple<Node<T>?, Node<T>?, Int> {
    var currNode = this.head
    var prevNode: Node<T>? = null
    var count = 0
    while (currNode != null && count < stopIndex) {
      prevNode = currNode
      currNode = currNode.next
      count++
    }
    return Triple(prevNode, currNode, count)
  }
  
  // TODO: make iterable
}

fun main() {
  var mylist = LinkedList<Int>()
  
  mylist.append(4)
  mylist.append(5)
  mylist.append(6)
  mylist.append(8)
  
  mylist.prepend(3)
  mylist.prepend(2)
  mylist.prepend(1)
  mylist.prepend(0)
  
  mylist.insert(7, 7)
  
  mylist.pop(mylist.length - 1)
  mylist.pop(0)
  mylist.pop(5)
  
  println(mylist.get(1))
  println(mylist)
  println(mylist.length)
  
  // for loop
}

// kotlinc linkedlist.kt -include-runtime -d linkedlist.jar && java -jar linkedlist.jar