package uint32_slice_packer_to_byte

func Marshal(data []KeyValStore) []byte {
	var keySize, valSize, mapLen, oneKeyValueSize, mapLenSize, formatSize, headerSize, maxVal, currentPosition, format uint32

	keySize, valSize, maxVal = calcSizeAndMaxVal(data)
	format = keySize<<4 | valSize
	formatSize = 1
	mapLen = uint32(len(data))
	mapLenSize = 4
	headerSize = formatSize + mapLenSize + valSize
	oneKeyValueSize = keySize + valSize
	resultBytes := make([]byte, headerSize+(mapLen*oneKeyValueSize))
	resultBytes[currentPosition] = byte(format)
	currentPosition++

	putUint(resultBytes, mapLen, currentPosition)
	currentPosition += mapLenSize

	putUint(resultBytes, maxVal, currentPosition)
	currentPosition += valSize

	for _, item := range data {
		putUint(resultBytes, item.Key, currentPosition)
		currentPosition += keySize

		putUint(resultBytes, item.Val, currentPosition)
		currentPosition += valSize
	}

	return resultBytes
}

func calcSizeAndMaxVal(data []KeyValStore) (keySize, valSize, maxCnt uint32) {
	for _, item := range data {
		tSize := GetRequiredNumberOfBytesUint32(item.Key)
		if keySize < tSize {
			keySize = tSize
		}

		tSize = GetRequiredNumberOfBytesUint32(item.Val)
		if valSize < tSize {
			valSize = tSize
		}

		if maxCnt < item.Val {
			maxCnt = item.Val
		}
	}
	return
}

func GetRequiredNumberOfBytesUint32(num uint32) uint32 {
	switch true {
	case num <= 1<<8-1: // 255
		return 1
	case num <= 1<<16-1: // 65535
		return 2
	case num <= 1<<24-1: // 16777215
		return 3
	default: // 4294967295
		return 4
	}
}

func putUint(buf []byte, num, start_pos uint32) {
	for num > 0 {
		buf[start_pos] = byte(num)
		num >>= 8
		start_pos++
	}
}