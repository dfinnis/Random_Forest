package forest

import "math/rand"

// shuffle randomizes the order of the data samples
func shuffle(data [][]float32) {
	rand.Shuffle(len(data), func(i, j int) {
		data[i], data[j] = data[j], data[i]
	})
}

// split splits data into training & test sets
func split(data [][]float32) (train_set [][]float32, test_set [][]float32) {
	var split float32 = 0.8
	var sample int
	for ; sample < int((float32(len(data)) * split)); sample++ {
		train_set = append(train_set, data[sample])
	}
	for ; sample < len(data); sample++ {
		test_set = append(test_set, data[sample])
	}
	return
}

// splitData shuffles data & creates training & test sets
func splitData(data [][]float32 /*, flags flags*/) (train_set, test_set [][]float32) {
	// shuffle(data) // seed!
	train_set, test_set = split(data)
	printSplit(len(train_set), len(test_set))
	return
}
