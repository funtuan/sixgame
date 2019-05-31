package main

import (
	"crypto/tls"
	"log"
	"net"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Mongo struct {
	DB      *mgo.Database
	BotPool *mgo.Collection
}

func (m *Mongo) login() {
	m.DB = m.getDB()
	m.DB.Login("sixgame", "TCfTmXhONIWug0bL")
	m.BotPool = m.DB.C("botPool")
}

func (m *Mongo) saveBot(bot Bot) {
	err := m.BotPool.Insert(&bot)
	if err != nil {
		panic(err)
	}
}

func (m *Mongo) getBotByName(name string) Bot {
	bot := Bot{}
	err := m.BotPool.Find(bson.M{"name": name}).Sort("-gen").One(&bot)
	if err != nil {
		log.Fatal(err)
	}
	return bot
}

func (m *Mongo) getBotsByGen(gen int) []Bot {
	bots := []Bot{}
	err := m.BotPool.Find(bson.M{"gen": gen}).Sort("-record.win").All(&bots)
	if err != nil {
		log.Fatal(err)
	}
	return bots
}

func (m *Mongo) getMaxGenBot() Bot {
	bot := Bot{}
	err := m.BotPool.Find(bson.M{}).Sort("-gen").One(&bot)
	if err != nil {
		log.Fatal(err)
	}
	return bot
}

func (m *Mongo) updateBot(bot Bot) {
	err := m.BotPool.Update(bson.M{"name": bot.Name, "gen": bot.Gen}, &bot)
	if err != nil {
		log.Fatal(err)
	}
}

func (m *Mongo) getDB() *mgo.Database {

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
