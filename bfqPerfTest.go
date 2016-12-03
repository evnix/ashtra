package main

import (
	"fmt"
	"github.com/evnix/ashtra/server/qfileop"
	"log"
	"time"
	"os"
	"strconv"
	"runtime"
	"runtime/debug"
	"io/ioutil"
)

//

func main() {

	fmt.Println("testing BigFastQueue PerfInsert...")
	
	b, err := ioutil.ReadFile(os.Args[1]) 

    if err != nil {
    
        fmt.Println("error occured");
        log.Printf("error opening file")
    }
    
    num,_ := strconv.ParseInt(os.Args[2], 10, 64)
	doWork(b,num)



}

func doWork(workload []byte,num int64){

    var i int64

	qfileop.CreateMetaFile("test/testPerf", int64(num))

	qfp := qfileop.QFileOp{}

	ret := qfp.OpenMetaFile("test/testPerf")

	if ret != nil {

		panic(ret)
	}

	for m := 0; m < 20; m++ {

		start := time.Now()
		elapsed := time.Since(start)

		for i = 0; i < num; i++ {

			if(i%10000==0){
		    
		    
		            elapsed = time.Since(start)
		            log.Printf("%d process at .. %f", i,elapsed.Seconds())
		    }


			qfp.PushElement(1, 2, 3, workload)

		}

		log.Printf("%d push took %s", m,elapsed)
        fmt.Println("push took ", elapsed.Seconds())

		runtime.GC()
		debug.FreeOSMemory()
		time.Sleep(1 * time.Second)

		start = time.Now()

		for i = 0; i < num; i++ {


			if(i%1000==0){
		    
		            elapsed = time.Since(start)
		            log.Printf("%d process at ... %f", i,elapsed.Seconds())
		    }

			qfp.PopElement()

		}

		elapsed = time.Since(start)
		log.Printf("pop took %s", elapsed)
		fmt.Println("pop took ", elapsed.Seconds())
		runtime.GC()
		debug.FreeOSMemory()
		time.Sleep(1 * time.Second)

	}


}
