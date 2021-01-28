package addtwonumbers_2

type ListNode struct {
	Val int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	tempResult := result
	var up int
	for {
		if l1 == nil && l2 == nil {
			break
		}
		var v1, v2 int
		if l1 != nil {
			v1 = l1.Val
		}

		if l2 != nil {
			v2 = l2.Val
		}

		val := (v1 + v2 + up) % 10
		up = (v1 + v2 + up) / 10

		tempResult.Next = &ListNode{
			Val: val,
			Next: nil,
		}
		tempResult = tempResult.Next

		if l1 != nil {
			l1 = l1.Next
		}

		if l2 != nil {
			l2 = l2.Next
		}
	}

	if up > 0 {
		tempResult.Next = &ListNode{
			Val:up,
		}
	}

	return result.Next
}
