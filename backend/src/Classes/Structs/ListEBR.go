package structs

import (
	"fmt"
)

type Node struct {
	Ebr  *EBR
	Next *Node
	Prev *Node
}

func NewNode(ebr *EBR) *Node {
	return &Node{ebr, nil, nil}
}

type ListEBR struct {
	StartPartition int
	LastStart      int
	Size           int
	First          *Node
	Last           *Node
}

func NewListEBR(startPartition, size int) *ListEBR {
	return &ListEBR{
		startPartition,
		1024,
		size,
		nil,
		nil,
	}
}

func (l *ListEBR) Insert(newEBR *EBR) {
	if l.First != nil {
		if l.First.Ebr.Mount != '0' && newEBR.Start == l.StartPartition {
			newNode := NewNode(newEBR)
			if l.First.Next != nil {
				newEBR.Next = l.First.Next.Ebr.Start
				l.First.Next.Prev = newNode
				newNode.Next = l.First.Next
			}
			l.First = newNode
			l.LastStart = newEBR.Start
			return
		}
		if newEBR.Start > l.LastStart {
			l.Last.Next = NewNode(newEBR)
			l.Last.Ebr.Next = l.Last.Next.Ebr.Start
			l.Last.Next.Prev = l.Last
			l.Last = l.Last.Next
			l.LastStart = newEBR.Start
			return
		}
		current := l.First.Next
		newNode := NewNode(newEBR)
		for current != nil {
			if newEBR.Start > current.Ebr.Start && newEBR.Start < current.Next.Ebr.Start {
				current.Ebr.Next = newNode.Ebr.Start
				newNode.Ebr.Next = current.Next.Ebr.Start
				newNode.Prev = current
				newNode.Next = current.Next
				current.Next.Prev = newNode
				current.Next = newNode
				l.LastStart = newEBR.Start
				return
			}
			current = current.Next
		}
		return
	}
	l.First = NewNode(newEBR)
	l.Last = l.First
}

func (l *ListEBR) Delete(name string) *EBR {
	if l.First.Ebr.Name == name {
		deleted := l.First.Ebr
		newFirst := NewNode(&EBR{Next: l.First.Ebr.Next})
		newFirst.Next = l.First.Next
		l.First = newFirst
		return deleted
	}
	current := l.First
	for current.Next != nil {
		if current.Next.Ebr.Name == name {
			deleted := current.Next.Ebr
			current.Ebr.Next = current.Next.Ebr.Next
			if current.Next != nil {
				if current.Next.Next != nil {
					current.Next.Next.Prev = current
				}
			}
			return deleted
		}
		current = current.Next
	}
	return nil
}

func (l *ListEBR) SearchEmptySpace(newSize int) [][]int {
	emptySpaces := [][]int{}
	lastNoEmptyByte := l.StartPartition - 1
	if l.First.Ebr.Mount == '1' {
		lastNoEmptyByte += l.First.Ebr.Size
	}
	current := l.First.Next
	for current != nil {
		if current.Ebr.Start-lastNoEmptyByte > 2 && current.Ebr.Start-lastNoEmptyByte >= newSize+1 {
			emptySpaces = append(emptySpaces, []int{lastNoEmptyByte + 1, current.Ebr.Start - lastNoEmptyByte - 1})
		}
		lastNoEmptyByte = current.Ebr.Start + current.Ebr.Size - 1
		current = current.Next
	}
	if l.StartPartition+l.Size-lastNoEmptyByte > 2 && l.StartPartition+l.Size-lastNoEmptyByte >= newSize+1 {
		emptySpaces = append(emptySpaces, []int{lastNoEmptyByte + 1, l.StartPartition + l.Size - lastNoEmptyByte - 1})
	}
	return emptySpaces
}

func (l *ListEBR) Print() {
	current := l.First
	for current != nil {
		fmt.Print(current.Ebr)
		current = current.Next
	}
}

func (l *ListEBR) GetIterable() []*EBR {
	ebrs := []*EBR{}
	current := l.First
	for current != nil {
		ebrs = append(ebrs, current.Ebr)
		current = current.Next
	}
	return ebrs
}
