package main

func searchMatrix(matrix [][]int, target int) bool {
	M := len(matrix)

	if M != 0 {
		N := len(matrix[0])

		for i := 0; i < M; i++ {
			for j := 0; j < N; j++ {

				currentNum := matrix[i][j]

				if currentNum < target {
					continue
				} else if currentNum > target {
					break
				} else {
					return true
				}
			}
		}
	}

	return false
}
