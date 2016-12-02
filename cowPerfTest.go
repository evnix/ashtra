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
	 "github.com/boltdb/bolt"
)

//

func main() {

	fmt.Println("testing COW(Bolt) PerfInsert...")
	
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

    db, err := bolt.Open("test/testcow.db", 0777, nil)
    if err != nil {
        log.Fatal(err)
    }
    
    defer db.Close()
    
    tx, err2 := db.Begin(true)
    if err2 != nil {
        log.Fatal(err2)
    }
    defer tx.Rollback()

    // Use the transaction...
    b, err3 := tx.CreateBucket([]byte("MyBucket"))

    if err3 != nil {
        log.Fatal(err2)
    }

	for m := 0; m < 20; m++ {

		start := time.Now()
		elapsed := time.Since(start)

		for i = 0; i < num; i++ {
		
		
		    if(i%10000==0){
		    
		    
		            elapsed = time.Since(start)
		            log.Printf("%d process at .. %f", i,elapsed.Seconds())
		    }


                    b.Put([]byte(strconv.FormatInt(i,10)), []byte(workload))

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

			b.Delete([]byte(strconv.FormatInt(i,10)))

		}

		elapsed = time.Since(start)
		log.Printf("pop took %s", elapsed)
		fmt.Println("pop took ", elapsed.Seconds())
		runtime.GC()
		debug.FreeOSMemory()
		time.Sleep(1 * time.Second)

	}
	
	    if err9 := tx.Commit(); err9 != nil {
        log.Fatal(err9)
    }


}
