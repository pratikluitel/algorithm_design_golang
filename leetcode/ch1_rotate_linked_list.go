/*
	Code Author	: github.com/pratikluitel
	Book		: The Algorithm Design Manual, written by Steven S Skiena
	Chapter 	: 1

	Q			: Given the head of a linked list, rotate the list to the right by k places.

	Link		: https://leetcode.com/problems/rotate-list/
*/

package leetcode

import (
	"fmt"
	"time"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Display the whole linked list when the head of the list is supplied.
func display_list(head *ListNode) {
	fmt.Printf("%v->", head.Val)
	for {
		if head.Next == nil {
			break
		}
		fmt.Printf("%v->", head.Next.Val)
		head = head.Next

	}
	fmt.Printf("\n\n")
	return
}

// converts non circular linked list to circular
func countAndConnectList(head *ListNode) (*ListNode, int) {
	if head == nil {
		return head, 0
	}
	first_node := head
	count := 1
	for head.Next != nil {
		head = head.Next
		count++
	}
	head.Next = first_node

	return head.Next, count //return the first node
}

// Rotates a linked list right by k times
func rotateRight(head *ListNode, k int) *ListNode {

	var effective_rotations int
	var new_head *ListNode
	head, count := countAndConnectList(head)

	if count == 0 { // the same list if zero elements in list
		return head
	}

	// since rotation is cyclic e.g. rotating 6 number of times is the same as rotating 1 time for a 5 length linked list
	effective_rotations = k % count

	// changes the head and tail of the list to fit rotation requirements
	for i := 0; i < count-effective_rotations-1; i++ {
		head = head.Next
	}
	new_head = head.Next
	head.Next = nil

	return new_head
}

func Run_2() {
	// creating a 1->2->3->4->5->...->59-> linked list
	list := []*ListNode{}

	for i := 1; i < 60; i++ {
		if i == 1 {
			node := ListNode{60 - i, nil}
			list = append(list, &node)
		} else {
			node := ListNode{60 - i, list[len(list)-1]}
			list = append(list, &node)
		}

	}
	head := list[len(list)-1]
	display_list(head)

	k := 40
	start := time.Now()
	new_head := rotateRight(head, k)
	fmt.Printf("Rotation complete, rotated %d times, execution time:%s\n\nRotated result:\n", k, time.Since(start))

	display_list(new_head)
}
