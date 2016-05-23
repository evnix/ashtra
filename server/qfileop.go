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


    objPush :=  BuildPushStruct(data)

    objPush =objPush

    // pushHeadbin := data[16:64]
    // fmt.Println((pushHeadbin))
    // fmt.Println(len(pushHeadbin))

    // popHeadbin := data[64:128]
    // fmt.Println((popHeadbin))
    // fmt.Println(len(popHeadbin))


    m.pushFP = pushFP

}

func BuildPushStruct(data []byte) PushStruct{


	ps1:= PushStruct{}
	pushHeadbin := data[16:64]
	ps1.crc32 = binhelp.Bin_to_int64(pushHeadbin[0:8])
	ps1.pushId = binhelp.Bin_to_int64(pushHeadbin[8:16])
	ps1.pushCount = binhelp.Bin_to_int64(pushHeadbin[16:24])

	fmt.Println( int64(crc32.ChecksumIEEE(pushHeadbin[8:24])))
	fmt.Println(ps1.crc32)

	ps2:= PushStruct{}
	ps2.crc32 = binhelp.Bin_to_int64(pushHeadbin[24:32])
	ps2.pushId = binhelp.Bin_to_int64(pushHeadbin[32:40])
	ps2.pushCount = binhelp.Bin_to_int64(pushHeadbin[40:48])	


    fmt.Println((ps2.crc32))
    //fmt.Println(len(pushHeadbin))

    return ps1;



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

