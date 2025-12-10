package sorted_squares

/*
977. 有序数组的平方
给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。

示例 1：
输入：nums = [-4,-1,0,3,10]
输出：[0,1,9,16,100]
解释：平方后，数组变为 [16,1,0,9,100]
排序后，数组变为 [0,1,9,16,100]

示例 2：
输入：nums = [-7,-3,2,3,11]
输出：[4,9,9,49,121]

提示：
1 <= nums.length <= 10^4
-10^4 <= nums[i] <= 10^4
nums 已按 非递减顺序 排序

进阶：
请你设计时间复杂度为 O(n) 的算法解决本问题
*/
func sortedSquares(nums []int) []int {
	if nums[len(nums)-1] <= 0 {
		i, j := 0, len(nums)-1
		for i <= j {
			nums[i], nums[j] = nums[j]*nums[j], nums[i]*nums[i]
			i++
			j--
		}

		return nums
	}
	mid := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] >= 0 && mid == -1 {
			mid = i
		}
		nums[i] = nums[i] * nums[i]
	}

	if mid == -1 {
		return nums
	}

	target := []int{}
	i, j := mid-1, mid
	for i >= 0 && j <= len(nums)-1 {
		if nums[i] <= nums[j] {
			target = append(target, nums[i])
			i--
		} else {
			target = append(target, nums[j])
			j++
		}
	}

	if i >= 0 {
		for ; i >= 0; i-- {
			target = append(target, nums[i])
		}
	}

	if j < len(nums) {
		for ; j < len(nums); j++ {
			target = append(target, nums[j])
		}
	}

	return target
}
