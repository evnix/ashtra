package server
import ("testing")


func TestCreate(t *testing.T){


	CreateMetaFile("testq",10)

	qfp:= QFileOp{};

	ret := qfp.OpenMetaFile("testq")


	if(ret!=nil){

		panic(ret)
	}



	var ret2 error


	//for i:=0;i<100;i++ {

		ret2 = qfp.PushElement(1,2,3,[]byte("hello"))

	//}


	if(ret2!=nil){

		panic(ret2)
	}

		qfp.PopElement()


}





