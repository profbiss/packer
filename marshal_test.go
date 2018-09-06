package packer_test

import (
	"reflect"
	"testing"
	"uint32-slice-packer-to-byte"
)

func BenchmarkMarshal(b *testing.B) {
	data := make([]packer.KeyValStore, 3000)
	var i uint32
	for i = 0; i < 3000; i++ {
		data[i] = packer.KeyValStore{255 + i, i}
	}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		packer.Marshal(data)
	}
}

var validResult = []byte{33, 3, 0, 0, 0, 2, 255, 0, 0, 0, 1, 1, 1, 1, 2}

func TestMarshal(t *testing.T) {
	data := make([]packer.KeyValStore, 3)
	var i uint32
	for i = 0; i < 3; i++ {
		data[i] = packer.KeyValStore{255 + i, i}
	}

	result := packer.Marshal(data)
	if !reflect.DeepEqual(validResult, result) {
		t.Error("Result incorrect")
	}
}

var numsToTest = map[uint32]uint32{
	255:      1,
	256:      2,
	65535:    2,
	65536:    3,
	16777215: 3,
	16777216: 4,
}

func TestNeedBytesUint32(t *testing.T) {
	for num, correctResult := range numsToTest {
		result := packer.NeedBytesUint32(num)
		if result != correctResult {
			t.Error("Result incorrect")
		}
	}
}
