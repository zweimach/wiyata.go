package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var result *ListNode
	carry := false

	for n1, n2 := l1, l2; n1 != nil || n2 != nil || carry; n1, n2 = next(n1), next(n2) {
		val := val(n1) + val(n2)

		if carry {
			val += 1
			carry = false
		}
		if val > 9 {
			val %= 10
			carry = true
		}
		result = appendListNode(result, val)
	}
	return result
}

func appendListNode(l *ListNode, val int) *ListNode {
	n := &ListNode{Val: val, Next: nil}
	if l == nil {
		return n
	}
	p := l
	for p.Next != nil {
		p = p.Next
	}
	p.Next = n
	return l
}

func val(l *ListNode) int {
	if l != nil {
		return l.Val
	}
	return 0
}

func next(l *ListNode) *ListNode {
	if l != nil {
		return l.Next
	}
	return nil
}

func newListNode(vals ...int) *ListNode {
	var result *ListNode
	for _, v := range vals {
		result = appendListNode(result, v)
	}
	return result
}

func compareListNode(l *ListNode, n *ListNode) bool {
	eq := l != nil && n != nil
	for p, q := l, n; p != nil && q != nil; p, q = p.Next, q.Next {
		eq = eq && p.Val == q.Val
	}
	return eq
}
