package main

import "fmt"

func main2() {
	source := []string{"Apple", "orange", "plum", "banana", "grape"}

	slice := source[2:3:3]

	fmt.Printf("%s, %x\n", slice, slice)

	slice = append(slice, "kiwi")

	fmt.Printf("%s, %x\n", slice, slice)

	// iteration
	slice_itr := []int{10, 20, 30, 40}
	for index, value := range slice_itr {
		fmt.Printf("index : %d, value :%d, valueaddr : %X, itemaddr : %X\n", index, value, &value, &slice_itr[index])
	}
	fmt.Printf("len:%d, cap:%d\n", len(slice_itr), cap(slice_itr))

	// multidimensional slices
	slice_md := [][]int{{10}, {100, 200}}
	fmt.Println(slice_md)
	fmt.Printf("outer len : %d\n", len(slice_md))
	for i, v := range slice_md {
		fmt.Printf("inner[%d] %o, %d\n", i, v, len(v))
	}
	fmt.Printf("before append : %X\n", &slice_md[0][0])
	slice_md[0] = append(slice_md[0], 20)
	fmt.Println(slice_md)
	fmt.Printf("after append : %X\n", &slice_md[0][0])
}
