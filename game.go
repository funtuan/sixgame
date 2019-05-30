package main

type Game struct {
	JudgeBot Bot
}

func (g *Game) init() {
	g.JudgeBot = Bot{}
	g.JudgeBot.loadJSON([]byte(`
    {"name":"e897e27d87972891811d2e3554f24211","feature":{"6L":900000000000,"5L:11":0,"5L:01":0,"5J1":0,"5P1":0,"4L:22":0,"4L:12":0,"4L:02":0,"4L:11":0,"4J1:11":0,"4J1:01":0,"4J2":0,"4J1P1":0,"4P2":0,"4J1S1":0,"3L:33":0,"3L:23":0,"3L:13":0,"3L:03":0,"3L:22":0,"3L:12":0,"3J1:22":0,"3J1:12":0,"3J1:02":0,"3J1:11":0,"3J2:11":0,"3J2:01":0,"3J3":0,"3J1P1:11":0,"3J1P1:01":0,"3J2P1":0,"2L:44":0,"2L:34":0,"2L:24":0,"2L:14":0,"2L:04":0,"2L:33":0,"2L:23":0,"2L:13":0,"2L:22":0,"2J1:33":0,"2J1:23":0,"2J1:13":0,"2J1:03":0,"2J1:22":0,"2J1:12":0,"2J2:22":0,"2J2:12":0,"2J2:02":0,"2J2:11":0,"2J3:11":0,"2J3:01":0,"2J4":0}}
  `))
}

func (g *Game) gameJudge(checkerboard [][]int8) int8 {

	if valuation(checkerboard, 1, g.JudgeBot.FeatureTypeList) > 100000000000 {
		return 1
	}
	if valuation(checkerboard, -1, g.JudgeBot.FeatureTypeList) > 100000000000 {
		return -1
	}

	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			if checkerboard[x][y] == 0 {
				//遊戲尚未結束
				return 0
			}
		}
	}
	//平手
	return 2
}

func (g *Game) start(bot1 Bot, bot2 Bot) int8 {
	checkerboard := openSgf("sixgameTest.sgf")
	var gameResult int8 = 0

	for g.gameJudge(checkerboard) == 0 {
		checkerboard = bot1.play(checkerboard, -1)
		printCheckerboard(checkerboard)
		if g.gameJudge(checkerboard) == 0 {
			checkerboard = bot2.play(checkerboard, 1)
			printCheckerboard(checkerboard)
		}
	}
	if g.gameJudge(checkerboard) != 2 {
		gameResult = gameResult - g.gameJudge(checkerboard)
	}

	checkerboard = openSgf("sixgameTest.sgf")

	for g.gameJudge(checkerboard) == 0 {
		checkerboard = bot2.play(checkerboard, -1)
		printCheckerboard(checkerboard)
		if g.gameJudge(checkerboard) == 0 {
			checkerboard = bot1.play(checkerboard, 1)
			printCheckerboard(checkerboard)
		}
	}
	if g.gameJudge(checkerboard) != 2 {
		gameResult = gameResult + g.gameJudge(checkerboard)
	}
	message = make(chan string)
	message <- "ok"
	return gameResult
}
