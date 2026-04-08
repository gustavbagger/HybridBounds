package main

import (
	"errors"
	"fmt"
	"math"

	pr "github.com/fxtlabs/primes"
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

func primesInHybrids(primes []int, moduli []int, needed int, cutoff int) ([]int, error) {
	count := 0
	var out []int
	length := len(moduli)
primeloop:
	for _, val := range primes {
		if cutoff == 0 || val%moduli[cutoff-1] == 1 {
			for i := length - 1; i >= cutoff; i-- {
				if val%moduli[i] == 1 {
					continue primeloop
				}
			}
			out = append(out, val)
			count++
		}
		if count >= needed {
			break
		}
	}
	if count < needed && (cutoff != 0 || moduli[0] != 2) {
		return nil, errors.New("Not enough primes in slice")
	}
	return out, nil
}

func makePrimeSets(primes []int, moduli []int, needed int) ([][]int, error) {
	length := len(moduli)
	out := make([][]int, length+1)
	var err error
	for cutoff := length; cutoff >= 0; cutoff-- {
		out[cutoff], err = primesInHybrids(primes, moduli, needed, cutoff)
		if err != nil {
			return nil, err
		}
	}
	return out, nil
}

func partialExponents(n [][]int) []float64 {
	out := make([]float64, len(n)+1)
	out[0] = 1.0
	for i := 0; i < len(n); i++ {
		prime := float64(n[i][0])
		exp := float64(n[i][1])
		out[i+1] = out[i] / math.Pow(prime, exp)
	}
	return out
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

func findLargestBound(primeSets [][]int, spread []int, exps []float64) (float64, error) {
	var bound float64
	var current, try float64
	if len(spread) != len(exps) {
		return 0, errors.New("mismatched set lengths, exps/spread")
	}
	if len(spread) != len(primeSets) {
		return 0, errors.New("mismatched set lengths, primeSets/spread")
	}
	for i := range len(spread) {
		bound += logProd(primeSets[i][:spread[i]])
		try = exps[i] * bound
		if try > current {
			current = try
		}
	}
	return current, nil
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

func findWorstBound(primeSets [][]int, exps []float64, omega int) (float64, error) {

	partitions := additivePartitions(len(exps), omega)
	current, err := findLargestBound(primeSets, partitions[0], exps)
	if err != nil {
		return 0, err
	}
	for _, spread := range partitions {
		new, err := findLargestBound(primeSets, spread, exps)
		if err != nil {
			return 0, err
		}
		if new < current {
			current = new
		}
	}
	return current, nil
}

func main() {
	fk.Factor(10)

	primes := pr.Sieve(10000)
	omega := 40
	n := [][]int{{5, 1}}
	primeSets, _ := makePrimeSets(primes, []int{5}, omega)
	exps := partialExponents(n)
	fmt.Println(primeSets, "   ", exps)
	logBound, _ := findWorstBound(primeSets, exps, omega)
	fmt.Println(math.Pow(math.E, logBound))
}
