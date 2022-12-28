package main

import (
	"fmt"
	"time"
)

type Score int

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func (p Person) String() string {
	return fmt.Sprintf("%s %s, age %d", p.FirstName, p.LastName, p.Age)
}
func basicUsage() {
	p := Person{
		FirstName: "Fred",
		LastName:  "Fredson",
		Age:       52,
	}
	output := p.String()
	fmt.Println(output)
}

// receiver type (value & pointer)
type Counter struct {
	total       int
	lastUpdated time.Time
}

func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

func (c Counter) String() string {
	return fmt.Sprintf("total: %d, last updated : %v", c.total, c.lastUpdated)
}

func usingPointReceiver() {
	var c Counter
	fmt.Println(c.String())
	c.Increment() // 값 타입인 c 를 포인터 리시버로 전달할 수 있음 (자동 변환해줌)
	fmt.Println(c.String())
	(&c).Increment()
	fmt.Println(c.String())
}

// 값 타입 파라미터로 받아 포인터리시버로 호출하면 로컬 복제본이 변경되는거임
func doUpdateWrong(c Counter) {
	c.Increment()
	fmt.Println("in doUpdateWrong:", c.String())
}

func doUpdateRight(c *Counter) {
	c.Increment()
	fmt.Println("in doUpdateRight:", c.String())
}

func funcLocalCopy() {
	var c Counter
	doUpdateWrong(c)
	fmt.Println("in main:", c.String())
	doUpdateRight(&c)
	fmt.Println(c)
	//doUpdateRight(c) // mismatch param type
	//fmt.Println(c)
}

// nil receiver
type IntTree struct {
	val         int
	left, right *IntTree
}

func (it *IntTree) Insert(val int) *IntTree {
	if it == nil {
		return &IntTree{val: val}
	}
	if val < it.val {
		it.left = it.left.Insert(val)
	} else if val > it.val {
		it.right = it.right.Insert(val)
	}
	return it
}

// Contains 는 리시버를 수정하지 않지만, nil 리시버 지원을 위해 포인터타입 사용
func (it *IntTree) Contains(val int) bool {
	switch {
	case it == nil:
		return false
	case val < it.val:
		return it.left.Contains(val)
	case val > it.val:
		return it.right.Contains(val)
	default:
		return true
	}
}
func nilReceiver() {
	var it *IntTree
	it = it.Insert(5)
	it = it.Insert(3)
	it = it.Insert(10)
	it = it.Insert(2)
	fmt.Println(it)
	fmt.Println(it.Contains(2))
	fmt.Println(it.Contains(12))
}

// method is a function
type Adder struct {
	start int
}

func (a Adder) AddTo(val int) int {
	return a.start + val
}

func callMethods() {
	myAdder := Adder{start: 10}
	fmt.Println(myAdder.AddTo(5))
	fmt.Println(myAdder)

	f1 := myAdder.AddTo
	fmt.Println(f1(10))
	f2 := Adder.AddTo
	fmt.Println(f2(myAdder, 15))
}

func userDefinedTypeExam() {
	type HighScore Score
	type Employee Person

	var i int = 300
	var s Score = 100
	var hs HighScore //= 200
	//hs = s  // type mismatch
	//s = i // type mismatch
	s = Score(i)
	hs = HighScore(s)

	fmt.Println(i, s, hs)
	fmt.Println(hs + 10)
	//fmt.Println(i+s) // type mismatch
}

func iotaEnum() {
	type MailCategory int
	const (
		aaa = 4
		bbb
		Uncategorized MailCategory = iota
		Personal
		Spam
		Social = 6
		Advertisements
		kkk
	)
	fmt.Println(aaa, bbb, Uncategorized, Personal, Spam, Social, Advertisements, kkk)
}

func main2() {
	//basicUsage()
	//usingPointReceiver()
	//funcLocalCopy()
	//nilReceiver()
	callMethods()
	userDefinedTypeExam()

	iotaEnum()
}
