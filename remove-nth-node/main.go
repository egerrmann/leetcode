package main

import "fmt"

func main() {
	l := buildList(1, 2, 3, 4, 5)
	printList(l)

	fmt.Println()

	l = removeNthFromEnd(l, 1)
	printList(l)
}

func buildList(nums ...int) *ListNode {
	var res *ListNode
	for i := len(nums) - 1; i >= 0; i-- {
		res = &ListNode{Next: res, Val: nums[i]}
	}
	return res
}

func printList(head *ListNode) {
	if head == nil {
		return
	}
	fmt.Printf("%d ", head.Val)
	printList(head.Next)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	listLen := calculateLen(head)

	if listLen == n {
		return head.Next
	}

	ctr := listLen
	headCpy := head
	for ctr > n+1 {
		headCpy = headCpy.Next
		ctr--
	}

	curr, next := headCpy, headCpy.Next
	curr.Next = next.Next

	return head
}

func calculateLen(head *ListNode) int {
	listLen := 1
	curr := head
	for curr.Next != nil {
		listLen++
		curr = curr.Next
	}
	return listLen
}
