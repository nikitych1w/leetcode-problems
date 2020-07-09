package main

import (
	"sort"
)

func groupThePeople(groupSizes []int) [][]int {

	dict := map[int][]int{}

	for i, el := range groupSizes {
		dict[el] = append(dict[el], i)
	}

	keys := make([]int, 0, len(dict))
	for k := range dict {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	var res [][]int

	for _, groupSize := range keys {
		if len(dict[groupSize]) > groupSize {

			startIdx := 0
			batchSize := len(dict[groupSize]) / groupSize

			for batch := 0; batch < batchSize; batch++ {
				res = append(res, dict[groupSize][startIdx:startIdx+groupSize])
				startIdx += groupSize
			}

		} else {
			res = append(res, dict[groupSize])
		}
	}

	return res
}

//func main() {
//	/**
//	Input: groupSizes = [3,3,3,3,3,1,3]
//	Output: [[5],[0,1,2],[3,4,6]]
//
//	Input: groupSizes = [2,1,3,3,3,2]
//	Output: [[1],[0,5],[2,3,4]]
//	*/
//
//	//src := []int {3,3,3,3,3,1,3}
//	src := []int {2,1,3,3,3,2}
//	fmt.Println(groupThePeople(src))
//
//}
