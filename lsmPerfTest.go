package main

import (
	"fmt"
	"log"
	"time"
	"os"
	"strconv"
	"runtime"
	"runtime/debug"
	"io/ioutil"
	 "github.com/syndtr/goleveldb/leveldb"
)

//

func main() {

	fmt.Println("testing LSM(leveldb) PerfInsert...")
	
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

    db, err := leveldb.OpenFile("testlevel.db", nil)
    if err != nil {
        log.Fatal(err)
    }
    
    defer db.Close()
    
    

	for m := 0; m < 20; m++ {

		start := time.Now()
		elapsed := time.Since(start)

		for i = 0; i < num; i++ {
		
		
		    if(i%10000==0){
		    
		    
		            elapsed = time.Since(start)
		            log.Printf("%d process at .. %f", i,elapsed.Seconds())
		    }


                    db.Put([]byte(strconv.FormatInt(i,10)), []byte(workload),nil)

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

			db.Delete([]byte(strconv.FormatInt(i,10)),nil)

		}

		elapsed = time.Since(start)
		log.Printf("pop took %s", elapsed)
		fmt.Println("pop took ", elapsed.Seconds())
		runtime.GC()
		debug.FreeOSMemory()
		time.Sleep(1 * time.Second)

	}



}
