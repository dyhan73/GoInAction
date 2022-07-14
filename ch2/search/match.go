package search

import "log"

// Result 검색 결과를 저장할 구조체
type Result struct {
	Field   string
	Content string
}

// Matcher 새로운 검색 타입을 구현할 때 필요한 동작을 정의한 인터페이스 정의
// Go 네이밍 규칙 : 인터페이스가 하나의 메서드만 선언하면 이름에 "~er" 로 끝나게 정의
type Matcher interface {
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}

// Match 함수는 고루틴으로 호출되며
// 개별 피드 타입에 대한 검색을 동시에 수행
func Match(matcher Matcher, feed *Feed, searchTerm string, results chan<- *Result) {
	// 지정된 검색기를 이용해 검색을 수행
	searchResults, err := matcher.Search(feed, searchTerm)
	if err != nil {
		log.Println(err)
		return
	}

	// 검색 결과를 채널에 기록
	for _, result := range searchResults {
		results <- result
	}
}

// Display 함수는 개별 고루틴이 전달한 검색 결과를 콘솔 창에 출력
func Display(results chan *Result) {
	// 채널은 검색 결과가 기록될 때까지 접근이 차단된다.
	// 채널이 닫히면 for 루프가 종료된다
	for result := range results {
		log.Printf("%s:\n%s\n\n", result.Field, result.Content)
	}
}
