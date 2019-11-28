package main

import "fmt"

func twoSum(nums []int, target int) []int {
	mp := make(map[int][]int)

	for index, num := range nums {
		arr, in := mp[num]

		if !in {
			arr = make([]int, 0)
		}
		mp[num] = append(arr, index)
	}

	for key, arr := range mp {

		halfIndex := arr[len(arr) -1]
		mp[key] = arr[:len(arr)-1]

		searchedArr, in := mp[target-key]

		if in && len(searchedArr) > 0 {
			return []int{halfIndex, searchedArr[len(searchedArr)-1]}
		}

		mp[key] = append(arr, halfIndex)

	}

	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 11}, 22))
}
