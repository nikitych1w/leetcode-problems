package main

func singleNumber(nums []int) int {
	res := 0

	if len(nums) > 0 {
		res = nums[0]

		for i := 1; i < len(nums); i++ {
			res = res ^ nums[i]
		}
	}

	return res
}
