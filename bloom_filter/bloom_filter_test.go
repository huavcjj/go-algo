package algorithm

import (
	"testing"
)

func TestBloomFilter(t *testing.T) {
	bf := NewBloomFilter(8, 1<<20)
	bf.Add("hello")
	bf.Add("golang")

	if !bf.Exists("hello") {
		t.Fail()
	}
	if !bf.Exists("golang") {
		t.Fail()
	}
	if bf.Exists("world") {
		t.Fail()
	}
}

func TestDumpAndLoadBloomFilter(t *testing.T) {
	bf := NewBloomFilter(8, 1<<20)
	bf.Add("hello")
	bf.Add("golang")

	tmpFile := "bloom_filter.gob"
	if err := bf.Dump(tmpFile); err != nil {
		t.Error(err)
	}

	loadedBf, err := LoadBloomFilter(tmpFile)
	if err != nil {
		t.Error(err)
	}

	if !loadedBf.Exists("hello") {
		t.Fail()
	}
	if !loadedBf.Exists("golang") {
		t.Fail()
	}
	if loadedBf.Exists("world") {
		t.Fail()
	}
}

func BenchmarkBloomFilter(b *testing.B) {
	bf := NewBloomFilter(8, 1<<20)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bf.Add("hello")
		bf.Exists("hello")
	}
}

func BenchmarkBloomFilterDump(b *testing.B) {
	bf := NewBloomFilter(8, 1<<20)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bf.Dump("bloom_filter.gob")
	}
}

func BenchmarkBloomFilterLoad(b *testing.B) {
	bf := NewBloomFilter(8, 1<<20)
	bf.Dump("bloom_filter.gob")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LoadBloomFilter("bloom_filter.gob")
	}
}
