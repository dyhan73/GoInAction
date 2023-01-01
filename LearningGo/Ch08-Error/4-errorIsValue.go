package main

import (
	"errors"
	"fmt"
)

type Status int

const (
	InvalidLogin Status = iota + 1
	NotFound
)

type StatusErr struct {
	Status  Status
	Message string
}

func (se StatusErr) Error() string {
	return se.Message
}

func LoginAndGetData(uid, pwd, file string) ([]byte, error) {
	err := login(uid, pwd)
	if err != nil {
		return nil, StatusErr{
			Status:  InvalidLogin,
			Message: fmt.Sprintf("invalid credentials for user %s", uid),
		}
	}
	data, err := getData(file)
	if err != nil {
		return nil, StatusErr{
			Status:  NotFound,
			Message: fmt.Sprintf("file %s not found", file),
		}
	}
	return data, nil
}

func login(uid, pwd string) error {
	//return errors.New("mismatched")
	return nil
}

func getData(file string) ([]byte, error) {
	return nil, errors.New("file not found")
}

func errorTest1() {
	data, err := LoginAndGetData("asdf", "pwd", "path/file")
	if err != nil {
		fmt.Println("Error returned", err)
	}
	fmt.Println("result :", data)
}

func GenerateError(flag bool) error {
	var genErr StatusErr // 사용자 정의 에러는 기본값으로 초기화되면 nil이 아님
	fmt.Println(genErr.Status, genErr.Message)
	if flag {
		genErr = StatusErr{
			Status: NotFound,
		}
	}
	return genErr
}

func GenerateError2(flag bool) error {
	if flag {
		return StatusErr{
			Status: NotFound,
		}
	}
	return nil
}

func GenerateError3(flag bool) error {
	var genErr error // 변수를 error 타입으로 정의
	if flag {
		genErr = StatusErr{ // 변수에 사용자정의 에러 재할당
			Status: NotFound,
		}
	}
	return genErr
}

func main4() {
	errorTest1()
	err := GenerateError(true)
	fmt.Println(err != nil) // true
	err = GenerateError(false)
	fmt.Println(err != nil) // true

	err = GenerateError2(true)
	fmt.Println(err != nil) // true
	err = GenerateError2(false)
	fmt.Println(err != nil) // false

	err = GenerateError3(true)
	fmt.Println(err != nil) // true
	err = GenerateError3(false)
	fmt.Println(err != nil) // false
}
