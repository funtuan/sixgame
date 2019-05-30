package main

import (
	"sort"
)

type Point struct {
	X   int
	Y   int
	val float64
}
type Points []Point

func (p Points) Len() int { return len(p) }
func (p Points) Less(i, j int) bool {
	return p[i].val > p[j].val
}
func (p Points) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func selectionPoints(checkerboard [][]int8, user int8, featureTypeList []ChessType) Points {
	selectPoint := Points{}
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			effective := false
			for k := 1; k < 4; k++ {
				if x+k < size {
					if checkerboard[x+k][y] != 0 {
						effective = true
					}
					if y+k < size && checkerboard[x+k][y+k] != 0 {
						effective = true
					}
					if y-k >= 0 && checkerboard[x+k][y-k] != 0 {
						effective = true
					}
				}
				if x-k >= 0 {
					if checkerboard[x-k][y] != 0 {
						effective = true
					}
					if y+k < size && checkerboard[x-k][y+k] != 0 {
						effective = true
					}
					if y-k >= 0 && checkerboard[x-k][y-k] != 0 {
						effective = true
					}
				}
				if y+k < size {
					if checkerboard[x][y+k] != 0 {
						effective = true
					}
				}
				if y-k >= 0 {
					if checkerboard[x][y-k] != 0 {
						effective = true
					}
				}
			}
			if effective && checkerboard[x][y] == 0 {
				checkerboard[x][y] = user
				selfVal := valuation(checkerboard, user, featureTypeList)
				oppVal := valuation(checkerboard, user*-1, featureTypeList)
				selectPoint = append(selectPoint, Point{x, y, selfVal - oppVal})
				checkerboard[x][y] = 0
			}
		}
	}
	sort.Sort(selectPoint)
	// fmt.Println(selectPoint)
	return selectPoint
}

func selectionNode(checkerboard [][]int8, user int8, featureTypeList []ChessType, depth int) (float64, Points) {
	if depth%2 == 0 {
		user = user * -1
	}
	depth--
	selectPoint := selectionPoints(checkerboard, user, featureTypeList)
	// fmt.Println(selectPoint)
	if depth == 0 {
		if user == 1 {
			return selectPoint[0].val, Points{selectPoint[0]}
		} else {
			return selectPoint[0].val * -1, Points{selectPoint[0]}
		}
	}
	if selectPoint[0].val > 1000000000 || selectPoint[0].val < -1000000000 {
		if user == 1 {
			return selectPoint[0].val, Points{selectPoint[0]}
		} else {
			return selectPoint[0].val * -1, Points{selectPoint[0]}
		}
	}

	var bestVal float64
	var bestPointRecord Points
	branchs := 1
	// if depth == 10 {
	// 	branchs = 2
	// }
	for b := 0; b < branchs; b++ {
		nodeCheckerboard := make([][]int8, size)
		for i := range checkerboard {
			nodeCheckerboard[i] = make([]int8, len(checkerboard[i]))
			copy(nodeCheckerboard[i], checkerboard[i])
		}
		nodeCheckerboard[selectPoint[b].X][selectPoint[b].Y] = user
		// printCheckerboard(nodeCheckerboard)
		nodeVal, pointRecord := selectionNode(nodeCheckerboard, user, featureTypeList, depth)
		if b == 0 {
			bestVal = nodeVal
			pointRecord = append(pointRecord, selectPoint[b])
			bestPointRecord = pointRecord
		} else if bestVal < nodeVal {
			bestVal = nodeVal
			pointRecord = append(pointRecord, selectPoint[b])
			bestPointRecord = pointRecord
		}
	}
	if user == 1 {
		return bestVal, bestPointRecord
	} else {
		return bestVal * -1, bestPointRecord
	}
}

func selectionBud(checkerboard [][]int8, user int8, featureTypeList []ChessType) Points {
	selectPoint := selectionPoints(checkerboard, user, featureTypeList)
	// fmt.Println(selectPoint)
	var bestVal float64
	var bestPointRecord Points
	for b := 0; b < 3; b++ {
		nodeCheckerboard := make([][]int8, size)
		for i := range checkerboard {
			nodeCheckerboard[i] = make([]int8, len(checkerboard[i]))
			copy(nodeCheckerboard[i], checkerboard[i])
		}
		nodeCheckerboard[selectPoint[b].X][selectPoint[b].Y] = user
		nodeVal, pointRecord := selectionNode(nodeCheckerboard, user, featureTypeList, 3)
		pointRecord = append(pointRecord, selectPoint[b])
		if b == 0 {
			bestVal = nodeVal
			bestPointRecord = pointRecord
		} else if bestVal < nodeVal {
			bestVal = nodeVal
			bestPointRecord = pointRecord
		}
	}
	return bestPointRecord
}
