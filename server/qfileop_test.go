package server
import ("testing"
		"fmt")


func TestCreate(t *testing.T){


	fmt.Println("testing...")

	CreateMetaFile("test/testq",10)

	qfp:= QFileOp{};

	ret := qfp.OpenMetaFile("test/testq")


	if(ret!=nil){

		panic(ret)
	}



	var ret2 error

	str1:=""
	str2:=""

	uid:=""
	
fmt.Println("pushing...")

	for i:=0;i<15;i++ {

		ret2,uid = qfp.PushElement(1,2,3,[]byte("hello"))
		if(ret2!=nil){

			panic(ret2)
		}
		str1+=uid
	}




	record := Record{}
fmt.Println("popping...")
	for i:=0;i<15;i++ {

		ret2,record = qfp.PopElement()
		if(ret2!=nil){

			panic(ret2)
		}
		str2+=((record.uid))
	}


	if(str1==str2){

		fmt.Println("good..it works :)")
	}


	

}


func TestCreate1000(t *testing.T){


	fmt.Println("testing...")

	num:=1000000

	CreateMetaFile("test/test1000",int64(num))

	qfp:= QFileOp{};

	ret := qfp.OpenMetaFile("test/test1000")


	if(ret!=nil){

		panic(ret)
	}



	var ret2 error

	str1:=""
	str2:=""

	uid:=""
	
fmt.Println("pushing 10,000...")

	for i:=0;i<num;i++ {

		ret2,uid = qfp.PushElement(1,2,3,[]byte("hello"))

		if(i%10000==0){
			//fmt.Println(i)
		}

		if(ret2!=nil){

			panic(ret2)
		}
		str1=uid
	}




	record := Record{}
fmt.Println("popping 10,000...")
	for i:=0;i<num;i++ {

		ret2,record = qfp.PopElement()

		if(i%10000==0){
			//fmt.Println(i)
		}

		if(ret2!=nil){

			panic(ret2)
		}
		str2=((record.uid))
	}


	if(str1==str2){

		fmt.Println("good..it works :)")
	}


	

}




