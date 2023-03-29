package main

// heap algorithm
func heapPermutate(mystr []string, n int, permutations *[][]string) {
	if n == 1 {
		*permutations = append(*permutations, append([]string{}, mystr...))
		return
	}
	for i := 0; i < n; i++ { //calculate the permutation using heap
		heapPermutate(mystr, n-1, permutations)
		if n%2 == 0 {
			mystr[i], mystr[n-1] = mystr[n-1], mystr[i]
		} else {
			mystr[0], mystr[n-1] = mystr[n-1], mystr[0]
		}
	}
}

// backtrack
func backtrackPermutate(nums []string) [][]string {
	var result [][]string
	backtrack(nums, &result, 0)
	return result
}

func backtrack(nums []string, result *[][]string, index int) {
	if index == len(nums) {
		temp := make([]string, len(nums))
		copy(temp, nums)
		*result = append(*result, temp)
		return
	}
	for i := index; i < len(nums); i++ {
		nums[index], nums[i] = nums[i], nums[index]
		backtrack(nums, result, index+1)
		nums[index], nums[i] = nums[i], nums[index]
	}
}

// 繰り返し法
func repetitionPermute(nums []string) [][]string {
	var result [][]string
	result = append(result, []string{})
	for _, num := range nums {
		size := len(result)
		for j := 0; j < size; j++ {
			temp := result[0]
			result = result[1:]
			for k := 0; k <= len(temp); k++ {
				newPerm := insert(temp, num, k)
				result = append(result, newPerm)
			}
		}
	}
	return result
}

func insert(nums []string, num string, index int) []string {
	var result []string
	result = append(result, nums[:index]...)
	result = append(result, num)
	result = append(result, nums[index:]...)
	return result
}

// 部分集合を用いたアルゴリズム
// TODO
