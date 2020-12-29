package trie

import (
	"math/rand"
	"testing"
	"time"
)

var data [1000]string

const bytesPerString = 30
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func init() {
	rand.Seed(time.Now().UnixNano())

	for d := 0; d < len(data); d++ {
		b := make([]byte, bytesPerString)
		for i := range b {
			b[i] = letterBytes[rand.Intn(len(letterBytes))]
		}
		data[d] = string(b)
	}
}

// Store 30,000 random bytes repeatedly to see how many bytes are allocated by the trie.
func BenchmarkTrieMemory(b *testing.B) {
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		tr := New()

		for _, v := range data {
			tr.Insert(v)
		}
	}
}

func BenchmarkInsert(b *testing.B) {
	b.ReportAllocs()

	tr := New()

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		tr.Insert(data[n%len(data)])
	}

}

func BenchmarkExists(b *testing.B) {
	b.ReportAllocs()

	tr := New()

	for _, v := range data {
		tr.Insert(v)
	}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		tr.Exists(data[n%len(data)])
	}
}
