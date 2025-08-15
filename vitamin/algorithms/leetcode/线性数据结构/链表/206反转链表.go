package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head // terminal condition
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil // the final step
	return newHead
}

func reverseList2(head *ListNode) *ListNode {
	var prev *ListNode // a general case
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev, cur = cur, next
	}
	return prev
}
