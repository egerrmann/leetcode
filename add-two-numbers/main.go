package main

import "fmt"

func main() {
	a := initListNode([]int{9, 9, 9, 9, 9, 9, 9})
	b := initListNode([]int{9, 9, 9, 9})
	res := addTwoNumbers(a, b)
	printListNode(res)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	resList := new(ListNode)
	currList := resList
	var hasTransferredDigit bool
	for {
		var a, b int
		if l1 != nil {
			a = l1.Val
		}
		if l2 != nil {
			b = l2.Val
		}

		sum := a + b
		if hasTransferredDigit {
			sum += 1
		}

		hasTransferredDigit = false
		if sum > 9 {
			hasTransferredDigit = true
		}

		resDigit := sum % 10
		currList.Val = resDigit
		if l1.Next == nil && l2.Next == nil {
			if hasTransferredDigit {
				l1 = new(ListNode)
				l2 = new(ListNode)
			} else {
				break
			}
		} else {
			if l1.Next == nil {
				l1.Val = 0
			} else {
				l1 = l1.Next
			}

			if l2.Next == nil {
				l2.Val = 0
			} else {
				l2 = l2.Next
			}
		}

		currList.Next = new(ListNode)
		currList = currList.Next
	}

	return resList
}

func initListNode(arr []int) *ListNode {
	res := new(ListNode)
	currList := res
	for i, v := range arr {
		currList.Val = v
		if i == len(arr)-1 {
			break
		}
		currList.Next = new(ListNode)
		currList = currList.Next
	}

	return res
}

func printListNode(list *ListNode) {
	for list != nil {
		fmt.Printf("%v ", list.Val)
		list = list.Next
	}
}
