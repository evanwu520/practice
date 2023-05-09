package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	const ylen, xLen int = 5, 4

	var directMap map[string][]int = make(map[string][]int)

	directMap["up"] = []int{1, 0}
	directMap["down"] = []int{-1, 0}
	directMap["right"] = []int{0, 1}
	directMap["left"] = []int{0, -1}

	var arrays = [ylen][xLen]string{
		{"b", "a", "3", "3"},
		{"b", "a", "1", "1"},
		{"b", "a", "4", "1"},
		{"b", "a", "4", "1"},
		{"2", "2", "a", "a"},
	}

	var tempMap map[string]map[int][]string = make(map[string]map[int][]string)

	for y := 0; y < ylen; y++ {

		for x := 0; x < xLen; x++ {

			// fmt.Println(y, x, arrays[y][x])

			s := arrays[y][x]

			if _, exist := tempMap[s]; !exist {

				// 表示s不存在，記錄

				v := fmt.Sprintf("%d,%d", y, x)
				tempMap[s] = make(map[int][]string)
				tempMap[s][0] = append(tempMap[s][0], v)

			} else {

				have := false
				index := 0

				// 表示s存在，檢查是否有連線
				for key, lines := range tempMap[s] {

					index = key

					for _, point := range lines {

						isConnect := false

						old := strings.Split(point, ",")
						oldY, _ := strconv.Atoi(old[0])
						oldX, _ := strconv.Atoi(old[1])

						for _, direct := range directMap {

							addY := direct[0]
							addX := direct[1]

							if int(oldY)+addY == y && oldX+addX == x {
								isConnect = true
								v := fmt.Sprintf("%d,%d", y, x)
								tempMap[s][key] = append(tempMap[s][key], v)
								have = true
								break
							}
						}

						if isConnect {
							break
						}
					}

				}

				// new line
				if !have {
					v := fmt.Sprintf("%d,%d", y, x)
					tempMap[s][index+1] = append(tempMap[s][index+1], v)
				}

			}
		}
	}

	//  原圖
	for _, v := range arrays {
		fmt.Println(v)
	}

	for key, lines := range tempMap {

		for i, line := range lines {

			fmt.Println(fmt.Sprintf("%v的第%d條線", key, i+1))

			chart := [ylen][xLen]string{}

			// init value "0"
			for y := 0; y < ylen; y++ {

				for x := 0; x < xLen; x++ {
					chart[y][x] = "0"
				}
			}

			// fill key
			for _, point := range line {

				v := strings.Split(point, ",")
				vY, _ := strconv.Atoi(v[0])
				vX, _ := strconv.Atoi(v[1])
				chart[vY][vX] = key
			}

			for _, v := range chart {
				fmt.Println(v)
			}
		}
	}
}
