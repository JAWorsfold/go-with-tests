package arrays

// playing with recursion
// func Sum(numbers []int) int {
// 	if len(numbers) == 1 {
// 		return numbers[0]
// 	}
// 	return numbers[0] + Sum(numbers[1:])
// }

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}