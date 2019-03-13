package v1

func Sum(numbers []int) (sum int) {
	for _, value := range numbers{
		sum += value
	}
	return
}

func SumAll(numbersToSum ...[]int) (sums []int) {
	length := len(numbersToSum)
	sums = make([]int, length)
	for index,numbers := range numbersToSum {
		sums[index] = Sum(numbers)
	}
	return
}

func SumAllTail(numbersToSum ...[]int) (sums []int) {
	for _,numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return
}


