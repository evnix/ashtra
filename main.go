package main

import "github.com/gin-gonic/gin"

import "github.com/evnix/ashtra/web"

import ("fmt"
		"io/ioutil"
		log "github.com/Sirupsen/logrus"
		)

func main() {

	fmt.Print(" ")
	log.Info("starting Ashtra...")

	byteArray, err := ioutil.ReadFile("res/logo.txt")


	if(err==nil){

			fmt.Print(string(byteArray))

	}
    
       r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })

    r.GET("/", ashtraweb.Index)

    r.Static("/assets", "./web")


    if (true){

    	    log.SetLevel(log.InfoLevel)
    }


	log.WithFields(log.Fields{
	    "animal": "walrus",
	  }).Info("A walrus appears")

    r.Run() // listen and server on 0.0.0.0:8080
}

  