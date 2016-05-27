package server
import ("io/ioutil"
		"os"
		"fmt"
		"github.com/evnix/ashtra/server/binhelp"
		"hash/crc32"
		"errors"
		)


var version int64 = 1
var minSupportedVersion int64 = 1

type QFileOp struct {
   
   pushId int64
   pushCount int64
   pushFP* os.File
   pushDataFP* os.File
   prevShard int64
   currentShard int64


   popId int64
   popCount int64
   headOffset int64
   popFP* os.File


   recordsPerShard int64

}



type PushStruct struct{

	crc32 int64
	pushId int64
    pushCount int64


}

type PopStruct struct{

	crc32 int64
	popId int64
	popCount int64
    headOffset int64
	
}


func (m* QFileOp) PushElement(data []byte)  (error) {

	fmt.Println("strting push")

	m.currentShard = m.pushCount / m.recordsPerShard

	if(m.pushDataFP == nil || m.currentShard!=m.prevShard){

		m.prevShard = m.currentShard

	}
	
	return nil

}

func (m* QFileOp) OpenMetaFile(filepath string) (error) {

	

	data := make([]byte, 128)

	pushFP, err1 := os.OpenFile(filepath, os.O_RDWR,0777);
	popFP, err2 := os.OpenFile(filepath, os.O_RDWR,0777);

	if err1 != nil || err2!=nil {
        
        return errors.New("Error opening the metadata file: "+filepath)
    }

    pushFP.Read(data)


    objPush,res1 :=  BuildPushStruct(data)

    objPop,res2 := BuildPopStruct(data)
 
	currentMetaVersion := binhelp.Bin_to_int64(data[0:8])
	if(currentMetaVersion < minSupportedVersion){

		 return errors.New("Old MetaData file format not supported")
	}



	m.recordsPerShard = binhelp.Bin_to_int64(data[8:16])

	fmt.Println(m.recordsPerShard)


    if(res1==false || res2==false)    {

    	
    	return errors.New("Corrupt MetaData file: "+filepath)

    		// REturn error

    }


    m.pushId = objPush.pushId
    m.pushCount = objPush.pushCount
    m.popId = objPop.popId
    m.popCount = objPop.popCount
    m.headOffset = objPop.headOffset


    //assign file pointers
    m.pushFP = pushFP
    m.popFP = popFP

    return nil

}



func BuildPopStruct(data []byte) (PopStruct,bool) {


	evenPS:= PopStruct{}
	popHeadbin := data[64:128]
	evenPS.crc32 = binhelp.Bin_to_int64(popHeadbin[0:8])
	evenPS.popId = binhelp.Bin_to_int64(popHeadbin[8:16])
	evenPS.popCount = binhelp.Bin_to_int64(popHeadbin[16:24])
	evenPS.headOffset = binhelp.Bin_to_int64(popHeadbin[24:32])


	evenData := popHeadbin[8:32]

	// fmt.Println( int64(crc32.ChecksumIEEE(evenData)))
	// fmt.Println(evenPS.crc32)


	oddPS:= PopStruct{}
	oddPS.crc32 = binhelp.Bin_to_int64(popHeadbin[32:40])
	oddPS.popId = binhelp.Bin_to_int64(popHeadbin[40:48])
	oddPS.popCount = binhelp.Bin_to_int64(popHeadbin[48:56])	
	oddPS.headOffset = binhelp.Bin_to_int64(popHeadbin[56:64])

	oddData := popHeadbin[40:64]


	// fmt.Println( int64(crc32.ChecksumIEEE(oddData)))
	// fmt.Println(oddPS.crc32)

	if(ValidCRC32(evenPS.crc32,evenData) && evenPS.popId > oddPS.popId){

		return evenPS,true

	} else if(ValidCRC32(oddPS.crc32,oddData) && oddPS.popId > evenPS.popId){

		return oddPS,true

	} else if(ValidCRC32(evenPS.crc32,evenData) && ValidCRC32(oddPS.crc32,oddData) && evenPS.popId == oddPS.popId){

		return oddPS,true

	} else if (ValidCRC32(evenPS.crc32,evenData)){

		return evenPS,true

	} else if(ValidCRC32(oddPS.crc32,oddData)){

		return oddPS,true

	} else{

		return evenPS,false
	}



}

func BuildPushStruct(data []byte) (PushStruct,bool) {


	evenPS:= PushStruct{}
	pushHeadbin := data[16:64]
	evenPS.crc32 = binhelp.Bin_to_int64(pushHeadbin[0:8])
	evenPS.pushId = binhelp.Bin_to_int64(pushHeadbin[8:16])
	evenPS.pushCount = binhelp.Bin_to_int64(pushHeadbin[16:24])

	// fmt.Println( int64(crc32.ChecksumIEEE(pushHeadbin[8:24])))
	// fmt.Println(evenPS.crc32)

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

	//var version int64 = 1
		
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

