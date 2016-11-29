package main

import "github.com/gin-gonic/gin"

import "github.com/evnix/ashtra/web"

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/boltdb/bolt"
	"io/ioutil"
)

func main() {

	fmt.Print(" ")
	log.Info("starting Ashtra...")
	db, err := bolt.Open("ashtra.db", 0600, nil)
	db.Update(func(tx *bolt.Tx) error {
		tx.CreateBucketIfNotExists([]byte("BB")) //basebucket
		b := tx.Bucket([]byte("BB"))
		err := b.Put([]byte("answer"), []byte("42"))
		return err
	})

	byteArray, err := ioutil.ReadFile("res/logo.txt")

	if err == nil {

		fmt.Print(string(byteArray))

	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/", ashtraweb.Index)

	r.Static("/web", "./web")

	if true {

		log.SetLevel(log.InfoLevel)
	}

	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")

	r.Run() // listen and server on 0.0.0.0:8080
}
