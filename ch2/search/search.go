package search

import (
	"log"
	"sync"
)

// 패키지 수준의 변수 선언
// 소문자로 시작하는 변수는 외부에 노출하지 않는 변수임 (== 다른 패키지에서 직접 접근 불가)
var matchers = make(map[string]Matcher)

// Run 함수 - 검색 로직을 수행
func Run(searchTerm string) {
	// 검색할 피드의 목록 조회
	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatal(err)
	}

	// 버퍼가 없는 채널을 생성, 화면에 표시할 검색 결과를 전달받음
	results := make(chan *Result)

	// 모든 피드를 처리할 때 까지 기다릴 대기 그룹 설정 (WaitGroup 은 counting 세마포어)
	var waitGroup sync.WaitGroup

	// 개별 피드를 처리하는 동안 대히할 고루틴 개수 설정
	waitGroup.Add(len(feeds))

	// 각기 다른 종류의 피드를 처리할 고루틴 실행
	for _, feed := range feeds {
		// 검색을 위해 검색기를 조회
		matcher, exists := matchers[feed.Type] // map 에서 값과 존재여부 받고
		if !exists {                           // 존재여부 확인 코드 (예외처리 제공하지 않아요)
			matcher = matchers["default"]
		}

		// 검색을 실행하기 위해 고루틴 실행
		go func(matcher Matcher, feed *Feed) { // 익명함수를 고루틴으로 실행
			Match(matcher, feed, searchTerm, results) // searchTerm, results 는 익명함수에 파라미터가 아니어도 클로저를 통해 사용
			waitGroup.Done()                          // waiteGroup 카운트 감소, waitGroup 이 익명함수에 파라미터로 전달되지 않아도 클로저를 통해 사용 가능
		}(matcher, feed) // 익명함수에 파라미터 전달 (Go 의 모든 변수는 값에 의해 전달, 포인터 변수는 주소가 전달되어 참조 유지)
		// matcher 와 feed 도 클로저로 접근 가능하나, 상위 범위에서 루프를 돌면서 값이 변경되므로 파라미터로 현재 값의 복사본 전달이 필요함
	}

	// 모든 작업이 완료되었는지 모니터링할 익명함수를 고루틴으로 실행
	go func() { // 파라미터 없음 -> 모든 변수를 클로저를 통해 접근함
		// 모든 작업이 처리 (count 가 0) 될 때까지 기다림
		waitGroup.Wait()

		// Display 함수에게 프로그램 종료할 수 있음을 알리기 위한 채널 닫음
		close(results)
	}()

	// 검색 결과를 화면에 표시하고 마지막 결과를 표시한 후 리턴
	Display(results)
}

// Register - 프로그램에서 사용할 검색기 등록
func Register(feedType string, matcher Matcher) {
	if _, exists := matchers[feedType]; exists {
		log.Fatalln(feedType, "검색기가 이미 등록되었습니다.")
	}

	log.Println("등록 완료:", feedType, " 검색기")
	matchers[feedType] = matcher
}
