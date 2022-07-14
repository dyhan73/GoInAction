package search

// 기본 검색기를 구현할 타입
// 빈 구조체 타입의 값이 생성될 때 0바이트 메모리 할당, 타입이 필요하지만 상태관리가 필요없는 경우에 사용
type defaultMatcher struct{}

// 기본 검색기를 프로그램에 등록
// search 패키지가 import 될 때 init() 함수가 발견되고 컴파일러는 init 이 main 함수 호출 전에 호출하도록 예약함
func init() {
	var matcher defaultMatcher
	Register("default", matcher)
}

// Search 기본검색기 함수 구현 - defaultMatcher type은 Matcher 인터페이스 요구사항인 Search 를 구현함
// defaultMatcher 타입에 대한 value receiver 로 선언 -> 해당 메서드는 지정된 수신기 타입에만 연결
// 일반적으로 포인터 수신기를 사용하는 것이 권장되나 여기서 defaultMatcher 는 0바이트이므로 포인터보다 메모리 점유율이 낮아 포인터 사용하지 않음
// 포인터 수신기에 선언 : 인터페이스 타입에 대한 포인터를 통해서만 호출가능
// 값 수신기에 선언 : 값과 포인터 변수 모두를 통해 호출이 가능
func (m defaultMatcher) Search(feed *Feed, searchTerm string) ([]*Result, error) {
	return nil, nil
}
