package forest

import "fmt"

// truthTally counts true & false, positives & negatives
func truthTally(predictions, truth []bool) (float32, float32, float32, float32) {
	var tp float32 // True Positive		// Predicted True & Is True
	var fn float32 // False Negative	// Predicted False & Is True
	var fp float32 // False Positive	// Predicted True & Is False
	var tn float32 // True Negative		// Predicted False & Is False

	for i := 0; i < len(predictions); i++ {
		if truth[i] { // Is True
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
	return tp, fn, fp, tn
}

func predict(forest forest, test_set [][]float32, flags flags) {
	fmt.Printf("\n%v%vPredict%v\n\n", BOLD, UNDERLINE, RESET)
	var predictions []bool
	var truth []bool
	// for sample := 0; sample < len(test_set); sample++ {

	tree := forest.trees[0]
	// printNode(&tree, 0)

	for _, sample := range test_set {
		// fmt.Printf("sample %v: %v\n", i, sample[0])
		// if sample[0] == 1 { // Malignant
		// 	// fmt.Printf("M\n")
		// }
		// truth := append(truth, bool(sample[0]))
		node := tree
		for {
			// index := [31]string{"None (leaf)", "Radius Mean", "Texture Mean", "Perimeter Mean", "Area Mean", "Smoothness Mean", "Compactness Mean", "Concavity Mean", "Concave points Mean", "Symmetry Mean", "Fractal dimension Mean", "Radius se", "Texture se", "Perimeter se", "Area se", "Smoothness se", "Compactness se", "Concavity se", "Concave points se", "Symmetry se", "Fractal dimension se", "Radius Worst", "Texture Worst", "Perimeter Worst", "Area Worst", "Smoothness Worst", "Compactness Worst", "Concavity Worst", "Concave points Worst", "Symmetry Worst", "Fractal dimension Worst"} // data.csv column titles
			// fmt.Printf("node.depth: %v\n", node.depth)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          //////////////
			// fmt.Printf("node.feature: %v\n", index[node.feature])                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               //////////////
			// fmt.Printf("node.split: %v\n", node.split)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          //////////////
			// fmt.Printf("sample[node.feature]: %v\n", sample[node.feature])                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      //////////////
			// fmt.Printf("len(node.data): %v\n\n", len(node.data))                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                //////////////
			if sample[node.feature] < node.split {
				node = *node.childLeft
			} else {
				node = *node.childRight
			}
			if node.feature == 0 {
				// printNode(&node, node.depth) ////////////
				// prediction := node.diagnosis
				predictions = append(predictions, node.diagnosis)
				// fmt.Printf("prediction: %v\n", prediction) //////////////
				// truth := sample[0]
				if sample[0] == 1 { // Malignant
					truth = append(truth, true)
				} else {
					truth = append(truth, false)
				}
				// fmt.Printf("truth: %v\n", truth) //////////////
				break
			}
			// break //
		}
		// break //
	}
	// for i := 0; i < len(predictions); i++ {
	// 	fmt.Printf("%-3v prediction: %-5v, truth: %v\n", i, predictions[i], truth[i])
	// }
	tp, fn, fp, tn := truthTally(predictions, truth)
	confusionMatrix(tp, fn, fp, tn)
}
