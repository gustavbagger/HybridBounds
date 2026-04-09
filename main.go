package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	var mode string
	for {
		fmt.Println("Do you want a range of n-values,omega-values or single-value? (n,o,s)")
		reader := bufio.NewReader(os.Stdin)
		// ReadString will block until the delimiter is entered
		input, err := reader.ReadString('\n')
		if err != nil {
			continue
		}
		input = strings.TrimSpace(input)
		if len(input) != 1 || !(input == "n" || input == "o" || input == "s") {
			continue
		}
		mode = input
		break
	}
	switch mode {
	case "n":
	nRange:
		for {
			fmt.Printf("nMin,nMax,omega = ")
			reader := bufio.NewReader(os.Stdin)
			// ReadString will block until the delimiter is entered
			input, err := reader.ReadString('\n')
			if err != nil {
				continue
			}

			// remove the delimeter from the string
			input = strings.TrimSpace(input)
			args := strings.Split(input, ",")
			if len(args) != 3 {
				continue
			}
			nMin, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Println(err)
				continue
			}
			nMax, err := strconv.Atoi(args[1])
			if err != nil {
				fmt.Println(err)
				continue
			}
			omega, err := strconv.Atoi(args[2])
			if err != nil {
				fmt.Println(err)
				continue
			}
			for n := nMin; n <= nMax; n++ {
				bound, err := hybridBound(n, omega)
				if err != nil {
					fmt.Println(err)
					continue nRange
				}
				fmt.Printf("For n = %v,omega = %v: q^%v > %.2e\n", n, omega, n, math.Pow(math.Exp(bound), float64(n)))
			}
		}
	case "o":
	oRange:
		for {
			fmt.Printf("nMin,omegaMin,omegaMax = ")
			reader := bufio.NewReader(os.Stdin)
			// ReadString will block until the delimiter is entered
			input, err := reader.ReadString('\n')
			if err != nil {
				continue
			}

			// remove the delimeter from the string
			input = strings.TrimSpace(input)
			args := strings.Split(input, ",")
			if len(args) != 3 {
				continue
			}
			n, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Println(err)
				continue
			}
			omegaMin, err := strconv.Atoi(args[1])
			if err != nil {
				fmt.Println(err)
				continue
			}
			omegaMax, err := strconv.Atoi(args[2])
			if err != nil {
				fmt.Println(err)
				continue
			}
			for omega := omegaMin; omega <= omegaMax; omega++ {
				bound, err := hybridBound(n, omega)
				if err != nil {
					fmt.Println(err)
					continue oRange
				}
				fmt.Printf("For n = %v,omega = %v: q^%v > %.2e\n", n, omega, n, math.Pow(math.Exp(bound), float64(n)))
			}
		}
	case "s":
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
}
