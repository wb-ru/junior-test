package mergeArrays

func Merge(arr1 []int, arr2 []int) []int {
	if len(arr1) > len(arr2) {
		arr1, arr2 = arr2, arr1
	}
	res := []int{}
	m := make(map[int]int)
	for _, item := range arr1 {
		m[item] += 1
	}
	for _, item := range arr2 {
		freq, ok := m[item]
		if ok && freq > 0 {
			res = append(res, item)
			freq--
			m[item] = freq
		}
	}

	return res
}
