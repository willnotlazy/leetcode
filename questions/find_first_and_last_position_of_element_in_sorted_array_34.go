package questions

func SearchRange(nums []int, target int) []int {
	left_l, left_r := 0, 0
	right_l, right_r := len(nums)-1, len(nums)-1

	for left_l <= right_l || left_r <= right_r {
		// 确定左边界
		if left_l <= right_l {
			mid_l := left_l + (right_l-left_l)/2
			if nums[mid_l] == target {
				right_l = mid_l - 1
			} else if nums[mid_l] > target {
				right_l = mid_l - 1
			} else if nums[mid_l] < target {
				left_l = mid_l + 1
			}
		}

		// 确定右边界
		if left_r <= right_r {
			mid_r := left_r + (right_r-left_r)/2
			if nums[mid_r] == target {
				left_r = mid_r + 1
			} else if nums[mid_r] > target {
				right_l = mid_r - 1
			} else if nums[mid_r] < target {
				left_r = mid_r + 1
			}
		}
	}
	left_ans, right_ans := -1, -1
	if left_l != len(nums) && nums[left_l] == target {
		left_ans = left_l
	}
	if left_r-1 != -1 && nums[left_r-1] == target {
		right_ans = left_r - 1
	}

	return []int{left_ans, right_ans}
}
