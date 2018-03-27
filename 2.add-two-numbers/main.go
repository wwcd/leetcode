package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var carry int
	var resl, tmpl *ListNode

	for l1 != nil || l2 != nil || carry != 0 {
		sum := 0

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		if resl == nil {
			tmpl = new(ListNode)
			resl = tmpl
		} else {
			tmpl.Next = new(ListNode)
			tmpl = tmpl.Next
		}

		tmpl.Val = (sum + carry) % 10
		carry = (sum + carry) / 10
	}

	return resl
}

func AddTwoNumbers(n1 []int, n2 []int) []int {
	var l1, l2, tmpl *ListNode

	for _, v := range n1 {
		if l1 == nil {
			tmpl = new(ListNode)
			l1 = tmpl
		} else {
			tmpl.Next = new(ListNode)
			tmpl = tmpl.Next
		}

		tmpl.Val = v
	}

	for _, v := range n2 {
		if l2 == nil {
			tmpl = new(ListNode)
			l2 = tmpl
		} else {
			tmpl.Next = new(ListNode)
			tmpl = tmpl.Next
		}

		tmpl.Val = v
	}

	var resn []int
	tmpl = addTwoNumbers(l1, l2)
	for tmpl != nil {
		resn = append(resn, tmpl.Val)
		tmpl = tmpl.Next
	}

	return resn
}

func main() {}
