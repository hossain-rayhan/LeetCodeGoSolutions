//LeetCode- 002: Add Two Numbers
//Concept of LinkedList
//Time: O(N); Space O(1)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    head := &ListNode{}
    current := &ListNode{}
    
    carry := 0
    for ; l1 != nil || l2 != nil;{
        val := carry
        if l1 != nil {
            val += l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            val += l2.Val
            l2 = l2.Next
        }
        carry = val / 10
        val = val % 10
        n := &ListNode{val, nil}
        if  head.Next == nil {
            head.Next = n
            current = n
        }
        current.Next = n
        current = current.Next
    }
    return head.Next
}
