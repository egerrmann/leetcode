package main

import "math"

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	var head, mergedList *ListNode

	for {
		minNode := popMinNode(lists)
		if minNode == nil {
			break
		}

		if mergedList == nil {
			mergedList = minNode
			head = mergedList
		} else {
			mergedList.Next = minNode
			mergedList = mergedList.Next
		}
	}

	return head
}

func popMinNode(lists []*ListNode) *ListNode {
	var minVal int = math.MaxInt
	var minValInd int

	for i, head := range lists {
		if head == nil {
			continue
		}

		if head.Val < minVal {
			minVal = head.Val
			minValInd = i
		}
	}

	if minVal == math.MaxInt {
		return nil
	}

	minNode := lists[minValInd]
	lists[minValInd] = minNode.Next
	minNode.Next = nil
	return minNode
}
