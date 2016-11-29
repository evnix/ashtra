package binhelp

import (
	"bytes"
	"encoding/binary"
)

func Int64_to_bin(num int64) []byte {

	numBuf := new(bytes.Buffer)
	binary.Write(numBuf, binary.BigEndian, num)
	return numBuf.Bytes()

} 

func Int32_to_bin(num int32) []byte {

	numBuf := new(bytes.Buffer)
	binary.Write(numBuf, binary.BigEndian, num)
	return numBuf.Bytes()

}

func Int8_to_bin(num int8) []byte {

	numBuf := new(bytes.Buffer)
	binary.Write(numBuf, binary.BigEndian, num)
	return numBuf.Bytes()

}

func Bin_to_int64(data []byte) int64 {

	var myint int64
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.BigEndian, &myint)
	return myint

}

func Bin_to_int32(data []byte) int32 {

	var myint int32
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.BigEndian, &myint)
	return myint

}

func Bin_to_int8(data []byte) int8 {

	var myint int8
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.BigEndian, &myint)
	return myint

}
