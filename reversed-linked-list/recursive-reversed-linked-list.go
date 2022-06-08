package reversed_linked_list

type ListNode2 struct {
	Val  int
	Next *ListNode2
}

func recursiveReverseList(head *ListNode2) *ListNode2 {
	if head == nil {
		return nil
	}
	var newHead = head
	if head.Next != nil {
		newHead = recursiveReverseList(head.Next)
		head.Next.Next = head
		head.Next = nil

	}
	return newHead

}
