package twosum_1

func TwoSum(nums []int, target int) []int {
	tempMap := map[int]int{}

	for k, v := range nums {
		tempMap[v] = k
	}

	for k, v := range nums {
		delta := target - v
		idx, ok := tempMap[delta]
		if ok && idx != k {
			return  []int{k, idx}
		}
	}

	return []int{}
}
