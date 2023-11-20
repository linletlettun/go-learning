package main

import "fmt"

func main() {
	counts := make([][]int, 3)

	counts[0] = []int{56, 65, 70, 59, 87, 28, 65, 39, 75, 60}
	counts[1] = []int{52, 90, 100, 30, 60}
	counts[2] = []int{90, 60, 80, 40, 90, 100, 30, 20, 0, 40, 50, 60, 97, 50, 44}

	for i := 0; i < len(counts); i++ {
		var full, flying, pass, fail int

		for j := 1; j < len(counts[i]); j++ {
			if counts[i][j] == 100 {
				full++
			} else if counts[i][j] > 80 {
				flying++
			} else if counts[i][j] > 40 {
				pass++
			} else {
				fail++
			}
		}
		fmt.Printf("Class %d: Full mark: %d, Flying color: %d, Pass: %d, Fail: %d\n", i, full, flying, pass, fail)
	}

}
