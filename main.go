package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	checkerboard := [][]int8{
		[]int8{0, 0, 0, 0, 0, 0, 0, 0},
		[]int8{0, 0, 0, 0, 0, 0, 0, 0},
		[]int8{0, 0, 0, 0, 0, 0, 0, 0},
		[]int8{0, 0, 0, 0, 0, 0, 0, 0},
		[]int8{0, 0, 1, 0, 1, 0, 0, 0},
		[]int8{0, 0, 0, -1, 0, 1, -1, 0},
		[]int8{0, 0, 0, 0, 0, 0, 1, 0},
		[]int8{0, 0, 0, 0, 0, 0, 0, 0},
	}
	jsonStr := []byte(`
    {"name":"e897e27d87972891811d2e3554f24211","feature":{"6L":900000000000,"5L:11":1000,"5L:01":900,"5J1":900,"5P1":900,"4L:22":900,"4L:12":900,"4L:02":900,"4L:11":900,"4J1:11":900,"4J1:01":900,"4J2":900,"4J1P1":900,"4P2":900,"4J1S1":900,"3L:33":900,"3L:23":900,"3L:13":900,"3L:03":900,"3L:22":900,"3L:12":900,"3J1:22":900,"3J1:12":900,"3J1:02":900,"3J1:11":900,"3J2:11":900,"3J2:01":900,"3J3":900,"3J1P1:11":900,"3J1P1:01":900,"3J2P1":900,"2L:44":900,"2L:34":900,"2L:24":900,"2L:14":900,"2L:04":900,"2L:33":900,"2L:23":900,"2L:13":900,"2L:22":900,"2J1:33":900,"2J1:23":900,"2J1:13":900,"2J1:03":900,"2J1:22":900,"2J1:12":900,"2J2:22":900,"2J2:12":900,"2J2:02":900,"2J2:11":900,"2J3:11":900,"2J3:01":900,"2J4":900}}
  `)
	jsonMap := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonStr), &jsonMap)
	if err != nil {
		panic(err)
	}
	maxVal := 0.00
	if feature, ok := jsonMap["feature"].(map[string]interface{}); ok {
		for x := 0; x < size; x++ {
			for y := 0; y < size; y++ {
				if checkerboard[x][y] == 0 {
					checkerboard[x][y] = 1
					val := valuation(checkerboard, 1, feature)
					if val > maxVal {
						maxVal = val
						fmt.Println(x, y)
					}
					checkerboard[x][y] = 0
				}
			}
		}
		fmt.Println(maxVal)
	}

}
