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

// WeightedMean calculates a weighted mean of an array of observations given an array of weights.
func WeightedMean(obs []float64, wgts []float64) float64 {
	var obs_sum, wgts_sum, mean float64
	for i := 0; i < len(obs); i++ {
		obs_sum += obs[i] * wgts[i]
		wgts_sum += wgts[i]
	}
	mean = obs_sum / wgts_sum
	return mean
}

// GeometricMean calculates a geometric mean of an array of observations (used in geometic growth contexts).
func GeometricMean(obs []float64) float64 {
	var prod, mean float64
	prod = 1.0
	for i := 0; i < len(obs); i++ {
		prod *= obs[i]
	}
	mean = math.Pow(prod, 1.0/float64(len(obs)))
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

// ZScore calculates the z-score for a particular observation.
func ZScore(obs float64, mean float64, sd float64) float64 {
	return (obs - mean) / sd
}
