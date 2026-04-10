package main

import "math"

func delta(vals []int) (float64, bool) {
	var out float64 = 1.0
	for _, val := range vals {
		out -= 1.0 / float64(val)
	}
	return out, out > 0
}

func logSieve(sievingPrimes, sievingPolys []int, omega, omegaCyclo int) (float64, bool) {
	del, valid := delta(append(sievingPrimes, sievingPolys...))
	if !valid {
		return 0, false
	}
	Del := 2.0 + float64(4*len(sievingPrimes)+len(sievingPrimes)-1)/del
	return math.Log(4) + math.Log(Del) + (float64(omegaCyclo+4*omega))*math.Log(2), true
}
