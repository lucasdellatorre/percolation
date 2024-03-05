package util

import "math/rand"

func GenerateUniqueRandomNumbers(n int) []int {
	randomNumbers := make([]int, n)
	usedNumbers := make(map[int]bool)

	for i := 0; i < n; i++ {
		var num int
		for {
			num = rand.Intn(n)
			if !usedNumbers[num] {
				break
			}
		}
		randomNumbers[i] = num
		usedNumbers[num] = true
	}

	return randomNumbers
}
