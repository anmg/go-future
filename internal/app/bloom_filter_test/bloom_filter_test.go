package bloom_filter_test

import (
	"encoding/base64"
	"fmt"
	"github.com/httpimp/bloomfilter"
	"testing"
)

func TestBloomFilter(t *testing.T) {
	m, k := bloomfilter.EstimateParameters(10, 1e-6)
	bf := bloomfilter.New(m, k)
	bf.Add([]byte("foo"))
	bf.Add([]byte("bar"))
	encoded := base64.StdEncoding.EncodeToString(bf.ToBytes())
	fmt.Println(m)
	fmt.Println(k)
	fmt.Println(string(encoded))
}
