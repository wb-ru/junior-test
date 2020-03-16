package main

import "fmt"

func twoSum(nums []int, target int) []int {

	firstIndex := 0
	secondIndex := 0

	for i, _ := range nums {

		for j, _ := range nums {
			if nums[i]+nums[j] == target && i != j {
				firstIndex = j
				secondIndex = i

				return []int{firstIndex, secondIndex}
			}
		}
	}

	return []int{firstIndex, secondIndex}
}

func main() {

	nums := []int{2, 7, 11, 15}

	fmt.Println(twoSum(nums, 9))
}
