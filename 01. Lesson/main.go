package twoSums

func twoSum(nums []int, target int) []int {
	if len(nums) == 0 {
		return nil
	}
	ret := []int{0, 0}
	for i, num := range nums {
		for j, num2 := range nums {
			if (i!=j)&&(num+num2==target) {
				ret[0], ret[1] = i, j
				return ret
			}
		}
	}
	return nil

}

