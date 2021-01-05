package trie

import (
	"math/rand"
	"testing"
	"time"
)

var data [2000000]string

const bytesPerString = 7

func init() {
	rand.Seed(time.Now().UnixNano())

	for d := 0; d < len(data); d++ {
		b := make([]byte, bytesPerString)
		for i := range b {
			b[i] = alphabet[rand.Intn(len(alphabet))]
		}
		data[d] = string(b)
	}
}

func BenchmarkTotalRandomMemory(b *testing.B) {
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		tr := New()

		for _, v := range data {
			tr.Insert(v)
		}
	}
}

func BenchmarkRandomInsert(b *testing.B) {
	b.ReportAllocs()

	tr := New()

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		tr.Insert(data[n%len(data)])
	}

}

func BenchmarkRandomExists(b *testing.B) {
	b.ReportAllocs()

	tr := New()

	for _, v := range data {
		tr.Insert(v)
	}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		if !tr.Exists(data[n%len(data)]) {
			b.Error("Something went wrong")
		}
	}
}
