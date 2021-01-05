package trie

import (
	"encoding/csv"
	"io"
	"os"
	"strings"
	"testing"
)

var postcodes []string

func init() {
	onspd := "../ONSPD/Data/ONSPD_MAY_2020_UK.csv"

	file, err := os.Open(onspd)
	if err != nil {
		panic(err)
	}
	r := csv.NewReader(file)
	r.ReuseRecord = true

	// Skip header
	_, err = r.Read()
	if err != nil {
		panic(err)
	}

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		postcode := strings.ReplaceAll(record[0], " ", "")
		postcodes = append(postcodes, postcode)
	}

}

func BenchmarkImportONSPD(b *testing.B) {
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		tr := New()
		for _, v := range postcodes {
			tr.Insert(v)
		}
	}
}

func BenchmarkONSPDExists(b *testing.B) {
	b.ReportAllocs()

	tr := New()
	for _, v := range postcodes {
		tr.Insert(v)
	}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		if !tr.Exists(postcodes[n%len(postcodes)]) {
			b.Error("something went wrong")
		}
	}
}
