package main

func main() {

}

func searchInsert(nums []int, target int) int {
	p := 0
	if target <= nums[0] {
		return p
	}
	if len(nums) == 1 && target > nums[0] {
		return 1
	}
	for i := 1; i < len(nums); i++ {
		if target >= nums[i] || (target < nums[i] && target > nums[i-1]) {
			p = i
		}
		if i == len(nums)-1 && target > nums[i] {
			p = i + 1
		}
	}
	return p
}
