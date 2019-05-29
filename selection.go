package main

import (
	"fmt"
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

func selection(checkerboard [][]int8, user int8, feature map[string]interface{}) {
	selectPoint := Points{}
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			effective := false
			for k := 1; k < 6; k++ {
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
				selfVal := valuation(checkerboard, user, feature)
				oppVal := valuation(checkerboard, user*-1, feature)
				selectPoint = append(selectPoint, Point{x, y, selfVal - oppVal})
				checkerboard[x][y] = 0
			}
		}
	}
	sort.Sort(selectPoint)
	fmt.Println(selectPoint)
}
