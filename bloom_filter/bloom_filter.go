package algorithm

import (
	"bufio"
	"encoding/gob"
	"os"
	"time"

	"math/rand"

	"github.com/spaolacci/murmur3"
)

type BloomFilter struct {
	BitArray []byte
	BitSize  uint
	Seeds    []uint32
}

func NewBloomFilter(hashNum, arrayLen int) *BloomFilter {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	byteSize := (arrayLen + 7) / 8
	seeds := make([]uint32, hashNum)
	for i := 0; i < hashNum; i++ {
		seeds[i] = uint32(r.Int31())
	}

	return &BloomFilter{
		BitArray: make([]byte, byteSize),
		BitSize:  uint(arrayLen),
		Seeds:    seeds,
	}
}

func (bf *BloomFilter) getBit(index uint) bool {
	byteIndex := index / 8
	bitPosition := index % 8
	return bf.BitArray[byteIndex]&(1<<bitPosition) != 0
}

func (bf *BloomFilter) setBit(index uint) {
	byteIndex := index / 8
	bitPosition := index % 8
	bf.BitArray[byteIndex] |= 1 << bitPosition
}

func (bf *BloomFilter) Add(element string) {
	data := []byte(element)
	for _, seed := range bf.Seeds {
		hash := murmur3.New32WithSeed(seed)
		hash.Write(data)
		index := uint(hash.Sum32()) % bf.BitSize
		bf.setBit(index)
	}
}

func (bf *BloomFilter) Exists(element string) bool {
	data := []byte(element)
	for _, seed := range bf.Seeds {
		hash := murmur3.New32WithSeed(seed)
		hash.Write(data)
		index := uint(hash.Sum32()) % bf.BitSize
		if !bf.getBit(index) {
			return false
		}
	}
	return true
}

func (bf *BloomFilter) Dump(file string) error {
	f, err := os.OpenFile(file, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o666)
	if err != nil {
		return err
	}
	defer f.Close()

	writer := bufio.NewWriter(f)
	defer writer.Flush()

	encoder := gob.NewEncoder(writer)
	return encoder.Encode(*bf)
}

func LoadBloomFilter(file string) (*BloomFilter, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	decoder := gob.NewDecoder(reader)

	var bf BloomFilter
	if err := decoder.Decode(&bf); err != nil {
		return nil, err
	}
	return &bf, nil
}
