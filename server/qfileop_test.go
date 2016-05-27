package server
import ("testing")


func TestCreate(t *testing.T){


	CreateMetaFile("testq",10)

	qfp:= QFileOp{};

	ret := qfp.OpenMetaFile("testq")


	if(ret!=nil){

		panic(ret)
	}


	ret2 := qfp.PushElement([]byte("hello"))

	if(ret2!=nil){

		panic(ret2)
	}


}


