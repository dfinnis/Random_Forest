package forest

import (
	"fmt"
)

// truthTally counts true & false, positives & negatives
func truthTally(predictions, truth []bool) (tp, fn, fp, tn uint) {
	// var tp uint // True Positive		// Predicted True & Is True
	// var fn uint // False Negative	// Predicted False & Is True
	// var fp uint // False Positive	// Predicted True & Is False
	// var tn uint // True Negative		// Predicted False & Is False

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
	return
}

// predictTally predicts & returns true & false, positives & negatives
func predictTally(forest forest, data [][]float32) (tp, fn, fp, tn uint) {
	var predictions []bool
	var truth []bool
	for _, sample := range data {
		var prediction []bool
		for _, tree := range forest.trees {
			node := tree
			for {
				// Predict at leaf
				if node.feature == 0 {
					prediction = append(prediction, node.diagnosis)
					break
				}
				// Move deeper in Tree
				if sample[node.feature] < node.split {
					node = *node.childLeft
				} else {
					node = *node.childRight
				}
			}
		}
		// Count votes
		var count float32
		for _, pred := range prediction {
			if pred {
				count += 1
			}
		}
		if count/float32(len(prediction)) > 0.5 {
			predictions = append(predictions, true)
		} else {
			predictions = append(predictions, false)
		}

		// Truth
		if sample[0] == 1 { // Malignant
			truth = append(truth, true)
		} else {
			truth = append(truth, false)
		}
	}
	return truthTally(predictions, truth)
}

// predict prints metrics & confusion matrix for test_set
func predict(forest forest, test_set [][]float32) {
	fmt.Printf("\n%v%vPredict%v\n\n", BOLD, UNDERLINE, RESET)
	tp, fn, fp, tn := predictTally(forest, test_set)
	printMetrics(tp, fn, fp, tn)
	confusionMatrix(tp, fn, fp, tn)
}
