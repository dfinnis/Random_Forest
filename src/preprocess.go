package forest

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

// readCsv reads data.csv into a 2d array of floats
func readCsv(filePath string) [][]float32 {
	f, err := os.Open(filePath)
	checkError("Unable to read input data file", err)
	defer f.Close()
	csvReader := csv.NewReader(f)
	// index := [32]string{"id", "diagnosis", "radius_mean", "texture_mean", "perimeter_mean", "area_mean", "smoothness_mean", "compactness_mean", "concavity_mean", "concave points_mean", "symmetry_mean", "fractal_dimension_mean", "radius_se", "texture_se", "perimeter_se", "area_se", "smoothness_se", "compactness_se", "concavity_se", "concave points_se", "symmetry_se", "fractal_dimension_se", "radius_worst", "texture_worst", "perimeter_worst", "area_worst", "smoothness_worst", "compactness_worst", "concavity_worst", "concave points_worst", "symmetry_worst", "fractal_dimension_worst"} // data.csv column titles

	var data [][]float32
	for {
		var drop bool
		dataStr, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		checkError("Unable to parse data file as CSV", err)
		var sample []float32
		for column, dataPoint := range dataStr {
			if column == 1 {
				if dataPoint == "M" {
					sample = append(sample, 1) // M = 1
				} else if dataPoint == "B" {
					sample = append(sample, 0) // B = 0
				} else {
					errorExit("Invalid data file format")
				}
			}
			if column > 1 {
				if dataPoint == "0" { // drop samples with empty data points
					drop = true
					break
				}
				float, err := strconv.ParseFloat(dataPoint, 64)
				checkError("Unable to parse data file as float", err)
				sample = append(sample, float32(float))
			}
		}
		if !drop {
			data = append(data, sample)
		}
	}
	return data
}

// checkData ensures data exists & each sample is 31 long
func checkData(data [][]float32, filePath string) {
	if len(data) == 0 {
		usageError("Data file empty: ", filePath)
	}
	length := len(data[0])
	for i, sample := range data {
		if len(sample) != length || len(sample) < 2 {
			fmt.Printf("%vERROR Data file invalid, len(data[%v]): %v%v\n", RED, i, len(sample), RESET)
			os.Exit(1)
		}
	}
}

// preprocess reads data.csv & standardizes data
func preprocess(dataPath string) [][]float32 {
	data := readCsv(dataPath)
	checkData(data, dataPath)
	fmt.Printf("Data loaded from: %v\n\n", dataPath)
	return data
}
