package server
import ("testing"
		//"fmt"
		"time")
//import "fmt"



func TestStart(t *testing.T){

	var  a=1;

	proc := QProc{}	
	proc.Start()
	a++;

	time.Sleep(10*time.Second)

	//fmt.Println("before push")
	proc.Push()



		time.Sleep(10*time.Second)
}


// func TestPush(){

// }


// func TestPop(){

// }


