/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */


func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    cur1, cur2 := list1, list2
	var cur, prev, head *ListNode

	if list1 == nil && list2 == nil {
		return nil
	}

	if list1 == nil && list2 != nil {
		return list2
	}

	if list1 != nil && list2 == nil {
		return list1
	}

	if list1.Val < list2.Val {
		cur, prev, head = list1, list1, list1
	} else {
		cur, prev, head = list2, list2, list2
	}


	for cur != nil {
		prev = cur
		
		if cur1 == nil && cur2 == nil {
			break
		}

		if cur1 != nil && cur2 == nil {
			cur = cur1
			cur1 = cur1.Next
			prev.Next = cur 
			continue
		}

		if cur1 == nil && cur2 != nil {
			cur = cur2
			cur2 = cur2.Next
			prev.Next = cur
			continue
		}
		
		if cur1.Val < cur2.Val {
			cur = cur1
			cur1 = cur1.Next
		} else {
			cur = cur2
			cur2 = cur2.Next
		}
		
		prev.Next = cur
	}

	return head
}

// func print(node *ListNode) {
// 	if node == nil {
// 		fmt.Println("nil")
// 		return
// 	}

// 	n := -1
// 	if node.Next != nil {
// 		n = node.Next.Val
// 	}

// 	fmt.Printf("%d, %d \n", node.Val, n)
// }
