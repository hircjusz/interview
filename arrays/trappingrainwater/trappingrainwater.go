package trappingrainwater

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

// {3, 0, 0, 2, 0, 4}
func traprainwater(arr []int) int {

	nivelLeft := make([]int, len(arr))

	maxL := arr[0]

	for i := 0; i < len(arr); i++ {
		if maxL < arr[i] {
			maxL = arr[i]
		}
		nivelLeft[i] = maxL
	}

	nivelRight := make([]int, len(arr))

	maxR := arr[len(arr)-1]

	for i := len(arr) - 1; i >= 0; i-- {
		if maxR < arr[i] {
			maxR = arr[i]
		}
		nivelRight[i] = maxR
	}

	var val int
	for i := 0; i < len(arr); i++ {
		val += min(nivelLeft[i], nivelRight[i]) - arr[i]
	}

	return val
}
