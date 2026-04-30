package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	tempVal := new(int)
	var curr, prev *ListNode = head, nil

	for i := 0; ; i++ {
		if curr == nil {
			return head
		}

		if i%2 == 1 {
			*tempVal = prev.Val
			prev.Val = curr.Val
			curr.Val = *tempVal
		}

		prev = curr
		curr = curr.Next
	}
}
