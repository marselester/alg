// Command 3sum-doubling produces experimental data for 3-sum algorithm.
// It generates a sequence of random inputs arrays, doubling the array size at each step,
// and prints the running times of 3-sum for each input size.
// The program also calculates the ratio of each running time with the previous.
// The ratio for TripleCountNaive is about 8 and we can predict the running times
// for next experiments by multiplying the last duration by 8.
package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	"github.com/marselester/alg/zerosum"
)

func main() {
	var size, maxSize int
	flag.IntVar(&size, "size", 125, "min input size")
	flag.IntVar(&maxSize, "max-size", 1e6, "max input size")
	wait := flag.Duration("wait", time.Minute, "wait for the experiment to finish")
	flag.Parse()

	var elapsed, prevElapsed, ratio float64
	for {
		if size >= maxSize || elapsed >= wait.Seconds() {
			break
		}

		nums := randNums(size)
		begun := time.Now()
		zerosum.TripleCountNaive(nums)
		elapsed = time.Since(begun).Seconds()

		if prevElapsed == 0 {
			ratio = 0
		} else {
			ratio = elapsed / prevElapsed
		}
		fmt.Printf("%d %0.2f %0.2f\n", size, elapsed, ratio)

		prevElapsed = elapsed
		size *= 2
	}

	// Predict the running times by observed ratio.
	for {
		if size >= maxSize {
			break
		}
		size *= 2

		elapsed = ratio * prevElapsed
		fmt.Printf("%d %0.2f %0.2f\n", size, elapsed, ratio)
		prevElapsed = elapsed
	}
}

func randNums(n int) []int {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = r.Int()
		// Flip a sign of an odd number.
		if nums[i]&1 == 1 {
			nums[i] *= -1
		}
	}
	return nums
}
