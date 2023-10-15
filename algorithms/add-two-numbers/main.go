package main

import (
	t "design-pattern/types"
	"fmt"
)

func main() {
	list1 := []int{1, 2, 3, 4, 9}
	list2 := []int{9, 9, 9, 9}

	fmt.Println(
		listToInt(
			addTwoNumbers(
				intToList(list1),
				intToList(list2))))
}

func addTwoNumbers(l1 *t.ListNode, l2 *t.ListNode) *t.ListNode {
	if l1 == nil {
		l1 = &t.ListNode{}
	}

	if l2 == nil {
		l2 = &t.ListNode{}
	}
	result := &t.ListNode{}
	result.Value = l1.Value + l2.Value

	if result.Value > 9 {
		result.Value -= 10
		if l2.Next == nil {
			l2.Next = &t.ListNode{Value: 0}
		}
		l2.Next.Value += 1
	}
	if l1.Next != nil || l2.Next != nil {
		result.Next = addTwoNumbers(l1.Next, l2.Next)
	}
	return result
}

func listToInt(head *t.ListNode) (res []int) {
	for head != nil {
		res = append(res, head.Value)
		head = head.Next
	}
	return
}

func intToList(nums []int) (res *t.ListNode) {
	if len(nums) == 0 {
		return nil
	}

	res = &t.ListNode{}
	tmp := res
	for _, v := range nums {
		tmp.Next = &t.ListNode{Value: v}
		tmp = tmp.Next
	}
	return res.Next
}
