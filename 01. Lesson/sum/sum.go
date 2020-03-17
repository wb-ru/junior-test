package sum

func twoSum(nums []int, target int) []int {

	if len(nums) == 0 {
		return nil
	}

	firstIndex := 0
	secondIndex := 0

	for i, _ := range nums {

		for j, _ := range nums {
			if nums[i]+nums[j] == target && i != j {
				firstIndex = i
				secondIndex = j

				return []int{firstIndex, secondIndex}
			}
		}
	}

	return nil
}
