package kstats

import "math"

// SampleMean calculates a point estimate of the mean of an array of observations.
func SampleMean(obs []float64) float64 {
	var sum, mean float64
	for i := 0; i < len(obs); i++ {
		sum += obs[i]
	}
	mean = sum / float64(len(obs))
	return mean
}

// SampleVariance calculates a point estimate of the variance of an array of observations.
func SampleVariance(obs []float64) float64 {
	var sample_mean, squared_error, sample_variance float64
	sample_mean = SampleMean(obs)
	for i := 0; i < len(obs); i++ {
		diff := obs[i] - sample_mean
		squared_error += math.Pow(diff, 2)
	}
	sample_variance = squared_error / float64(len(obs)-1)
	return sample_variance
}
