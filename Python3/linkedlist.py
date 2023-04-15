# Simple Linked List with Python

class Node:
    def __init__(self, value):
        self.value = value
        self.next_node = None


class LinkedList:
    def __init__(self):
        self.head = Node(None)
        self.tail = self.head  # references head node (same object)

    def append(self, node_val):
        next_node = Node(node_val)  # new node with given value as value
        temp = self.tail  # create temp reference to current tail
        temp.next_node = next_node
        self.tail = next_node

    def prepend(self, node_val):
        new_first = Node(node_val)
        temp = self.head.next_node
        self.head.next_node = new_first
        new_first.next_node = temp

    # inserted node will become nth element, nth element will become nth + 1
    def insert(self, node_val, index: int):
        new_node = Node(node_val)
        curr_node = self.head.next_node
        prev_node = None
        count = 0
        while count < index:
            prev_node = curr_node
            curr_node = curr_node.next_node
            count += 1
        prev_node.next_node = new_node
        new_node.next_node = curr_node

    def pop(self, index=-1):
        if index == 0:
            curr_node = self.head.next_node
            self.head.next_node = curr_node.next_node
        elif index > 0:
            count = 0
            curr_node = self.head.next_node
            prev_node = None
            next_node = None
            while count < index:
                prev_node = curr_node
                curr_node = curr_node.next_node
                next_node = curr_node.next_node
                count += 1
            prev_node.next_node = next_node
        # TODO: edit to allow for -2, -3, etc
        elif index == -1:
            curr_node = self.head
            prev_node = None
            while curr_node.next_node:
                prev_node = curr_node
                curr_node = curr_node.next_node
            prev_node.next_node = None
        else:
            pass

    def __iter__(self):
        self.curr_node = self.head
        return self

    def __next__(self):
        if self.curr_node.next_node:
            node_val = self.curr_node.next_node.value
            self.curr_node = self.curr_node.next_node
            return node_val
        raise StopIteration

    def __len__(self):
        return sum(1 for _ in self)

    def __str__(self):
        node = self.head
        out = ""
        while node.next_node:
            out += f"[{node.next_node.value}] -> "
            node = node.next_node
        return out[:-4]  # remove trailing arrow


def main():
    mylist = LinkedList()
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

    print(mylist)

    for node in mylist:
        print(node)


if __name__ == "__main__":
    main()
