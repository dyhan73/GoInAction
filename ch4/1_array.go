package main

import "fmt"

func main1() {
	// 5개 원소로 구성된 정수 배열 선언, 제로값 초기화
	var array [5]int
	fmt.Println("len :", len(array), ", idx(1) : ", array[1]) // 5, 0

	// 배열 리터럴을 사용한 초기화
	array2 := [5]int{10, 20, 30, 40, 50}
	fmt.Println("len :", len(array2), ", idx(1) :", array2[1])

	// 초기값에 따라서 길이는 알아서 해줘
	array3 := [...]int{10, 20, 30}
	println("len :", len(array3), ", idx(1) :", array3[1])

	// 길이는 결정하는데 일부만 초기화 할께 (나머지는 제로값 초기화)
	array4 := [5]int{1: 10, 2: 20}
	println("len :", len(array4), ", idx(1) :", array4[1], ", idx(3) :", array4[3])

	// 배열에 대한 포인터로 사용
	arrayp := [5]*int{0: new(int), 1: new(int)}
	*arrayp[0] = 10
	*arrayp[1] = 20
	println(&arrayp, arrayp[0], *arrayp[0], arrayp[1], *arrayp[1], arrayp[4])
	println(&arrayp[0], &arrayp[1])

	// 배열 값 복사하기
	var arraycp1 [5]string
	arraycp2 := [5]string{"red", "blue", "green", "yello", "pink"}
	arraycp1 = arraycp2
	println(&arraycp1[1], arraycp1[1], &arraycp2[1], arraycp2[1])
	// 배열 크기가 다르면 복사되나?
	//var arraycp3 [6]string
	//arraycp3 = arraycp1 // 컴파일 에러나요.
	//println(arraycp3[1], arraycp1[1])
	// 포인터 복사하면?
	var arrayp2 [5]*int
	arrayp2 = arrayp
	println(arrayp[0], arrayp2[0]) // 참조 주소가 복사됨 - 참조 공유함

	// 다차원 배열 선언 (제로 초기화)
	var arraym [4][2]int
	println(arraym[3][1])

	// 배열 리터럴을 이용한 다차원 배열 초기화
	arraym2 := [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}
	println(arraym2[2][1])

}
