package binhelp

import (
		"encoding/binary"
		"bytes"
		)

func Int64_to_bin(num int64) []byte {

	numBuf := new(bytes.Buffer)
	binary.Write(numBuf,binary.BigEndian, num)
	return numBuf.Bytes()

}