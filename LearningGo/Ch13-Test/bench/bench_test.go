package bench

import "testing"

func TestFileLen(t *testing.T) {
	result, err := FileLen("bench.go", 1)
	if err != nil {
		t.Fatal(err)
	}
	if result != 309 {
		t.Error("Expected 309, got", result)
	}
}

var blackhole int

func BenchmarkFileLen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result, err := FileLen("bench.go", 1)
		if err != nil {
			b.Fatal(err)
		}
		blackhole = result
	}
}
