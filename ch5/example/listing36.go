package example

import "fmt"

// 알림을 수행하는 동작을 정의하는 notifier 인터페이스 선언
type notifier interface {
	notify()
}

// 사용자를 표현하는 uer 타입 선언
type user struct {
	name  string
	email string
}

// 포인터 수신자를 이용하여 notify 메서드 구현
func (u *user) notify() {
	fmt.Printf("사용자에게 메일을 전송합니다.: %s<%s>\n", u.name, u.email)
}

func Main36() {
	u := user{"Bill", "bill@email.com"}

	sendNotification(&u)
}

func sendNotification(n notifier) {
	n.notify()
}
