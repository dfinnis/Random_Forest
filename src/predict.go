package forest

import "fmt"

// truthTally counts true & false, positives & negatives
func truthTally(predictions, truth []bool) (tp, fn, fp, tn uint) {
	// var tp uint // True Positive		// Predicted True & Is True
	// var fn uint // False Negative	// Predicted False & Is True
	// var fp uint // False Positive	// Predicted True & Is False
	// var tn uint // True Negative		// Predicted False & Is False

	for i := 0; i < len(predictions); i++ {
		// fmt.Printf("%-3v prediction: %-5v, truth: %v\n", i, predictions[i], truth[i]) ///////////
		if truth[i] { // Is True
			// fmt.Printf("Is True\n") ///////////
			if predictions[i] { // Predicted True
				tp += 1
			} else { // Predicted False
				fn += 1
			}
		} else { // Is False
			if predictions[i] { // Predicted True
				fp += 1
			} else { // Predicted False
				tn += 1
			}
		}
	}
	return
}

func predictTally(forest forest, data [][]float32) (tp, fn, fp, tn uint) {
	var predictions []bool
	var truth []bool

	tree := forest.trees[0]
	// printNode(&tree, 0)

	for _, sample := range data {
		node := tree
		for {
			// index := [31]string{"None (leaf)", "Radius Mean", "Texture Mean", "Perimeter Mean", "Area Mean", "Smoothness Mean", "Compactness Mean", "Concavity Mean", "Concave points Mean", "Symmetry Mean", "Fractal dimension Mean", "Radius se", "Texture se", "Perimeter se", "Area se", "Smoothness se", "Compactness se", "Concavity se", "Concave points se", "Symmetry se", "Fractal dimension se", "Radius Worst", "Texture Worst", "Perimeter Worst", "Area Worst", "Smoothness Worst", "Compactness Worst", "Concavity Worst", "Concave points Worst", "Symmetry Worst", "Fractal dimension Worst"} // data.csv column titles
			// // fmt.Printf("node.depth: %v\n", node.depth)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          //////////////
			// fmt.Printf("node.feature: %v\n", index[node.feature])          //////////////
			// fmt.Printf("node.split: %v\n", node.split)                     //////////////
			// fmt.Printf("node.diagnosis: %v\n", node.diagnosis)             //////////////
			// fmt.Printf("sample[node.feature]: %v\n", sample[node.feature]) //////////////
			// fmt.Printf("len(node.data): %v\n\n", len(node.data))                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                //////////////
			if node.feature == 0 {
				// printNode(&node, node.depth) ////////////
				// prediction := node.diagnosis
				predictions = append(predictions, node.diagnosis)

				if sample[0] == 1 { // Malignant
					truth = append(truth, true)
					// fmt.Printf("%v %vM%v %v\n", node.diagnosis, RED, RESET, sample[1]) ////////////////
				} else {
					truth = append(truth, false)
					// fmt.Printf("%v %vB%v %v\n", node.diagnosis, GREEN, RESET, sample[1]) ////////////////
				}
				// fmt.Printf("truth: %v\n", truth) //////////////
				break
			}
			if sample[node.feature] < node.split {
				node = *node.childLeft
			} else {
				node = *node.childRight
			}
			// break //
		}
		// break //
	}
	// for i := 0; i < len(predictions); i++ {
	// 	fmt.Printf("%-3v prediction: %-5v, truth: %v\n", i, predictions[i], truth[i])
	// }
	fmt.Printf("\n") //////////////
	return truthTally(predictions, truth)
}

func predict(forest forest, test_set [][]float32) {
	fmt.Printf("\n%v%vPredict%v\n\n", BOLD, UNDERLINE, RESET)
	tp, fn, fp, tn := predictTally(forest, test_set)
	printMetrics(tp, fn, fp, tn)
	confusionMatrix(tp, fn, fp, tn)
}
