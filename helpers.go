package main

import (
	"errors"
	"math"

	fk "github.com/johnkerl/goffl/pkg/intfactor"
)

type Number interface {
	float64 | int | uint64 | uint | uint32 | float32
}

func Min[T Number](vals []T) T {
	out := vals[0]
	for _, val := range vals {
		if val < out {
			out = val
		}
	}
	return out
}

func Max[T Number](vals []T) T {
	out := vals[0]
	for _, val := range vals {
		if val > out {
			out = val
		}
	}
	return out
}

func logProd[T Number](vals []T) float64 {
	var out float64
	for _, val := range vals {
		out += math.Log(float64(val))
	}
	return out
}

func primesInArith(primes []int, modulus int, needed int) ([]int, error) {
	count := 0
	var out []int
	for _, val := range primes {
		if val%modulus == 1 {
			out = append(out, val)
			count++
		}
		if count >= needed {
			break
		}
	}
	if count < needed {
		return nil, errors.New("Not enough primes in slice")
	}
	return out, nil
}

func logProdPrimeSets(primeSets [][]int, spread []int, exp int) (float64, error) {
	var out float64
	if len(primeSets) != len(spread) {
		return 0, errors.New("mismatched set lengths, primeSets/spread")
	}
	for i := range len(spread) {
		out += logProd(primeSets[i][:spread[i]])
	}

	return float64(exp) * out, nil
}
func prependManyTimes(val int, slices [][]int) [][]int {
	var out [][]int
	valSlice := []int{val}
	for _, slice := range slices {
		out = append(out, append(valSlice, slice...))
	}
	return out
}
func additivePartitions(piles, total int) [][]int {
	var out [][]int
	if piles == 1 {
		return [][]int{{total}}
	}
	for i := 1; i <= total+1-piles; i++ {
		out = append(out, prependManyTimes(i, additivePartitions(piles-1, total-i))...)
	}
	return out
}

func factorIntoSlice(n int) ([][]int, []int) {
	nFactors := fk.Factor(int64(n))
	var nPrimes []int
	var nSlice [][]int
	for i := 0; i < nFactors.NumDistinctFactors(); i++ {
		p, m := nFactors.Get(i)
		nSlice = append(nSlice, []int{int(p), m})
		nPrimes = append(nPrimes, int(p))
	}
	return nSlice, nPrimes
}
