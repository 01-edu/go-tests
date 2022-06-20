package solutions

func FindMissingNumber(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	for i := 0; i < len(nums); i++ {
		if i+1 < len(nums) && nums[i]+1 != nums[i+1] {
			return nums[i] + 1
		}
	}
	return -1
}
