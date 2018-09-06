package packer_test

import (
	"reflect"
	"testing"
	"uint32-slice-packer-to-byte"
)

func BenchmarkUnmarshal(b *testing.B) {
	data := make([]packer.KeyValStore, 3000)
	var i uint32
	for i = 0; i < 3000; i++ {
		data[i] = packer.KeyValStore{255 + i, i}
	}
	testBytes := packer.Marshal(data)

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		packer.Unmarshal(testBytes)
	}
}

func TestUnmarshal(t *testing.T) {
	data := make([]packer.KeyValStore, 3)
	var i uint32
	for _, start_num := range []uint32{255, 65535, 16777215} {
		for i = 0; i < 3; i++ {
			data[i] = packer.KeyValStore{start_num + i, i}
		}

		testBytes := packer.Marshal(data)

		result, _ := packer.Unmarshal(testBytes)
		if !reflect.DeepEqual(data, result) {
			t.Error("Result incorrect")
		}
	}
}
