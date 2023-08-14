package main

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left != right {
		center := (left + right + 1) / 2

		if nums[center] == target {
			return center
		}

		if nums[center] < target {
			left = center
		}

		if nums[center] > target {
			right = center - 1
		}
	}

	if nums[left] == target {
		return left
	}

	return -1
}
