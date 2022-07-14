package matchers

import (
	"GoInAction/ch2/search"
	"encoding/xml"
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"
)

type (
	// item 구조체는 RSS문서 내의 item 태그에 정의된 필드들에 대응하는 필드를 선언
	item struct {
		XMLName     xml.Name `xml:"item"`
		PubDate     string   `xml:"pubDate"`
		Title       string   `xml:"title"`
		Description string   `xml:"description"`
		Link        string   `xml:"link"`
		GUID        string   `xml:"guid"`
		GeoRssPoint string   `xml:"georss:point"`
	}

	// image 구조체는 RSS문서 내의 image 태그에 정의된 필드에 대응하는 필드를 선언
	image struct {
		XMLName xml.Name `xml:"image"`
		URL     string   `xml:"url"`
		Title   string   `xml:"title"`
		Link    string   `xml:"link"`
	}

	// channel 구조체는 RSS 문서 내의 channel 태그에 정의된 필드에 대응하는 필드를 선언
	channel struct {
		XMLName        xml.Name `xml:"channel"`
		Title          string   `xml:"title"`
		Description    string   `xml:"description"`
		Link           string   `xml:"link"`
		PubDate        string   `xml:"pubDate"`
		LastBuildDate  string   `xml:"lastBuildDate"`
		TTL            string   `xml:"ttl"`
		Language       string   `xml:"language"`
		ManagingEditor string   `xml:"managingEditor"`
		WebMaster      string   `xml:"webMaster"`
		Image          image    `xml:"image"`
		Item           []item   `xml:"item"`
	}

	// rssDocument 구조체는 RSS 문서에 정의된 필드들에 대응하는 필드늘 정의
	rssDocument struct {
		XMLName xml.Name `xml:"rss"`
		Channel channel  `xml:"channel"`
	}
)

// rssMatcher 타입을 선언 - Matcher 인터페이스를 구현
type rssMatcher struct{}

// init 함수를 통해 프로그램에 검색기 등록
// main() 의 import 에 _ matchers 로 패키지가 임포트 되면서 init() 을 검색하여 main 실행 전 초기화 실행
func init() {
	var matcher rssMatcher
	search.Register("rss", matcher)
}

// Search 함수는 지정된 문서에서 검색어를 검색함
func (m rssMatcher) Search(feed *search.Feed, searchTerm string) ([]*search.Result, error) {
	var results []*search.Result

	log.Printf("피드 종류[%s] 사이트[%s] 주소[%s]에서 검색을 수행합니다.\n", feed.Type, feed.Name, feed.URI)

	// 검색할 데이터를 조회
	document, err := m.retrieve(feed)
	if err != nil {
		return nil, err
	}

	for _, channelItem := range document.Channel.Item {
		// 제목에서 검색어 조회
		matched, err := regexp.MatchString(searchTerm, channelItem.Title)
		if err != nil {
			return nil, err
		}

		// 검색어가 발견되면 결과에 저장
		if matched {
			results = append(results, &search.Result{
				Field:   "Title",
				Content: channelItem.Title,
			})
		}

		// 상세 내용에서 검색어 검색
		matched, err = regexp.MatchString(searchTerm, channelItem.Description)
		if err != nil {
			return nil, err
		}

		// 검색어가 발견되면 결과에 저장
		if matched {
			results = append(results, &search.Result{
				Field:   "Description",
				Content: channelItem.Description,
			})
		}
	}

	return results, nil
}

// HTTP Get 요청을 수행해서 RSS 피드를 요청한 후 결과를 디코딩
// retrieve 는 비공개 메서드
func (m rssMatcher) retrieve(feed *search.Feed) (*rssDocument, error) {
	if feed.URI == "" {
		return nil, errors.New("검색할 RSS 피드가 정의되지 않았습니다")
	}

	// 웹에서 RSS 문서를 조회
	resp, err := http.Get(feed.URI)
	if err != nil {
		return nil, err
	}

	// 함수가 리턴될 때 응답 스트림을 닫음
	defer resp.Body.Close()

	// 상태 코드가 200 인지 검사해서 올바른 응답 수신여부 확인
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP 응답 오류: %d\n", resp.StatusCode)
	}

	// RSS 피드 문서를 구조체 타입으로 디코드
	// 호출 함수가 에러를 판단하므로 여기에서 에러처리는 하지 않음
	var document rssDocument
	err = xml.NewDecoder(resp.Body).Decode(&document)
	return &document, err
}
