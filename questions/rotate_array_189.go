package questions

// Rotate
// 189. 轮转数组
//给你一个数组，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。
//示例 1:
//
//输入: nums = [1,2,3,4,5,6,7], k = 3
//输出: [5,6,7,1,2,3,4]
//解释:
//向右轮转 1 步: [7,1,2,3,4,5,6]
//向右轮转 2 步: [6,7,1,2,3,4,5]
//向右轮转 3 步: [5,6,7,1,2,3,4]
//
//示例 2:
//
//输入：nums = [-1,-100,3,99], k = 2
//输出：[3,99,-1,-100]
//解释:
//向右轮转 1 步: [99,-1,-100,3]
//向右轮转 2 步: [3,99,-1,-100]
func Rotate(nums []int, k int) {
	l := len(nums)
	res := make([]int, l)
	for i := 0; i < l; i++ {
		res[(i+k)%(l)] = nums[i]
	}

	copy(nums, res)
}

func Rotate2(nums []int, k int) {
	for i := 0; i <= k; i++ {
		nums = append(nums, nums[i])
	}
	nums = nums[k+1:]
}

func Rotate3(nums []int, k int) {

}

//func Rotate3(nums []int, k int) {
//	k = k % len(nums)
//	l := len(nums)
//	next := nums[k]
//	for i := k; i >= 1; i-- {
//		tmp1 := nums[i-1]
//		nums[i] = nums[(k+i)%l]
//		nums[(i+k)%l] = next
//		next = tmp1
//	}
//}
