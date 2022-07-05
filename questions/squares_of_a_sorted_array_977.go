package questions

// SortedSquares
// 977. 有序数组的平方
//给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。
//
//
//
//示例 1：
//
//输入：nums = [-4,-1,0,3,10]
//输出：[0,1,9,16,100]
//解释：平方后，数组变为 [16,1,0,9,100]
//排序后，数组变为 [0,1,9,16,100]
func SortedSquares(nums []int) []int {
	left, right := 0, 0
	if nums[len(nums)-1] <= 0 {
		left = len(nums) - 1
		right = len(nums)
	} else if nums[0] >= 0 {
		left = -1
		right = 0
	} else {
		for i := 0; i < len(nums); i++ {
			if nums[i] >= 0 {
				left = i - 1
				right = i
				break
			}
		}
	}

	tmp := make([]int, len(nums))
	i := 0
	for left >= 0 || right < len(nums) {
		if right < len(nums) && left >= 0 {
			leftDouble := nums[left] * nums[left]
			rightDouble := nums[right] * nums[right]
			if leftDouble < rightDouble {
				tmp[i] = leftDouble
				left--
			} else {
				tmp[i] = rightDouble
				right++
			}
		} else if left >= 0 {
			tmp[i] = nums[left] * nums[left]
			left--
		} else {
			tmp[i] = nums[right] * nums[right]
			right++
		}
		i++
	}

	return tmp
}
