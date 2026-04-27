package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var head *ListNode
	curr := &ListNode{}

	if list1 == nil && list2 == nil {
		return nil
	}

	for list1 != nil || list2 != nil {
		curr.Next = &ListNode{}
		curr = curr.Next

		if head == nil {
			head = curr
		}

		var candidate *ListNode
		if list1 != nil && list2 != nil {
			if list1.Val < list2.Val {
				candidate = list1
				list1 = list1.Next
			} else {
				candidate = list2
				list2 = list2.Next
			}
		} else if list1 != nil {
			candidate = list1
			list1 = list1.Next
		} else if list2 != nil {
			candidate = list2
			list2 = list2.Next
		}

		curr.Val = candidate.Val
	}

	return head
}
