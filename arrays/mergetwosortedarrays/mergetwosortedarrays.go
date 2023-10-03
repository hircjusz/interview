package mergetwosortedarrays

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func mergetwosortedarrays(arr1 []int, arr2 []int) []int {

	resultArr := []int{}
	//lowlen := min(len(arr1), len(arr2))

	for len(arr1) > 0 && len(arr2) > 0 {

		mina, minb := arr1[0], arr2[0]

		if mina < minb {
			resultArr = append(resultArr, mina)
			arr1 = arr1[1:]
		} else {
			resultArr = append(resultArr, minb)
			arr2 = arr2[1:]
		}
	}
	if len(arr1) > 0 {
		resultArr = append(resultArr, arr1...)
	}
	if len(arr2) > 0 {
		resultArr = append(resultArr, arr2...)
	}

	return resultArr
}
