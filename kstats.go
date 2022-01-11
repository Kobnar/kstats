package kstats

// SampleMean calculates a point estimate of the mean of an array of observations.
func SampleMean(obs []float64) float64 {
	var sum, mean float64
	for i := 0; i < len(obs); i++ {
		sum += obs[i]
	}
	mean = sum / float64(len(obs))
	return mean
}
