package main

import (
	"fmt"
	"github.com/evnix/ashtra/server/qfileop"
	"log"
	"time"
	"os"
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
    
	doWork(b)



}

func doWork(workload []byte){

    	num := 1000000

	qfileop.CreateMetaFile("test/testPerf", int64(num))

	qfp := qfileop.QFileOp{}

	ret := qfp.OpenMetaFile("test/testPerf")

	if ret != nil {

		panic(ret)
	}

	for m := 0; m < 20; m++ {

		start := time.Now()

		for i := 0; i < num; i++ {

			qfp.PushElement(1, 2, 3, workload)

		}

		elapsed := time.Since(start)
		log.Printf("%d push took %s", m,elapsed)
        fmt.Println("push took ", elapsed)

		runtime.GC()
		debug.FreeOSMemory()
		time.Sleep(1 * time.Second)

		start = time.Now()

		for i := 0; i < num; i++ {

			qfp.PopElement()

		}

		elapsed = time.Since(start)
		log.Printf("pop took %s", elapsed)
		fmt.Println("pop took ", elapsed)
		runtime.GC()
		debug.FreeOSMemory()
		time.Sleep(1 * time.Second)

	}


}
