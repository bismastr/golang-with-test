package main

func Sum(arr []int) int {
	var sum int
	for _, num := range arr {
		sum += num
	}

	return sum
}

func SumAll(num ...[]int) []int {
	var result []int
	for _, num := range num {
		result = append(result, Sum(num))
	}

	return result
}

func SumAllTails(num ...[]int) []int {
	var result []int
	for _, n := range num {
		if len(n) == 0 {
			result = append(result, 0)
		} else {
			result = append(result, Sum(n[1:]))
		}
	}

	return result
}
