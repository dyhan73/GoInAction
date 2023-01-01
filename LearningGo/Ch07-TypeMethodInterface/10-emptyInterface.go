package main

import "fmt"

type LinkedList struct {
	Value interface{}
	Next  *LinkedList
}

func (ll *LinkedList) Insert(pos int, val interface{}) *LinkedList {
	if ll == nil || pos == 0 {
		return &LinkedList{
			Value: val,
			Next:  ll,
		}
	}
	ll.Next = ll.Next.Insert(pos-1, val)
	return ll
}

func main10() {
	l := &LinkedList{
		Value: "KK",
		Next:  nil,
	}
	fmt.Println(l)
	l = l.Insert(2, "ka1")
	fmt.Println(l, l.Next)
	l = l.Insert(2, "ka2")
	fmt.Println(l, l.Next, l.Next.Next)
	l = l.Insert(1, "ka3")
	fmt.Println(l, l.Next, l.Next.Next, l.Next.Next.Next)
}
