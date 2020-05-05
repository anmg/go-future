package bloom_filter_test

import (
	"encoding/base64"
	"fmt"
	"github.com/kevburnsjr/bloomfilter"
	"testing"
)

//哈希函数个数k、位数组大小m、加入的字符串数量n
func TestBloomFilter(t *testing.T) {
	m, k := bloomfilter.EstimateParameters(10000, 0.001)
	bf := bloomfilter.New(m, k)
	bf.Add([]byte("foo"))
	bf.Add([]byte("bar"))

	fmt.Println(bf)

	encoded := base64.StdEncoding.EncodeToString(bf.ToBytes())
	fmt.Println(m)
	fmt.Println(k)
	fmt.Println(string(encoded))
}
