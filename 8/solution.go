package main

func FindArrWithMaxSum(arr []int) int {
	max := 0
	current := 0

	for i := range arr {
		current += arr[i]

		if current < 0 {
			current = 0
		} else if max < current {
			max = current
		}
	}

	return max
}
