package main

import "fmt"

type Training struct {
	Mongo Mongo
}

func (t *Training) run() {
	maxGenBot := t.Mongo.getMaxGenBot()
	gen := maxGenBot.Gen
	fmt.Println("start:", gen)
	for {
		bots := t.Mongo.getBotsByGen(gen)
		endChan := make(chan string)
		for i := 0; i < len(bots); i++ {
			for k := i + 1; k < len(bots); k++ {

				game := Game{}
				game.init()
				go game.competition(&bots[i], &bots[k], t.Mongo, endChan)
			}
		}

		for i := 0; i < int(float64(len(bots)*(len(bots)-1)/2)*0.99); i++ {
			_ = <-endChan
		}

		gen++
		fmt.Println("createNewGen:", gen)
		createNewGen(t.Mongo, gen)
	}
}
