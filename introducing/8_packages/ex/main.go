func Max(xs []float64) float64 {
	var max float64
	for i, x := range xs {
		if i == 0 || x > max {
			max = x
		}
	}
}

func Min(xs []float64) float64 {
	var min float64
	for i, x := range xs {
		if i == 0 || x < min {
			min = x
		}
	}
	return min
}

