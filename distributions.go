package kstats

import "math"

type BinomialPMF struct {
	n int
	p float64
}

func (dist BinomialPMF) f(k int) float64 {
	var cmb int
	var n_, k_, p float64
	n_ = float64(dist.n)
	k_ = float64(k)
	p = float64(cmb)
	cmb = Combinations(dist.n, k)
	p *= math.Pow(dist.p, k_)
	p *= math.Pow(1-dist.p, n_-k_)
	return p
}

func (dist BinomialPMF) ExpectedValue() float64 {
	return float64(dist.n) * dist.p
}

func (dist BinomialPMF) Variance() float64 {
	return float64(dist.n) * dist.p * (1 - dist.p)
}

type PoissonPMF struct {
	lmbd float64
}

func (dist PoissonPMF) f(k int) float64 {
	var k_bang int
	var k_, p float64
	k_ = float64(k)
	k_bang = 1
	for i := 0; i < k; k++ {
		k_bang *= k - i
	}
	p = math.Pow(dist.lmbd, k_)
	p *= math.Pow(math.E, -dist.lmbd)
	p /= float64(k_bang)
	return p
}

func (dist PoissonPMF) ExpectedValue() float64 {
	return dist.lmbd
}

func (dist PoissonPMF) Variance() float64 {
	return dist.lmbd
}

type HypergeometricPMF struct {
	N, K, n int
}

func (dist HypergeometricPMF) f(n, k int) float64 {
	var num, den int
	var pb float64
	num = Combinations(dist.K, k)
	num *= Combinations(dist.N-dist.K, n-k)
	den = Combinations(dist.N, n)
	pb = float64(num) / float64(den)
	return pb
}

func (dist HypergeometricPMF) ExpectedValue() float64 {
	var n_, K_, N_ float64
	n_ = float64(dist.n)
	K_ = float64(dist.K)
	N_ = float64(dist.N)
	return n_ * K_ / N_
}

func (dist HypergeometricPMF) Variance() float64 {
	var n_, K_, N_, NK, Nn, N1, num, den float64
	n_ = float64(dist.n)
	K_ = float64(dist.K)
	N_ = float64(dist.N)
	NK = float64(dist.N - dist.K)
	Nn = float64(dist.N - dist.n)
	N1 = float64(dist.N - 1)
	num = n_ * K_ * NK * Nn
	den = N_ * N_ * N1
	return num / den
}

type UniformPDF struct {
	a, b float64
}

func (dist UniformPDF) f(x float64) float64 {
	if dist.a <= x || x <= dist.b {
		return 1 / (dist.b - dist.a)
	} else {
		return 0.0
	}
}

func (dist UniformPDF) ExpectedValue() float64 {
	return (dist.a + dist.b) / 2.0
}

func (dist UniformPDF) Variance() float64 {
	return math.Pow(dist.b-dist.a, 2.0) / 12.0
}

type ExponentialPDF struct {
	lmbd float64
}

func (dist ExponentialPDF) f(x float64) float64 {
	return math.Pow(math.E, -x/dist.lmbd) / dist.lmbd
}

func (dist ExponentialPDF) F(x float64) float64 {
	return 1 - math.Pow(math.E, -x/dist.lmbd)
}

func (dist ExponentialPDF) ExpectedValue() float64 {
	return 1 / dist.lmbd
}

func (dist ExponentialPDF) Variance() float64 {
	return 1 / math.Pow(dist.lmbd, 2.0)
}
