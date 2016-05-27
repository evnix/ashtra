package server
import ("io/ioutil"
		"os"
		"fmt"
		"github.com/evnix/ashtra/server/binhelp"
		"hash/crc32"
		)


type QFileOp struct {
   
   pushId int64
   pushCount int64
   pushFP* os.File

   popId int64
   popCount int64
   headOffset int64
}



type PushStruct struct{

	crc32 int64
	pushId int64
    pushCount int64


}

type PopStruct struct{

	crc32 int64
	popCount int64
    headOffset int64
	
}


func (m* QFileOp) OpenMetaFile(filepath string){

	

	data := make([]byte, 128)

	pushFP, err := os.OpenFile(filepath, os.O_RDWR,0777);

	if err != nil {
        panic(err)
    }

    pushFP.Read(data)


    objPush,res :=  BuildPushStruct(data)

    fmt.Println(res)

    objPush =objPush

    // pushHeadbin := data[16:64]
    // fmt.Println((pushHeadbin))
    // fmt.Println(len(pushHeadbin))

    // popHeadbin := data[64:128]
    // fmt.Println((popHeadbin))
    // fmt.Println(len(popHeadbin))


    m.pushFP = pushFP

}

func BuildPushStruct(data []byte) (PushStruct,bool) {


	evenPS:= PushStruct{}
	pushHeadbin := data[16:64]
	evenPS.crc32 = binhelp.Bin_to_int64(pushHeadbin[0:8])
	evenPS.pushId = binhelp.Bin_to_int64(pushHeadbin[8:16])
	evenPS.pushCount = binhelp.Bin_to_int64(pushHeadbin[16:24])

	fmt.Println( int64(crc32.ChecksumIEEE(pushHeadbin[8:24])))
	fmt.Println(evenPS.crc32)

	oddPS:= PushStruct{}
	oddPS.crc32 = binhelp.Bin_to_int64(pushHeadbin[24:32])
	oddPS.pushId = binhelp.Bin_to_int64(pushHeadbin[32:40])
	oddPS.pushCount = binhelp.Bin_to_int64(pushHeadbin[40:48])	

	if(ValidCRC32(evenPS.crc32,pushHeadbin[8:24]) && evenPS.pushId > oddPS.pushId){

		return evenPS,true

	} else if(ValidCRC32(oddPS.crc32,pushHeadbin[32:48]) && oddPS.pushId > evenPS.pushId){

		return oddPS,true

	} else if(ValidCRC32(evenPS.crc32,pushHeadbin[8:24]) && ValidCRC32(oddPS.crc32,pushHeadbin[32:48]) && evenPS.pushId == oddPS.pushId){

		return oddPS,true

	} else if (ValidCRC32(evenPS.crc32,pushHeadbin[8:24])){

		return evenPS,true

	} else if(ValidCRC32(oddPS.crc32,pushHeadbin[32:48])){

		return oddPS,true

	} else{

		return evenPS,false
	}



}


func ValidCRC32(numcrc32 int64, data []byte) bool{

	return ( numcrc32 ==  int64(crc32.ChecksumIEEE(data)) )

}

func CreateMetaFile(filepath string,  recordsPerShard int64){

	var version int64 = 1
		
	//Even Push ID 64 bits + Push Count 64 bits
	pushHeaderPart := append(binhelp.Int64_to_bin(0),binhelp.Int64_to_bin(0)...)
	pushHeaderCRC32 := int64(crc32.ChecksumIEEE(pushHeaderPart))

	//CRC32+00000000+00000000
	pushHeader_part := append(binhelp.Int64_to_bin(pushHeaderCRC32), pushHeaderPart...)

	//CRC32+00000000+00000000 + CRC32+00000000+00000000
	pushHeader := append(pushHeader_part,pushHeader_part...)

	//00000000+00000000+00000000
	popHeaderPart := append(binhelp.Int64_to_bin(0),pushHeaderPart...)
	popHeaderCRC32 := int64(crc32.ChecksumIEEE(popHeaderPart))

	//CRC32+00000000+00000000+00000000
	popHeader_part := append(binhelp.Int64_to_bin(popHeaderCRC32), popHeaderPart...)

	//CRC32+00000000+00000000+00000000  +  CRC32+00000000+00000000+00000000
	popHeader := append(popHeader_part,popHeader_part...)

	//Version+RecordsPerShard
	metaStart := append(binhelp.Int64_to_bin(version),binhelp.Int64_to_bin(recordsPerShard)...)

	//CRC32+00000000+00000000 + CRC32+00000000+00000000
	//CRC32+00000000+00000000+00000000  +  CRC32+00000000+00000000+00000000
	header_part := append(pushHeader,popHeader...)

	/* Here we have the complete Header file */
	//Version+RecordsPerShard
	//CRC32+00000000+00000000 + CRC32+00000000+00000000
	//CRC32+00000000+00000000+00000000  +  CRC32+00000000+00000000+00000000
	header := append(metaStart,header_part...)

	ioutil.WriteFile(filepath, header, os.ModeAppend | 0777)

}

