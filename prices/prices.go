package prices

import (
	"fmt"
	"pricecalculator/conversion"
	"pricecalculator/filemanager"
)

type TaxIncludedPriceJob struct {
	IOManager         filemanager.FileManager `json:"_"` // ignore from json file
	TaxRate           float64                 `json:"taxrate"`
	InputPrices       []float64               `json:"input_prices"`
	TaxIncludedPrices map[string]string       `json:"tax_Included_prices"`
}

// Read text form the file
func (job *TaxIncludedPriceJob) LoadData() error {

	// moved to filemanager.go
	/*
		file, err := os.Open("prices.txt")
		if err != nil {
			fmt.Println("Could not open file")
			fmt.Print(err)
			return
		}

		scanner := bufio.NewScanner(file) // read line by line from file
		var lines []string
		for scanner.Scan() {
			// step by step read by scan form file
			lines = append(lines, scanner.Text())
		}
		err = scanner.Err()
		if err != nil {
			fmt.Println("reading the file content failed")
			file.Close()
			fmt.Print(err)
			return
		}
	*/
	/*
		using make keyword, we can create placeholder for float64 type list/Array ,
		with desired length using len()
	*/

	// lines, err := filemanager.ReadLines("prices.txt")
	lines, err := job.IOManager.ReadLines()
	if err != nil {
		fmt.Println(err)
		return err
	}

	prices, err := conversion.StringToFloats(lines)

	if err != nil {
		fmt.Println("Converting price to float failed")
		fmt.Println(err)
		return err
	}

	/*
		// Moved to conversion package file - conversion.go
		// convert string to float64 value
		prices := make([]float64, len(lines))
		for lineIndex, line := range lines {
			floatPrice, err := strconv.ParseFloat(line, 64)
			if err != nil {
				fmt.Println("Converting price to float failed")
				fmt.Println(err)
				file.Close()
				return
			}
			prices[lineIndex] = floatPrice
		}
	*/

	job.InputPrices = prices
	return nil
}

func (job *TaxIncludedPriceJob) Process() error {
	err := job.LoadData()
	if err != nil {
		return err
	}
	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	job.TaxIncludedPrices = result

	// filemanager.WritJSON(fmt.Sprintf("result_%.0f.json", job.TaxRate*100), job)
	return job.IOManager.WriteResult(job)

	// fmt.Println(result)
}

func New(fm filemanager.FileManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager:   fm,
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
