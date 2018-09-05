package uint32_slice_packer_to_byte_test

import (
	"reflect"
	"testing"
	"uint32-slice-packer-to-byte"
)

func BenchmarkUnmarshal(b *testing.B) {
	data := make([]uint32_slice_packer_to_byte.KeyValStore, 3000)
	var i uint32
	for i = 0; i < 3000; i++ {
		data[i] = uint32_slice_packer_to_byte.KeyValStore{255 + i, i}
	}
	testBytes := uint32_slice_packer_to_byte.Marshal(data)

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		uint32_slice_packer_to_byte.Unmarshal(testBytes)
	}
}

func TestUnmarshal(t *testing.T) {
	data := make([]uint32_slice_packer_to_byte.KeyValStore, 3)
	var i uint32
	for _, start_num := range []uint32{255, 65535, 16777215} {
		for i = 0; i < 3; i++ {
			data[i] = uint32_slice_packer_to_byte.KeyValStore{start_num + i, i}
		}

		testBytes := uint32_slice_packer_to_byte.Marshal(data)

		result, _ := uint32_slice_packer_to_byte.Unmarshal(testBytes)
		if !reflect.DeepEqual(data, result) {
			t.Error("Result incorrect")
		}
	}
}
