package kstats

import "math"

// Mean calculates a point estimate of the mean of an array of observations.
func Mean(x []float64) float64 {
	var n, sm, mn float64
	n = float64(len(x))
	for i := 0; i < len(x); i++ {
		sm += x[i]
	}
	mn = sm / n
	return mn
}

// WeightedMean calculates a weighted mean of an array of observations given an array of weights.
func WeightedMean(x, w []float64) float64 {
	var xw_sm, w_sm, wmn float64
	for i := 0; i < len(x); i++ {
		xw_sm += x[i] * w[i]
		w_sm += w[i]
	}
	wmn = xw_sm / w_sm
	return wmn
}

// GeometricMean calculates a geometric mean of an array of observations (used in geometic growth contexts).
func GeometricMean(x []float64) float64 {
	var n, prd, gmn float64
	n = float64(len(x))
	prd = 1.0
	for i := 0; i < len(x); i++ {
		prd *= x[i]
	}
	gmn = math.Pow(prd, 1.0/n)
	return gmn
}

// Variance calculates the variance of an array of observations.
func Variance(x []float64, sample bool) float64 {
	var n, mn, se, vr float64
	if sample {
		n = float64(len(x) - 1)
	} else {
		n = float64(len(x))
	}
	mn = Mean(x)
	for i := 0; i < len(x); i++ {
		diff := x[i] - mn
		se += math.Pow(diff, 2)
	}
	vr = se / n
	return vr
}

// Covariance calculates the covariance of two arrays of observations.
func Covariance(x, y []float64, sample bool) float64 {
	var n, x_mn, y_mn, xy_er, cv float64
	if sample {
		n = float64(len(x) - 1)
	} else {
		n = float64(len(x))
	}
	x_mn = Mean(x)
	y_mn = Mean(y)
	for i := 0; i < len(x); i++ {
		x_diff := x[i] - x_mn
		y_diff := y[i] - y_mn
		xy_er += x_diff * y_diff
	}
	cv = xy_er / n
	return cv
}

// CorrelationCoefficient calculates the pearson Product Moment Corrrelation Coefficient of two arrays of observations.
func CorrelationCoefficient(x, y []float64, sample bool) float64 {
	var x_sd, y_sd, xy_cv, xy_cr float64
	x_sd = math.Sqrt(Variance(x, sample))
	y_sd = math.Sqrt(Variance(y, sample))
	xy_cv = Covariance(x, y, sample)
	xy_cr = xy_cv / (x_sd * y_sd)
	return xy_cr
}

// ZScore calculates the z-score for a particular observation.
func ZScore(x float64, mn float64, sd float64) float64 {
	return (x - mn) / sd
}
