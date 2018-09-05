package uint32_slice_packer_to_byte

type KeyValStore struct {
	Key, Val uint32
}

func Unmarshal(packedData []byte) ([]KeyValStore, uint32) {
	var format, keySize, valSize, mapLen, i, currentPosition, mapLenSize, maxVal uint32

	mapLenSize = 4

	format = uint32(packedData[currentPosition])
	currentPosition++

	keySize = format >> 4
	valSize = format & 15

	mapLen = bytesToUint32(packedData[currentPosition: currentPosition+mapLenSize])
	currentPosition += mapLenSize

	maxVal = bytesToUint32(packedData[currentPosition: currentPosition+valSize])
	currentPosition += valSize

	resultMap := make([]KeyValStore, mapLen)

	for i = 0; i < mapLen; i++ {
		key := bytesToUint32(packedData[currentPosition: currentPosition+keySize])
		currentPosition += keySize
		val := bytesToUint32(packedData[currentPosition: currentPosition+valSize])
		resultMap[i] = KeyValStore{key, val}

		currentPosition += valSize
	}

	return resultMap, maxVal
}

func bytesToUint32(bytes []byte) (result uint32) {
	for i, byte := range bytes {
		if i == 0 {
			result = uint32(byte)
		} else {
			result |= uint32(byte) << (8 * uint32(i))
		}
	}
	return
}
