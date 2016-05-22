package server
import ("io/ioutil"
		"fmt"
		"os"
		"github.com/evnix/ashtra/server/binhelp"
		"hash/crc32"
		)


func CreateQueueMetaFile(filepath string,  recordsPerShard int64){

	var version int64 = 255
	var checksum int64 = 0

	


		
	//Even Push ID 64 bits + Push Count 64 bits
	pushHeaderPart := append(binhelp.Int64_to_bin(0),binhelp.Int64_to_bin(0)...)
	pushHeaderCRC32 := int64(crc32.ChecksumIEEE(pushHeaderPart))


	pushHeader_part := append(binhelp.Int64_to_bin(pushHeaderCRC32), pushHeaderPart...)
	

	pushHeader := append(pushHeader_part,pushHeader_part...)

	t := append(binhelp.Int64_to_bin(version),binhelp.Int64_to_bin(recordsPerShard)...)

	x:=crc32.ChecksumIEEE(t) 

	checksum = int64( x)

	fmt.Println(pushHeader)
	fmt.Println(t)
	fmt.Println(checksum)


	ioutil.WriteFile(filepath, []byte("Hello World"), os.ModeAppend | 0777)

}