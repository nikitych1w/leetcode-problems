package main

func findDuplicate(nums []int) int {
	dict := map[int]int{}

	for _, el := range nums {
		dict[el]++
		if dict[el] > 1 {
			return el
		}
	}

	return 0
}
