package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	pr "github.com/fxtlabs/primes"
)

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

func findWorstBound(primeSets [][]int, exps []float64, omega int) (float64, error) {
	var partitions [][]int
	//Handling n even case
	if len(primeSets[0]) == 1 && primeSets[0][0] == 2 {
		//Forces exactly one divisor present in q-1 (there could be more but these could just as well be in q+1 = \Phi_2)
		partitions = prependManyTimes(1, additivePartitions(len(exps)-1, omega-1))
	} else {
		partitions = additivePartitions(len(exps), omega)
	}
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

// Power of 2 case, assuming n= 2^exp
func HybridBoundTwoPow(primes []int, exp, omega int) float64 {
	oneOnN := 1.0 / float64(int(1)<<exp)
	upper := oneOnN * logProd(primes[1:omega+1])
	lower := oneOnN * (logProd(primes[0:omega]) + float64(1+exp)*math.Log(2))
	if upper > lower {
		return lower
	} else {
		return upper
	}
}

func hybridBound(n, omega int) (float64, error) {
	nSlice, nPrimes := factorIntoSlice(n)
	primes := pr.Sieve(10000)

	if nPrimes[0] == 2 && len(nPrimes) == 1 {
		return HybridBoundTwoPow(primes, nSlice[0][1], omega), nil
	}
	primeSets, err := makePrimeSets(primes, nPrimes, omega)
	if err != nil {
		return 0, err
	}
	exps := partialExponents(nSlice)
	return findWorstBound(primeSets, exps, omega)
}

func main() {
	for {
		fmt.Printf("n,omega = ")
		reader := bufio.NewReader(os.Stdin)
		// ReadString will block until the delimiter is entered
		input, err := reader.ReadString('\n')
		if err != nil {
			continue
		}

		// remove the delimeter from the string
		input = strings.TrimSpace(input)
		args := strings.Split(input, ",")
		if len(args) != 2 {
			continue
		}
		n, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println(err)
			continue
		}
		omega, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println(err)
			continue
		}
		bound, err := hybridBound(n, omega)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("For n = %v,omega = %v: q^%v > %.2e\n", n, omega, n, math.Pow(math.Exp(bound), float64(n)))
	}
}
