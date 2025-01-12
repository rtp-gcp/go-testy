package main

import (
	"fmt"
	"github.com/montanaflynn/stats"
)

func main() {
	data := []float64{1.0, 2.1, 3.2, 4.823, 4.1, 5.8}

	median, _ := stats.Median(data)
	fmt.Println(median) // 3.65

	roundedMedian, _ := stats.Round(median, 0)
	fmt.Println(roundedMedian) // 4
}
