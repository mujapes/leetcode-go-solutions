import (
    //"math"
    //"fmt"
    "slices"
)
//used ai for reminder of fundamental theorem of arithmetic and that 
// k-factors are composed of the target's prime factors

func minDifference(n int, k int) []int {
    notPrimes := make([]bool, n + 1)
    primes := make([]int, 0)
    notPrimes[0] = true
    notPrimes[1] = true
    for i, v := range notPrimes {
        if !v {
            primes = append(primes, i)
            for j := i*i; j < len(notPrimes); j += i {
                notPrimes[j] = true
            }
        }
    }
    primeFactors := make([]int, 0)
    remainder := n
    for remainder > 1 {
        for i := 0; i < len(primes) ; i++{
            if remainder % primes[i] == 0 {
                primeFactors = append(primeFactors, primes[i])
                remainder /= primes[i]
                i = 0
            }
        }
    }
    kFactors := make([][][]int, 0)
    bestKFactor := make([]int, k)
    smallestDiff := 99999
    if len(primeFactors) == k {
        bestKFactor = append([]int{}, primeFactors...)
        smallestDiff = slices.Max(primeFactors) - slices.Min(primeFactors)
    } else if len(primeFactors) < k {
        padding := []int{}
        for _ = range k - len(primeFactors) {
            padding = append(padding, 1)
        }
        bestKFactor = append(primeFactors, padding...)
        smallestDiff = slices.Max(primeFactors) - 1
    }
    kFactors = append(kFactors, append([][]int{}, []int{}, primeFactors))
    for len(kFactors[len(kFactors) - 1][1]) > 1 {
        kFactors = append(kFactors, [][]int{{}})
        for possibleBestKFactor := range kFactors[len(kFactors) - 2] {
            for i, f := range kFactors[len(kFactors) - 2][possibleBestKFactor] {
                for j := i + 1; j < len(kFactors[len(kFactors) - 2][possibleBestKFactor]); j++ {
                    nextKFactor := []int{f*kFactors[len(kFactors) - 2][possibleBestKFactor][j]}
                    for iCopy, fCopy := range kFactors[len(kFactors) - 2][possibleBestKFactor] {
                        if iCopy != i && iCopy != j {
                            nextKFactor = append(nextKFactor, fCopy)
                        }
                    }
                    present := false
                    for _, fPresent := range kFactors[len(kFactors) - 1] {
                        if slices.Equal(fPresent, nextKFactor) {
                            present = true
                        }
                    }
                    if !present {
                        kFactors[len(kFactors) - 1] = append(kFactors[len(kFactors) - 1], nextKFactor)
                        nextKFactorDiff := 99999
                        if len(nextKFactor) < k {
                            nextKFactorDiff = slices.Max(nextKFactor) - 1
                            if nextKFactorDiff < smallestDiff {
                                smallestDiff = nextKFactorDiff
                                padding := []int{}
                                for _ = range k - len(nextKFactor) {
                                    padding = append(padding, 1)
                                }
                                bestKFactor = append(nextKFactor, padding...)
                            }
                        } else if len(nextKFactor) == k {
                            nextKFactorDiff = slices.Max(nextKFactor) - slices.Min(nextKFactor)
                            if nextKFactorDiff < smallestDiff {
                                smallestDiff = nextKFactorDiff
                                bestKFactor = nextKFactor
                            }
                        }
                    }
                }
            }
        }
    }
    return bestKFactor
}


//runtime: 829 ms, Beats 5.55%
//memory: 10.58 MB, Beats 6.94%

//Previously used a recursive approach which was even more memory intensive, 
// this solution still seems to be using too many resources
//Final program took 4-5 hours to produce over 2 days