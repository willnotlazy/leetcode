package main

func main() {
	removeElement(
		[]int{0, 1, 2, 2, 3, 0, 4, 2},
		2,
	)
}

func removeElement(nums []int, val int) int {
	slow := 0
	fast := 0
	for ; fast < len(nums) && slow < len(nums); fast++ {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
