package server

import( "time"
		"fmt")


type QProc struct {
    name string;
    req chan int
}


func(q* QProc)  Start(string Qname) int{

	c := make(chan int)
	quit := make(chan int)
	go run(c,quit)
	
	q.req=c;

	fmt.Print("here")


	for i:=0;i<4;i++ {

		c <- 1
		fmt.Println("inLoop")
		time.Sleep(1*time.Second)
	}

	
	return 0
}




func run(c,quit chan int){


	i :=0;
	i++


	for {

		fmt.Println("-XXX--")
		select {
		case v := <-c:
				fmt.Print("inVVVV")
				fmt.Print(v)
		// default:
		// 	fmt.Println("--C-")
    // receiving from c would block
		}

			//time.Sleep(1*time.Second)
	}

}


func(q* QProc) Push(){

	q.req <- 1
		fmt.Println("-pushed--")

}