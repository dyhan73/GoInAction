package search

import (
	"encoding/json"
	"os"
)

const dataFile = "/Users/1001235/MyDev/ws_study/GoInAction/ch2/data/data.json" // 비공개 상수

// Feed 정보 구조체 (json 문서 내 필드이름과 매핑)
type Feed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

func RetrieveFeeds() ([]*Feed, error) {
	// 파일 열기
	file, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}

	// 함수 리턴 시 열어둔 파일 닫기. finally 대신 함수 종료 직후 처리할 작업 정의할 때 defer 사용 (panic 종료시에도 반드시 실행됨)
	defer file.Close()

	// 파일을 읽어 Feed 구조체 포인터의 슬라이스로 변환
	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)

	// err 는 여기서 처리하지 않고 호출한 함수가 처리하도록 반환
	return feeds, err
}
