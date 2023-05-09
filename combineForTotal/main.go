package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var keepMap map[string]int = make(map[string]int)

func main() {

	fmt.Println("-----------------Exhausting-------------------------")

	Exhausting()

	fmt.Println("---------------CombineForTotal-----------------------------")

	CombineForTotal()

}

func Exhausting() {

	var n int = 5
	var target int = 5

	solve(n, target, "")

	for k, v := range keepMap {

		if v == target {
			fmt.Println(k)
		}
	}

}

func solve(n, target int, findedStr string) (ans string) {

	// fmt.Println(value)

	tmp := strings.Split(findedStr, " + ")

	var total int
	for _, v := range tmp {
		intVar, _ := strconv.Atoi(v)
		total = total + intVar
	}

	if total == target {
		keepMap[findedStr] = total
	}

	for i := 1; i <= n; i++ {

		if n >= 1 {

			var s string

			if findedStr == "" {
				s = fmt.Sprintf("%v", i)
			} else {
				s = fmt.Sprintf("%v + %v", findedStr, i)
			}

			// fmt.Println(s)
			solve(n-1, target, s)

		}

	}

	return ans
}

func CombineForTotal() {

	var saveMap map[string]int = make(map[string]int)

	var nums []int = []int{1, 2, 3, 4, 5}

	var target int = 5

	for _, num := range nums {

		for i := 1; i <= len(nums)-1; i++ {

			if num == target {
				key := combineArray(0, 0, num)

				if _, exist := saveMap[key]; !exist {
					saveMap[key] = num
				}
				break
			}

			if i*num <= 5 {
				v := target - (i * num)

				if v > 0 {
					key := combineArray(num, i, v)

					if _, exist := saveMap[key]; !exist {
						saveMap[key] = 5
					}
				}
			}
		}
	}

	for key, _ := range saveMap {
		fmt.Println(key)
	}
}

func combineArray(base, loop int, modV int) string {

	var array []int

	for i := 1; i <= loop; i++ {
		array = append(array, base)
	}

	array = append(array, modV)

	sort.Slice(array, func(i, j int) bool {
		return array[i] < array[j]
	})

	var tmpS string

	for i, s := range array {

		if i == 0 {
			tmpS = fmt.Sprintf("%v", s)
		} else {
			tmpS = fmt.Sprintf("%v+%v", tmpS, s)
		}
	}

	return tmpS

}
