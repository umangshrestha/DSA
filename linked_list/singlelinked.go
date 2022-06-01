package llist

import "fmt"

type LinkedList struct {
	next  *LinkedList
	value int
}

func (l *LinkedList) Insert(val int, pos int) {

	if pos == 0 {
		temp := *l
		l.next = &temp
		l.value = val

	} else {
		newNode := &LinkedList{value: val, next: nil}
		temp := l
		for i := 0; i < pos-1 && temp.next != nil; i++ {
			temp = temp.next
		}
		newNode.next = temp.next
		temp.next = newNode
	}
}

func (l *LinkedList) Append(val int) {
	newNode := &LinkedList{value: val, next: nil}
	temp := l
	for temp.next != nil {
		temp = l.next
	}
	temp.next = newNode
}

func (l *LinkedList) Print() {
	temp := l
	for temp != nil {
		fmt.Printf("%d ", temp.value)
		temp = temp.next
	}
	fmt.Println()
}

func (l *LinkedList) Sort() {
	lastNode := l
	for lastNode != nil {
		curNode := lastNode.next
		for curNode != nil {
			if lastNode.value > curNode.value {
				lastNode.value, curNode.value = curNode.value, lastNode.value
			}
			curNode = curNode.next
		}
		lastNode = lastNode.next
	}
}
