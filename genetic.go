package main

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"time"
)

var pm = 0.12

func creator(mongo Mongo, count int) {
	rand.Seed(int64(time.Now().UnixNano()))
	for i := 0; i < count; i++ {
		bot := Bot{}
		bot.loadJSON([]byte(`
  	  {"name":"e897e27d87972891811d2e3554f24211","feature":{"6L":900000000000,"5L:11":100000,"5L:01":90000,"5J1":85000,"5P1":80000,"4L:22":75000,"4L:12":7000,"4L:02":6500,"4L:11":6000,"4J1:11":5500,"4J1:01":4500,"4J2":4000,"4J1P1":3900,"4P2":3800,"4J1S1":3700,"3L:33":3600,"3L:23":3500,"3L:13":3400,"3L:03":3300,"3L:22":3200,"3L:12":3100,"3J1:22":3000,"3J1:12":2900,"3J1:02":2800,"3J1:11":2700,"3J2:11":2600,"3J2:01":2500,"3J3":2400,"3J1P1:11":2300,"3J1P1:01":2200,"3J2P1":2100,"2L:44":2000,"2L:34":1900,"2L:24":1800,"2L:14":1700,"2L:04":1600,"2L:33":1500,"2L:23":1400,"2L:13":1300,"2L:22":1200,"2J1:33":1100,"2J1:23":1000,"2J1:13":900,"2J1:03":800,"2J1:22":700,"2J1:12":600,"2J2:22":600,"2J2:12":500,"2J2:02":500,"2J2:11":500,"2J3:11":400,"2J3:01":300,"2J4":200}}
  	`))
		bot.Name = getRandName()
		fmt.Println(bot.Name)
		for i := 1; i < len(bot.FeatureTypeList); i++ {
			bot.FeatureTypeList[i].Val = rand.Float64() * 100000
		}
		mongo.saveBot(bot)
	}
}

func crossover(bot1 Bot, bot2 Bot) Bot {
	rand.Seed(int64(time.Now().UnixNano()))
	bot := Bot{}
	bot.loadJSON([]byte(`
		{"name":"e897e27d87972891811d2e3554f24211","feature":{"6L":900000000000,"5L:11":100000,"5L:01":90000,"5J1":85000,"5P1":80000,"4L:22":75000,"4L:12":7000,"4L:02":6500,"4L:11":6000,"4J1:11":5500,"4J1:01":4500,"4J2":4000,"4J1P1":3900,"4P2":3800,"4J1S1":3700,"3L:33":3600,"3L:23":3500,"3L:13":3400,"3L:03":3300,"3L:22":3200,"3L:12":3100,"3J1:22":3000,"3J1:12":2900,"3J1:02":2800,"3J1:11":2700,"3J2:11":2600,"3J2:01":2500,"3J3":2400,"3J1P1:11":2300,"3J1P1:01":2200,"3J2P1":2100,"2L:44":2000,"2L:34":1900,"2L:24":1800,"2L:14":1700,"2L:04":1600,"2L:33":1500,"2L:23":1400,"2L:13":1300,"2L:22":1200,"2J1:33":1100,"2J1:23":1000,"2J1:13":900,"2J1:03":800,"2J1:22":700,"2J1:12":600,"2J2:22":600,"2J2:12":500,"2J2:02":500,"2J2:11":500,"2J3:11":400,"2J3:01":300,"2J4":200}}
	`))
	bot.Name = getRandName()
	bot.Father = []string{bot1.Name, bot2.Name}
	for i := 0; i < len(bot.FeatureTypeList); i++ {
		bot.FeatureTypeList[i].Val = (bot1.FeatureTypeList[i].Val + bot2.FeatureTypeList[i].Val) / 2
		if rand.Float64() < pm {
			bot.FeatureTypeList[i].Val = bot.FeatureTypeList[i].Val * (0.5 + rand.Float64()*1)
		}
	}

	return bot
}

func createNewGen(mongo Mongo, gen int) {
	bots := mongo.getBotsByGen(gen - 1)
	sum := len(bots)
	elites := int(float64(len(bots)) * 0.3)
	for i := 0; i < elites; i++ {
		bots[i].Gen = gen
		bots[i].Record = Record{}
		mongo.saveBot(bots[i])
		sum--
	}

elitesLoop:
	for i := 0; i < elites; i++ {
		for k := i + 1; k < elites; k++ {
			bot := crossover(bots[i], bots[k])
			bot.Gen = gen
			mongo.saveBot(bot)
			sum--
			if sum == 0 {
				break elitesLoop
			}
		}
	}
}

func getRandName() string {
	data := []byte(fmt.Sprintf("%d", rand.Int()))
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has)
	return md5str1
}
