package qfileop

import (
	"errors"
	"fmt"
	"github.com/evnix/ashtra/server/binhelp"
	"github.com/satori/go.uuid"
	"hash/crc32"
	"io/ioutil"
	"os"
	"strconv"
)

var version int64 = 1
var minSupportedVersion int64 = 1

type QFileOp struct {
	filePath string

	pushId           int64
	pushFP           *os.File
	pushDataFP       *os.File
	prevPushShard    int64
	currentPushShard int64

	popId           int64
	headOffset      int64
	popFP           *os.File
	popDataFP       *os.File
	prevPopShard    int64
	currentPopShard int64

	recordsPerShard int64
}

type PushStruct struct {
	crc32  int64
	pushId int64
}

type PopStruct struct {
	crc32      int64
	popId      int64
	headOffset int64
}

type Record struct {
	uid        string
	errorCount int32
	expires    int64
	deleted    int8
	data       []byte
} 

func CreatePushHeader(pushId int64) []byte {

	pushIdBin := binhelp.Int64_to_bin(pushId)
	pushHeaderCRC32 := int64(crc32.ChecksumIEEE(pushIdBin))
	return append(binhelp.Int64_to_bin(pushHeaderCRC32), pushIdBin...)
}

func CreatePopHeader(popId int64, headOffset int64) []byte {

	popBin := append(binhelp.Int64_to_bin(popId), binhelp.Int64_to_bin(headOffset)...)
	popHeaderCRC32 := int64(crc32.ChecksumIEEE(popBin))
	return append(binhelp.Int64_to_bin(popHeaderCRC32), popBin...)
}

func CreateRecordBin(errorCount int32, expires int64, deleted int8, data []byte) ([]byte, string) {

	var dataLength int64 = int64(len(data))
	uid := uuid.NewV4().String()
	T0 := append(binhelp.Int32_to_bin(errorCount), binhelp.Int64_to_bin(expires)...)
	T1 := append(T0, binhelp.Int8_to_bin(deleted)...)
	T2 := append(T1, []byte(uid)...)
	T3 := append(T2, binhelp.Int64_to_bin(dataLength)...)
	T4 := append(T3, data...)
	return T4, uid

}

func (m *QFileOp) PopElement() (error, Record) {

	if m.popId == m.pushId {

		return nil, Record{}
	}

	m.currentPopShard = m.popId / m.recordsPerShard

	var err error
	if m.currentPopShard != m.prevPopShard || m.popDataFP == nil {

		if m.currentPopShard != m.prevPopShard {

			m.headOffset = 0
		}

		m.prevPopShard = m.currentPopShard
		filepath := m.filePath + "-" + strconv.FormatInt(m.currentPopShard, 10) + ".data"
		//fmt.Println(filepath)

		m.popDataFP, err = os.OpenFile(filepath, os.O_CREATE|os.O_RDWR, 0777)

		if err != nil {

			panic(err)

			return errors.New("onPop Error opening the data file: " + filepath), Record{}

		}

	}

	m.popDataFP.Seek(m.headOffset, 0)

	record := Record{}
	RecordMetaBin := make([]byte, 57)
	m.popDataFP.Read(RecordMetaBin)
	dataLength := binhelp.Bin_to_int64(RecordMetaBin[49:57])
	record.data = make([]byte, dataLength)
	m.popDataFP.Read(record.data)

	record.errorCount = (binhelp.Bin_to_int32(RecordMetaBin[0:4]))
	record.expires = (binhelp.Bin_to_int64(RecordMetaBin[4:12]))
	record.deleted = (binhelp.Bin_to_int8(RecordMetaBin[12:13]))
	record.uid = (string(RecordMetaBin[13:49]))

	m.popId++
	m.headOffset = m.headOffset + int64(len(RecordMetaBin)) + dataLength

	var popMetaHeaderOffset int64

	if m.popId%2 == 0 {

		popMetaHeaderOffset = 48

	} else {

		popMetaHeaderOffset = 72

	}

	m.popFP.Seek(popMetaHeaderOffset, 0)
	popHeader := CreatePopHeader(m.popId, m.headOffset)
	m.popFP.Write(popHeader)

	return nil, record
}

func (m *QFileOp) PushElement(errorCount int32, expires int64, deleted int8, data []byte) (error, string) {

	var err error
	m.currentPushShard = m.pushId / m.recordsPerShard

	if m.currentPushShard != m.prevPushShard || m.pushDataFP == nil {

		m.prevPushShard = m.currentPushShard
		filepath := m.filePath + "-" + strconv.FormatInt(m.currentPushShard, 10) + ".data"
		m.pushDataFP, err = os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)

		if err != nil {

			return errors.New("onPush Error opening the data file: " + filepath), ""

		}

	}

	recordBin, uid := CreateRecordBin(errorCount, expires, deleted, data)
	m.pushDataFP.Write(recordBin)

	m.pushId++

	var pushHeaderOffset int64

	if m.pushId%2 == 0 {

		pushHeaderOffset = 16

	} else {

		pushHeaderOffset = 32

	}

	m.pushFP.Seek(pushHeaderOffset, 0)
	pushHeader := CreatePushHeader(m.pushId)
	m.pushFP.Write(pushHeader)

	return nil, uid

}

func (m *QFileOp) OpenMetaFile(filepath string) error {

	m.filePath = filepath

	filepath = filepath + ".meta"

	//current size of header file
	data := make([]byte, 96)

	pushFP, err1 := os.OpenFile(filepath, os.O_RDWR, 0777)
	popFP, err2 := os.OpenFile(filepath, os.O_RDWR, 0777)

	if err1 != nil || err2 != nil {

		return errors.New("Error opening the metadata file: " + filepath)
	}

	pushFP.Read(data)

	objPush, res1 := BuildPushStruct(data)

	objPop, res2 := BuildPopStruct(data)

	currentMetaVersion := binhelp.Bin_to_int64(data[0:8])
	if currentMetaVersion < minSupportedVersion {

		return errors.New("Old MetaData file format not supported")
	}

	m.recordsPerShard = binhelp.Bin_to_int64(data[8:16])

	fmt.Println(m.recordsPerShard)

	if res1 == false || res2 == false {

		return errors.New("Corrupt MetaData file: " + filepath)

		// REturn error

	}

	m.pushId = objPush.pushId
	m.popId = objPop.popId
	m.headOffset = objPop.headOffset

	m.prevPopShard = m.popId / m.recordsPerShard

	//assign file pointers
	m.pushFP = pushFP
	m.popFP = popFP

	return nil

}

func BuildPopStruct(data []byte) (PopStruct, bool) {

	evenPS := PopStruct{}
	popHeadbin := data[48:96]
	evenPS.crc32 = binhelp.Bin_to_int64(popHeadbin[0:8])
	evenPS.popId = binhelp.Bin_to_int64(popHeadbin[8:16])
	evenPS.headOffset = binhelp.Bin_to_int64(popHeadbin[16:24])

	evenData := popHeadbin[8:24]

	// fmt.Println( int64(crc32.ChecksumIEEE(evenData)))
	// fmt.Println(evenPS.crc32)

	oddPS := PopStruct{}
	oddPS.crc32 = binhelp.Bin_to_int64(popHeadbin[24:32])
	oddPS.popId = binhelp.Bin_to_int64(popHeadbin[32:40])
	oddPS.headOffset = binhelp.Bin_to_int64(popHeadbin[40:48])

	oddData := popHeadbin[32:48]

	// fmt.Println( int64(crc32.ChecksumIEEE(oddData)))
	// fmt.Println(oddPS.crc32)

	if ValidCRC32(evenPS.crc32, evenData) && evenPS.popId > oddPS.popId {

		return evenPS, true

	} else if ValidCRC32(oddPS.crc32, oddData) && oddPS.popId > evenPS.popId {

		return oddPS, true

	} else if ValidCRC32(evenPS.crc32, evenData) && ValidCRC32(oddPS.crc32, oddData) && evenPS.popId == oddPS.popId {

		return oddPS, true

	} else if ValidCRC32(evenPS.crc32, evenData) {

		return evenPS, true

	} else if ValidCRC32(oddPS.crc32, oddData) {

		return oddPS, true

	} else {

		return evenPS, false
	}

}

func BuildPushStruct(data []byte) (PushStruct, bool) {

	evenPS := PushStruct{}
	pushHeadbin := data[16:48]
	evenPS.crc32 = binhelp.Bin_to_int64(pushHeadbin[0:8])
	evenPS.pushId = binhelp.Bin_to_int64(pushHeadbin[8:16])

	evenData := pushHeadbin[8:16]

	// fmt.Println( int64(crc32.ChecksumIEEE(pushHeadbin[8:24])))
	// fmt.Println(evenPS.crc32)

	oddPS := PushStruct{}
	oddPS.crc32 = binhelp.Bin_to_int64(pushHeadbin[16:24])
	oddPS.pushId = binhelp.Bin_to_int64(pushHeadbin[24:32])

	oddData := pushHeadbin[24:32]

	if ValidCRC32(evenPS.crc32, evenData) && evenPS.pushId > oddPS.pushId {

		return evenPS, true

	} else if ValidCRC32(oddPS.crc32, oddData) && oddPS.pushId > evenPS.pushId {

		return oddPS, true

	} else if ValidCRC32(evenPS.crc32, evenData) && ValidCRC32(oddPS.crc32, oddData) && evenPS.pushId == oddPS.pushId {

		return oddPS, true

	} else if ValidCRC32(evenPS.crc32, evenData) {

		return evenPS, true

	} else if ValidCRC32(oddPS.crc32, oddData) {

		return oddPS, true

	} else {

		return evenPS, false
	}

}

func ValidCRC32(numcrc32 int64, data []byte) bool {

	return (numcrc32 == int64(crc32.ChecksumIEEE(data)))

}

func (m *QFileOp) Close() {

	if m.pushFP != nil {

		m.pushFP.Close()
	}

	if m.pushDataFP != nil {

		m.pushDataFP.Close()
	}

	if m.popFP != nil {

		m.popFP.Close()
	}

	if m.popDataFP != nil {

		m.popDataFP.Close()
	}

}

func CreateMetaFile(filepath string, recordsPerShard int64) {

	//var version int64 = 1

	filepath = filepath + ".meta"

	//Even Push ID 64 bits
	pushHeaderPart := binhelp.Int64_to_bin(0)
	pushHeaderCRC32 := int64(crc32.ChecksumIEEE(pushHeaderPart))

	//CRC32+00000000
	pushHeader_part := append(binhelp.Int64_to_bin(pushHeaderCRC32), pushHeaderPart...)

	//CRC32+00000000 + CRC32+00000000
	pushHeader := append(pushHeader_part, pushHeader_part...)

	//00000000+00000000
	popHeaderPart := append(pushHeaderPart, pushHeaderPart...)
	popHeaderCRC32 := int64(crc32.ChecksumIEEE(popHeaderPart))

	//CRC32+00000000+00000000
	popHeader_part := append(binhelp.Int64_to_bin(popHeaderCRC32), popHeaderPart...)

	//CRC32+00000000+00000000 +  CRC32+00000000+00000000
	popHeader := append(popHeader_part, popHeader_part...)

	//Version+RecordsPerShard
	metaStart := append(binhelp.Int64_to_bin(version), binhelp.Int64_to_bin(recordsPerShard)...)

	//CRC32+00000000 + CRC32+00000000
	//CRC32+00000000+00000000  +  CRC32+00000000+00000000
	header_part := append(pushHeader, popHeader...)

	/* Here we have the complete Header file */
	//Version+RecordsPerShard
	//CRC32+00000000 + CRC32+00000000
	//CRC32+00000000+00000000  +  CRC32+00000000+00000000
	header := append(metaStart, header_part...)

	ioutil.WriteFile(filepath, header, 0777)

}
