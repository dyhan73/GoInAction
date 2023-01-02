package main

import (
	"GoInAction/LearningGo/Ch09-Packages/3-exam/math"
	crand "crypto/rand"
	"encoding/binary"
	"fmt"
	"math/rand"
)

func packageExam() {
	num := math.Double(2)
	output := print.Format(num)
	fmt.Println(output)
}

func seedRand() *rand.Rand {
	var b [8]byte
	_, err := crand.Read(b[:]) // using renamed package
	if err != nil {
		panic("cannot seed with cryptographic random number generator")
	}
	r := rand.New(rand.NewSource(int64(binary.LittleEndian.Uint64(b[:]))))
	return r
}

// renameForUsingPackages An example for duplicated name between packages
// usually using rename for package is the solution.
func renameForUsingPackages() {
	r := seedRand()
	fmt.Println(r)
	fmt.Println(*r)
}

// shadowing
//
//	func rand() {
//		fmt.Println("redefined rand()")
//	}
func main() {
	packageExam()
	renameForUsingPackages()
	//rand()
}
