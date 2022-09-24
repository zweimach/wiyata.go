package leetcode

func twoSum(nums []int, target int) []int {
	var result []int
loop:
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				result = []int{i, j}
				break loop
			}
		}
	}
	return result
}
