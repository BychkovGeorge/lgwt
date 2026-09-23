package arrays

import "fmt"

func Sum(slice []int) (result int) {
	for _, v := range slice {
		result += v
	}
	return
}

func SumAll(slices ...[]int) []int {
	var result []int
	for _, v := range slices {
		result = append(result, Sum(v))
	}
	return result
}

func SumAllTails(slices ...[]int) []int {
	var result []int
	for _, v := range slices {
		result = append(result, Sum(v[1:]))
	}
	fmt.Println("res", result)
	return result
}
