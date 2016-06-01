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

func Int32_to_bin(num int32) []byte {

	numBuf := new(bytes.Buffer)
	binary.Write(numBuf,binary.BigEndian, num)
	return numBuf.Bytes()

}

func Int8_to_bin(num int8) []byte {

	numBuf := new(bytes.Buffer)
	binary.Write(numBuf,binary.BigEndian, num)
	return numBuf.Bytes()

}


func Bin_to_int64(data []byte) int64 {

	var myint int64
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.BigEndian, &myint)
	return myint

}