package main

import ("github.com/evnix/ashtra/server/qfileop"
	"log"
	"fmt"
	"time")


func main(){

	fmt.Println("testing BigFastQueue PerfInsert...")

	num:=1000000

	qfileop.CreateMetaFile("test/testPerf",int64(num))

	qfp:= qfileop.QFileOp{};

	ret := qfp.OpenMetaFile("test/testPerf")



	if(ret!=nil){

		panic(ret)
	}


	for m:=0;m<10;m++ {

	    start := time.Now()

		for i:=0;i<num;i++ {

			 qfp.PushElement(1,2,3,[]byte("hello"))

		}



	    elapsed := time.Since(start)
		log.Printf("push took %s", elapsed)


		    start = time.Now()

		for i:=0;i<num;i++ {

			 qfp.PopElement()

		}



	    elapsed = time.Since(start)
		log.Printf("pop took %s", elapsed)
	
	}

}

