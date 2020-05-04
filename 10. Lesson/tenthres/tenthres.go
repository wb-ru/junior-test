package tenthres

func intersection(nums1 []int, nums2 []int) []int {
	result := []int{}
	mp := make(map[int]int)
	for _, val := range nums1 {
		mp[val]++
	}
	for _, val := range nums2 {
		if mp[val] > 0 {
			result = append(result, val)
			mp[val]--
		}
	}
	return result
}
