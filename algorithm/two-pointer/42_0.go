package main

func trap2(height []int) int {
	res := 0
	if len(height) <= 2 {
		return res
	}
	left, right := 0, len(height)-1
	for left < right {
		if height[left] > height[left+1] {
			break
		}
		left++
	}
	for left < right {
		if height[right] > height[right-1] {
			break
		}
		right--
	}
	// fmt.Println(left, right)
	for i := left + 1; i < right; i++ {
		if height[i] >= height[left] {
			left = i
			continue
		}
		r := i + 1
		for j := i + 1; j <= right; j++ {
			if height[j] >= height[left] {
				r = j
				break
			}
			if height[j] > height[r] {
				r = j
			}
		}

		if height[r] <= height[i] {
			continue
		}
		res += min(height[left], height[r]) - height[i]
		// fmt.Println(i, " ", left, " ", r, " ", min(height[left], height[r])-height[i])
	}

	return res
}
