package main

import "testing"

func TestSum(t *testing.T) {

	t.Run("Given nums = [2, 7, 11, 15], target = 9", func(t *testing.T) {
		nums := []int{2, 7, 11, 15}
		target := 9
		i := twoSum(nums, target)
		got := 9

		if (got != target) || (i[0] == i[1]) {
			t.Errorf("got %d target %d given, %v", got, target, nums)
		}
	})

	t.Run("Given nums = [5, 3, 1, 6], target = 7", func(t *testing.T) {
		nums := []int{5, 3, 1, 6}
		target := 7
		i := twoSum(nums, target)
		got := nums[i[0]] + nums[i[1]]

		if (got != target) || (i[0] == i[1]) {
			t.Errorf("got %d target %d given, %v", got, target, nums)
		}
	})

	t.Run("Given nums = [1, 2], target = 3", func(t *testing.T) {
		nums := []int{1, 2}
		target := 3
		i := twoSum(nums, target)
		got := nums[i[0]] + nums[i[1]]

		if (got != target) || (i[0] == i[1]) {
			t.Errorf("got %d target %d given, %v", got, target, nums)
		}
	})
	t.Run("Given nums = [1, 1], target = 2", func(t *testing.T) {
		nums := []int{1, 1}
		target := 2
		i := twoSum(nums, target)
		got := nums[i[0]] + nums[i[1]]

		if (got != target) || (i[0] == i[1]) {
			t.Errorf("got %d target %d given, %v", got, target, nums)
		}
	})
	t.Run("Given nums = [3, 2, 1, 1], target = 2", func(t *testing.T) {
		nums := []int{3, 2, 1, 1}
		target := 2
		i := twoSum(nums, target)
		got := nums[i[0]] + nums[i[1]]

		if (got != target) || (i[0] == i[1]) {
			t.Errorf("got %d target %d given, %v", got, target, nums)
		}
	})
}
