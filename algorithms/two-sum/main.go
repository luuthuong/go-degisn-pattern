package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	result := twoSum(nums, 5)
	fmt.Println(result)
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		if j, ok := m[target-v]; ok {
			return []int{j, i}
		}
		m[v] = i
	}
	return nil
}
