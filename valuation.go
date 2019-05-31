package main

import (
	"errors"
	"strings"
)

var size = 19

type Valuation struct {
	FeatureTypeList []ChessType
}

func (v *Valuation) all(checkerboard [][]int8, user int8) float64 {
	val := 0.00
	// x軸
	for x := 0; x < size; x++ {
		temp := ""
		for y := 0; y < size; y++ {
			if checkerboard[x][y] == 0 {
				temp = temp + "?"
			} else if checkerboard[x][y] == user {
				temp = temp + "A"
			} else {
				if chessTypeVal, err := v.getBestType(temp); err == nil {
					val += chessTypeVal
				}
				temp = ""
			}
		}
		if temp != "" {
			if chessTypeVal, err := v.getBestType(temp); err == nil {
				val += chessTypeVal
			}
			temp = ""
		}
	}

	// y軸
	for y := 0; y < size; y++ {
		temp := ""
		for x := 0; x < size; x++ {
			if checkerboard[x][y] == 0 {
				temp = temp + "?"
			} else if checkerboard[x][y] == user {
				temp = temp + "A"
			} else {
				if chessTypeVal, err := v.getBestType(temp); err == nil {
					val += chessTypeVal
				}
				temp = ""
			}
		}
		if temp != "" {
			if chessTypeVal, err := v.getBestType(temp); err == nil {
				val += chessTypeVal
			}
			temp = ""
		}
	}

	// 右斜
	startX := 5
	startY := 0
	for startY < size-5 {
		x := startX
		y := startY
		temp := ""
		for x >= 0 && y < size {
			// fmt.Println(x, y)
			if checkerboard[x][y] == 0 {
				temp = temp + "?"
			} else if checkerboard[x][y] == user {
				temp = temp + "A"
			} else {
				if chessTypeVal, err := v.getBestType(temp); err == nil {
					val += chessTypeVal
				}
				temp = ""
			}
			x--
			y++
		}
		if temp != "" {
			if chessTypeVal, err := v.getBestType(temp); err == nil {
				val += chessTypeVal
			}
			temp = ""
		}
		if startX == size-1 {
			startY++
		} else {
			startX++
		}
	}

	// 左斜
	startX = size - 5
	startY = 0
	for startY < size-5 {
		x := startX
		y := startY
		temp := ""
		for x < size && y < size {
			// fmt.Println(x, y)
			if checkerboard[x][y] == 0 {
				temp = temp + "?"
			} else if checkerboard[x][y] == user {
				temp = temp + "A"
			} else {
				if chessTypeVal, err := v.getBestType(temp); err == nil {
					val += chessTypeVal
				}
				temp = ""
			}
			x++
			y++
		}
		if temp != "" {
			if chessTypeVal, err := v.getBestType(temp); err == nil {
				val += chessTypeVal
			}
			temp = ""
		}
		if startX == 0 {
			startY++
		} else {
			startX--
		}
	}

	return val
	// fmt.Println(checkerboard[0][0])
}

func (v *Valuation) getBestType(chessStructure string) (float64, error) {

	matchType := []ChessType{}
	for _, chessType := range v.FeatureTypeList {
	matchLoop:
		for _, structure := range chessType.Structure {
			matchIndex := strings.Index(chessStructure, structure)
			if matchIndex != -1 {
				matchType = append(matchType, chessType)
				break matchLoop
			}
		}
	}

	if len(matchType) == 0 {
		return 0, errors.New("No type")
	}
	// fmt.Println(chessStructure)
	// fmt.Println(chessTypes)

	// chessType := chessTypes[0]
	chessTypeVal := matchType[0].Val
	for index := 1; index < len(matchType); index++ {
		if matchType[index].Val > chessTypeVal {
			// chessType = chessTypes[index]
			chessTypeVal = matchType[index].Val
		}
	}
	return chessTypeVal, nil
}
