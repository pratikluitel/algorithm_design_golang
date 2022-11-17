/*
	Code Author	: github.com/pratikluitel
	Book		: The Algorithm Design Manual, written by Steven S Skiena
	Chapter 	: 1

	Q			: Given the head of a linked list, rotate the list to the right by k places.

	Link		: https://leetcode.com/problems/rotate-list/
*/

package leetcode

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func display_list(head *ListNode) { // display the whole linked list when the head of the list is supplied
	fmt.Printf("Linked list:\n%v->", head.Val)
	for {
		if head.Next == nil {
			break
		}
		fmt.Printf("%v->", head.Next.Val)
		head = head.Next

	}
	fmt.Printf("\n")
	return
}

func countList(head *ListNode) int { //counts the number of elements in the linked list
	count := 0
	if head == nil {
		return count
	}
	for head != nil {
		head = head.Next
		count++
	}

	return count
}

func connectList(head *ListNode, count int) *ListNode { // converts non circular linked list to circular
	var first_node *ListNode
	temp_head := head

	for i := 0; i < count; i++ {
		if i == 0 {
			first_node = temp_head
		}
		if head.Next == nil {
			temp_head.Next = first_node
		} else {
			head = head.Next
		}
		temp_head = temp_head.Next
	}
	return temp_head
}

func rotateRight(head *ListNode, k int) *ListNode {

	count := countList(head)

	if count == 0 || count == 1 { // the same list if no rotations or only one/zero elements in list
		return head
	}

	/// since rotation is cyclic i.e. rotating 6 number of times is the same as rotating 1 time for a 5 length linked list
	var effective_rotations int
	effective_rotations = k % count

	if effective_rotations == 0 {
		return head
	}

	temp_head := connectList(head, count)

	// changes the head and tail of the list to fit rotation requirements
	var new_head *ListNode
	for i := 0; i < count; i++ {
		if i == count-effective_rotations-1 {
			new_head = temp_head.Next
			temp_head.Next = nil
			break
		}
		temp_head = temp_head.Next
	}

	return new_head
}

func Run_2() {
	// a 1->2->3->4->5 linked list
	list := []ListNode{}

	for i := 1; i < 6; i++ {
		if i == 1 {
			node := ListNode{6 - i, nil}
			list = append(list, node)
		} else {
			node := ListNode{6 - i, &list[len(list)-1]}
			list = append(list, node)
		}

	}
	head := list[len(list)-1]
	display_list(&head)
	new_head := rotateRight(&head, 9)
	display_list(new_head)
}
