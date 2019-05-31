package main

import (
	"encoding/json"
	"fmt"
)

type Bot struct {
	Name            string
	FeatureTypeList []ChessType
	Gen             int
	Father          []string
	Record          Record
}

type Record struct {
	Total     int
	Win       int
	Lose      int
	MateFirst []string
	MateLast  []string
}

func (b *Bot) loadJSON(jsonStr []byte) {
	jsonMap := make(map[string]interface{})
	if err := json.Unmarshal([]byte(jsonStr), &jsonMap); err != nil {
		panic(err)
	}
	b.Name = jsonMap["name"].(string)
	if feature, ok := jsonMap["feature"].(map[string]interface{}); ok {
		newChessType := []ChessType{}
		for _, chessType := range chessTypeList {
			newChessType = append(newChessType, ChessType{
				chessType.Name,
				chessType.Structure,
				feature[chessType.Name].(float64),
			})
		}
		b.FeatureTypeList = newChessType
	}
}

func (b *Bot) play(checkerboard [][]int8, user int8) [][]int8 {
	bestPointRecord := selectionBud(checkerboard, user, b.FeatureTypeList)
	checkerboard[bestPointRecord[len(bestPointRecord)-1].X][bestPointRecord[len(bestPointRecord)-1].Y] = user
	checkerboard[bestPointRecord[len(bestPointRecord)-2].X][bestPointRecord[len(bestPointRecord)-2].Y] = user
	// printCheckerboard(checkerboard)
	// fmt.Println(bestPointRecord[0].val * -1)
	fmt.Println(bestPointRecord)
	// fmt.Println(gameJudge(checkerboard))
	return checkerboard
}
