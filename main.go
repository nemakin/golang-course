package main

import "fmt"

func twoSum(nums []int, target int) []int {
	idx := make(map[int]int)

	for i, x := range nums {
		if j, ok := idx[target-x]; ok {
			return []int{j, i}
		}
		idx[x] = i
	}

	return nil
}

func main() {
	fmt.Println(twoSum([]int{2, -1, 4, 8, 10}, 7))
}
