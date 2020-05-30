package kata

// Step 1: Split the array in two
func split(input []int) [2][]int{
	l := len(input)
	return [2][]int{input[0:l/2], input[l/2:l]}
}

// Step 2: Put the arrays on top of each other
// Step 3: Add them together
func sum(a1 []int, a2 []int) []int {
	arraySum := make([]int, len(a2))
	// case 1: equal length
	if len(a1) == len(a2) {
		for i, v := range a1 {
			arraySum[i] = v + a2[i]
		}
	} else {
		// case 2: not equal length, second array is 1 longer
		for i, v := range a2 {
			if i == 0 {
				// special case 
				arraySum[i] = v
			} else {
				arraySum[i] = v + a1[i-1]
			}
		}
	}

	return arraySum
}

func SplitAndAdd(numbers []int, n int) []int {
	var result []int
	result = numbers
	for n > 0 {
		splitArray := split(result)
		result = sum(splitArray[0], splitArray[1])
		n--
	}
	return result
}
