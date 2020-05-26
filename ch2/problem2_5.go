// Problem 2.5
package main

import (
	"fmt"

	"github.com/chausen/go-tools/ds"
)

func main() {
	l1 := ds.CreateLinkedList(5)
	l1.Add(4)
	l1.Add(3)

	l2 := ds.CreateLinkedList(2)
	l2.Add(6)
	l2.Add(9)

	s := sum(l1, l2)

	l1.Print()
	fmt.Println("+")
	l2.Print()
	fmt.Println("=")
	s.Print()
}

func sum(l1, l2 *ds.LinkedList) *ds.LinkedList {
	if l1 == nil && l2 == nil {
		return nil
	} else if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	d1 := l1.Head
	d2 := l2.Head

	var sum *ds.LinkedList

	if d1 != nil && d2 != nil {

		sum = ds.CreateLinkedList(d1.Data + d2.Data)
		carry := 0

		for d1.Next != nil && d2.Next != nil {
			d1 = d1.Next
			d2 = d2.Next

			s := d1.Data + d2.Data + carry
			carry = s / 10
			s = s % 10
			sum.Add(s)
		}

		if carry > 0 {
			sum.Add(carry)
		}
	}

	return sum
}
