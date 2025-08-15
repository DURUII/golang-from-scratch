package main

func getGroupEnd(head *ListNode, k int) *ListNode {
	cur := head
	for i := 0; i < k-1 && cur != nil; i++ {
		cur = cur.Next
	}
	return cur
}

func reverse(head *ListNode, nextGroupHead *ListNode) *ListNode {
	prev, cur := head, head
	for cur != nextGroupHead {
		next := cur.Next
		cur.Next = prev
		prev, cur = cur, next
	}
	return prev
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head}
	start, prevGroupEnd := head, dummy
	for start != nil {
		// 分组，找到每一组的开始和结束
		end := getGroupEnd(start, k)
		if end == nil {
			break
		}
		nextGroupHead := end.Next
		// 内部反转 + 翻转后的局部与整体的衔接
		prevGroupEnd.Next = reverse(start, nextGroupHead)
		start.Next = nextGroupHead
		// 分组迭代
		prevGroupEnd = start
		start = nextGroupHead
	}
	return dummy.Next
}
