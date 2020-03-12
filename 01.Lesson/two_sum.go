package main

import "fmt"

// Returns indexes of items sum of which equals target
// Returns nil if there're no such indexes or input slice is empty
func twoSum(nums []int, target int) []int {
	if len(nums) == 0 {
		return nil
	}
	ans := []int{0, 0}
	m := make(map[int]int)
	m[0] = -1 // this is in case target is in the array as a single value
	for curIdx, item := range nums {
		otherIdx, ok := m[target-item]
		if ok {
			ans[0], ans[1] = otherIdx, curIdx
			if ans[0] == -1 {
				return ans[1:]
			}
			return ans
		}
		m[item] = curIdx
	}
	return nil
}

func main() {
	fmt.Println(twoSum([]int{1}, 1))
}
