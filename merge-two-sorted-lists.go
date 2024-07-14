package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := new(ListNode)
	sortedList := dummy

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			sortedList.Next = list1
			list1 = list1.Next

		} else {
			sortedList.Next = list2
			list2 = list2.Next
		}

		sortedList = sortedList.Next
	}

	if list1 != nil {
		sortedList.Next = list1

	} else if list2 != nil {
		sortedList.Next = list2
	}

	return dummy.Next
}

func main() {
	list1 := &ListNode{1, &ListNode{2, &ListNode{4, &ListNode{5, nil}}}}
	list2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}

	merge := mergeTwoLists(list1, list2)
	for merge != nil {
		fmt.Println(merge.Val)
		merge = merge.Next
	}
}
