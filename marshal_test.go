package uint32_slice_packer_to_byte_test

import (
	"testing"
	"uint32-slice-packer-to-byte"
	"reflect"
)

func BenchmarkMarshal(b *testing.B) {
	data := make([]uint32_slice_packer_to_byte.KeyValStore, 3000)
	var i uint32;
	for i = 0; i < 3000; i++ {
		data[i] = uint32_slice_packer_to_byte.KeyValStore{255 + i, i}
	}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		uint32_slice_packer_to_byte.Marshal(data)
	}
}

var valid_result = []byte{33,3,0,0,0,2,255,0,0,0,1,1,1,1,2}

func TestMarshal(t *testing.T) {
	data := make([]uint32_slice_packer_to_byte.KeyValStore, 3)
	var i uint32;
	for i = 0; i < 3; i++ {
		data[i] = uint32_slice_packer_to_byte.KeyValStore{255+i, i}
	}

	result := uint32_slice_packer_to_byte.Marshal(data)
	if !reflect.DeepEqual(valid_result, result) {
		t.Error("Result incorrect")
	}
}

var nums_to_test = map[uint32]uint32{
	255: 1,
	256: 2,
	65535: 2,
	65536: 3,
	16777215: 3,
	16777216: 4,
}
func TestNeedBytesUint32(t *testing.T) {
	for num,correct_result := range nums_to_test {
		result := uint32_slice_packer_to_byte.NeedBytesUint32(num)
		if result != correct_result  {
			t.Error("Result incorrect")
		}
	}
}