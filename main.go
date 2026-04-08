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
		if val%moduli[cutoff-1] == 1 {
			for i := length - 1; i > cutoff-1; i-- {
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
	if count < needed {
		return nil, errors.New("Not enough primes in slice")
	}
	return out, nil
}

func makePrimeSets(primes []int, moduli []int, needed int) ([][]int, error) {
	length := len(moduli)
	out := make([][]int, length)
	var err error
	for cutoff := length; cutoff > 0; cutoff-- {
		out[cutoff-1], err = primesInHybrids(primes, moduli, needed, cutoff)
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

// NOT DONE
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

func main() {
	fk.Factor(10)
	primes := pr.Sieve(10000)
	primeSets, _ := makePrimeSets(primes, []int{2, 5, 7}, 10)
	fmt.Println(primeSets)
	fmt.Println(logProdPrimeSets(primeSets, []int{1, 5, 3}, 1))

	//fmt.Println(partialExponents([][]int{{2, 1}, {3, 7}, {7, 3}}))

}
