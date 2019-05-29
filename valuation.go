package main

import (
	"errors"
)

var size = 8

func valuation(checkerboard [][]int8, user int8, feature map[string]interface{}) float64 {
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
				if chessTypeVal, err := getBestType(temp, feature); err == nil {
					val += chessTypeVal
				}
				temp = ""
			}
		}
		if temp != "" {
			if chessTypeVal, err := getBestType(temp, feature); err == nil {
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
				if chessTypeVal, err := getBestType(temp, feature); err == nil {
					val += chessTypeVal
				}
				temp = ""
			}
		}
		if temp != "" {
			if chessTypeVal, err := getBestType(temp, feature); err == nil {
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
				if chessTypeVal, err := getBestType(temp, feature); err == nil {
					val += chessTypeVal
				}
				temp = ""
			}
			x--
			y++
		}
		if temp != "" {
			if chessTypeVal, err := getBestType(temp, feature); err == nil {
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
				if chessTypeVal, err := getBestType(temp, feature); err == nil {
					val += chessTypeVal
				}
				temp = ""
			}
			x++
			y++
		}
		if temp != "" {
			if chessTypeVal, err := getBestType(temp, feature); err == nil {
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

func getBestType(chessStructure string, feature map[string]interface{}) (float64, error) {

	chessTypes := getChessTypes(chessStructure)
	if len(chessTypes) == 0 {
		return 0, errors.New("No type")
	}
	// fmt.Println(chessStructure)
	// fmt.Println(chessTypes)

	// chessType := chessTypes[0]
	chessTypeVal := feature[chessTypes[0]].(float64)
	for index := 1; index < len(chessTypes); index++ {
		if feature[chessTypes[index]].(float64) > chessTypeVal {
			// chessType = chessTypes[index]
			chessTypeVal = feature[chessTypes[index]].(float64)
		}
	}
	return chessTypeVal, nil
}
