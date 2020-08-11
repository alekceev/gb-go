package statistic

func Average(xs []float64) float64 {
	var total float64

	if len(xs) == 0 {
		return total
	}

	return Sum(xs) / float64(len(xs))
}

func Sum(xs []float64) float64 {
	var sum float64

	for _, x := range xs {
		sum += x
	}
	return sum
}
