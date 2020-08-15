package main

func average(salary []int) float64 {
	max := 0
	min := 1<<63 - 1
	sum := 0

	for i, _ := range salary {
		if salary[i] > max {
			max = salary[i]
		}

		if salary[i] < min {
			min = salary[i]
		}
		sum += salary[i]
	}

	sum -= min
	sum -= max

	return float64(sum) / float64(len(salary)-2)
}
