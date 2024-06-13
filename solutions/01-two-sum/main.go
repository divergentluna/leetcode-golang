package main

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for index, number := range nums {
		if value, ok := numMap[target-number]; ok {
			return []int{value, index}
		}
		numMap[number] = index
	}

	return nil
}
