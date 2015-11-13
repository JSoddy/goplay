package math

func Average(xs []float64) float64 {
	total := float64(0)
	if len(xs) == 0 {
		return float64(0)
	}
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}

func Min(numList []float64) float64 {
	if len(numList) == 0 {
		return 0
	}
	curMin := numList[0]
	for _, x := range numList {
		if x < curMin {
			curMin = x
		}
	}
	return curMin
}

func Max(numList []float64) float64 {
	if len(numList) == 0 {
		return 0
	}
	curMax := numList[0]
	for _, x := range numList {
		if x > curMax {
			curMax = x
		}
	}
	return curMax
}
