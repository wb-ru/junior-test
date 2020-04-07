package twosum

func twoSum(nums []int, target int) []int {
	if len(nums) < 2 {
		return nil
	}
	getIndexByValue := make(map[int]int)
	for i, value := range nums {
		getIndexByValue[value] = i
	}
	for i, value := range nums {
		need := target - value
		currentIndex, ok := getIndexByValue[need]
		if ok && currentIndex != i {
			return []int{i, currentIndex}
		}
	}
	return []int{-1, -1}
}
