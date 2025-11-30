package main

import (
	"fmt"
	"pricecalculator/filemanager"
	"pricecalculator/prices"
)

func main() {
	// prices := []float64{10, 20, 30}
	taxrates := []float64{0, 0.7, 0.1, 0.15}

	// result := make(map[float64][]float64)

	// for _, taxRate := range taxrates {
	// 	taxIncludedPrices := make([]float64, len(prices))
	// 	for priceIndex, price := range prices {
	// 		taxIncludedPrices[priceIndex] = price * (1 + taxRate)
	// 	}
	// 	result[taxRate] = taxIncludedPrices
	// }

	// fmt.Println(result)

	for _, taxRate := range taxrates {
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		// cmdm := iomananger.New()
		priceJob := prices.New(fm, taxRate)
		err := priceJob.Process()
		if err != nil {
			fmt.Println("could not process job")
			fmt.Println(err)
		}
	}
	// fmt.Println(result)
}
