package questions

// Search
// 704. 二分查找
// description: 给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。
func Search(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	mid := start + (end-start)>>1
	for start <= end {
		if nums[mid] == target {
			return mid
		} else {
			if nums[mid] > target {
				end = mid - 1
			} else {
				start = mid + 1
			}
			mid = start + (end-start)>>1
		}
	}

	return -1
}
