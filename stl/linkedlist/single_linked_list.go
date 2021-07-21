package linkedlist

import (
	"bytes"
	"errors"
	"fmt"
)

// 不带头节点的链表
type SingleLinkedList struct {
	Head *SingleLinkedListNode
}

type SingleLinkedListNode struct {
	Item interface{}
	Next *SingleLinkedListNode
}

func NewSingleLinkedList() *SingleLinkedList {
	return &SingleLinkedList{
		Head: nil,
	}
}

func(this *SingleLinkedList) AddintoHead(val interface{})  {
	// 1. save head
	oldhead := this.Head
	// 2. create new node
	node := &SingleLinkedListNode{
		Item: val,
		Next: oldhead,
	}
	// 3. modify new head node derect
	this.Head = node
}

func(this *SingleLinkedList) DeletefromHead() error {
	if this.Head == nil {
		// 没有元素,还要删除
		return errors.New("error, try pop an empty stack")
	}
	this.Head = this.Head.Next
	return nil
}

// 链表反转reverse the order of the linked list's items
// cur要转为指向之前的，变向后调整pre和cur， 适配边界情况      如果是递归实现的话就是子问题就是反转第一个后继续反转后面的，以要反转的节点做参数
func(this *SingleLinkedList) Reverse() {
	var pre, cur *SingleLinkedListNode = nil,this.Head
	for cur != nil {
		oldnext := cur.Next
		cur.Next = pre
		pre = cur
		cur = oldnext
	}
	this.Head = pre
}

// 链表转为字符串
func(this SingleLinkedList) ToString()  string{
	var ret bytes.Buffer
	ret.WriteString("SingleLinkedList Content: Head -> ")
	for it := this.Head; it != nil; it = it.Next {
		ret.WriteString(fmt.Sprintf("%v", it.Item))
		ret.WriteString(" -> ")
	}
	ret.WriteString(" nil")
	return ret.String()
}

// 遍历到尾部然后插入 //TODO
func(this *SingleLinkedList) AddintoTail(val interface{}) {

}


