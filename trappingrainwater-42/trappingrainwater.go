package trappingrainwater_42

func trap(height []int) int {
	var left, right, leftMax, rightMax, total = 0, len(height) - 1, 0, 0, 0

	for ; left < right; {
		if height[left] < height[right] {
			if height[left] > leftMax {
				leftMax = height[left]
			} else {
				total += leftMax - height[left]
			}
			left++
		} else {
			if height[right] > rightMax {
				rightMax = height[right]
			} else {
				total += rightMax - height[right]
			}
			right--
		}
	}

	return total
}
