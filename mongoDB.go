package main

import (
	"crypto/tls"
	"log"
	"net"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Name string `bson:"name"`
	Age  int    `bson:"age"`
}

func saveBot(bot Bot) {
	mydb := getDB()
	mydb.Login("sixgame", "TCfTmXhONIWug0bL")
	c := mydb.C("botPool")
	err := c.Insert(&bot)
	if err != nil {
		panic(err)
	}
}

func getBot(name string) Bot {
	mydb := getDB()
	mydb.Login("sixgame", "TCfTmXhONIWug0bL")
	c := mydb.C("botPool")
	bot := Bot{}
	err := c.Find(bson.M{"name": name}).One(&bot)
	if err != nil {
		log.Fatal(err)
	}
	return bot
}

func getDB() *mgo.Database {

	tlsConfig := &tls.Config{}

	dialInfo := &mgo.DialInfo{}
	dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
		conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
		return conn, err
	}
	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		panic(err)
	}
	db := session.DB("sixgame") //root user is created in the admin authentication database and given the role of root.
	return db
}
