package forest

import "fmt"

const RESET = "\x1B[0m"
const BOLD = "\x1B[1m"
const UNDERLINE = "\x1B[4m"
const RED = "\x1B[31m"
const GREEN = "\x1B[32m"
const BLUE = "\x1B[34m"
const CYAN = "\x1B[36m"

// printHeader prints intro
func printHeader(flags flags) {
	if !flags.flagQ {
		fmt.Printf("\033[H\033[2J") // Clear screen
	} else {
		fmt.Printf("\n")
	}
	fmt.Printf("%v%vLaunching Random Forest%v\n\n", BOLD, UNDERLINE, RESET)
}

// printSplit shows how the data is split between training & test set
func printSplit(train_set, test_set int) {
	fmt.Printf("+--------------+---------+\n")
	fmt.Printf("|%v Data Split   %v|%v Samples %v|\n", BOLD, RESET, BOLD, RESET)
	fmt.Printf("+--------------+---------+\n")
	fmt.Printf("| Training set | %-7v |\n", train_set)
	fmt.Printf("| Test set     | %-7v |\n", test_set)
	fmt.Printf("+--------------+---------+\n\n")
}

// printNode prints info for one node
func printNode(current *node) {

	index := [31]string{"None (leaf)", "Radius Mean", "Texture Mean", "Perimeter Mean", "Area Mean", "Smoothness Mean", "Compactness Mean", "Concavity Mean", "Concave points Mean", "Symmetry Mean", "Fractal dimension Mean", "Radius se", "Texture se", "Perimeter se", "Area se", "Smoothness se", "Compactness se", "Concavity se", "Concave points se", "Symmetry se", "Fractal dimension se", "Radius Worst", "Texture Worst", "Perimeter Worst", "Area Worst", "Smoothness Worst", "Compactness Worst", "Concavity Worst", "Concave points Worst", "Symmetry Worst", "Fractal dimension Worst"} // data.csv column titles

	var malignant float32
	for i := 0; i < len(current.data); i++ {
		malignant += current.data[i][0]
	}

	fmt.Printf("+-----------+------------------------------+\n")
	if current.depth == 0 {
		fmt.Printf("| Depth     | %v0%v - %vRoot                     %v|\n", BOLD, RESET, CYAN, RESET)
	} else if current.feature == 0 {
		fmt.Printf("| Depth     | %v%v%v - %vLeaf                     %v|\n", BOLD, current.depth, RESET, BLUE, RESET)
	} else {
		fmt.Printf("| Depth     | %v%-28v %v|\n", BOLD, current.depth, RESET)
	}

	if current.feature != 0 {
		fmt.Printf("| Feature   | %-2v - %-23v |\n", current.feature, index[current.feature])
		fmt.Printf("| Split     | %-28v |\n", current.split)
	}
	fmt.Printf("| Gini      | %-28v |\n", giniImpurity(current.data))
	fmt.Printf("| Samples   | %-28v |\n", len(current.data))
	fmt.Printf("| Value     | %v%-3v%v, %v%-23v%v |\n", GREEN, len(current.data)-int(malignant), RESET, RED, malignant, RESET)
	if current.diagnosis {
		fmt.Printf("| Diagnosis |%v Malignant                    %v|\n", RED, RESET)
	} else {
		fmt.Printf("| Diagnosis |%v Benign                       %v|\n", GREEN, RESET)
	}
	fmt.Printf("+-----------+------------------------------+\n\n")
}

// printTree recursively explores & prints each node
func printTree(current *node) {
	printNode(current)
	if current.childLeft != nil {
		printTree(current.childLeft)
	}
	if current.childRight != nil {
		printTree(current.childRight)
	}
}

func printTrees(trees []node) {
	for _, tree := range trees {
		printTree(&tree)
		fmt.Printf("\n\n")
	}
}

// printForest prints info about trees
func printForest(treeInfos []treeInfo) {
	fmt.Printf("+--------------+-------------------------+\n")
	fmt.Printf("| Trees        | %v |\n", len(treeInfos))
	fmt.Printf("+--------------+-------------------------+\n")
	fmt.Printf("| Samples      | %v\n", treeInfos[0].samples)
	fmt.Printf("| Depth        | %v\n", treeInfos[0].depth)
	fmt.Printf("| Nodes        | %v\n", treeInfos[0].nodes)
	fmt.Printf("| Leafs        | %v\n", treeInfos[0].leafs)
	fmt.Printf("| Gini mean    | %v\n", treeInfos[0].impurity)
	fmt.Printf("| Samples/Leaf | %v\n", treeInfos[0].samples)

	// Samples / leaf     min mean max
	// leaf gini          min mean max
	fmt.Printf("+--------------+-------------------------+\n\n")
}

// getMetrics converts true & false, positives & negatives into metrics
func getMetrics(tpUint, fnUint, fpUint, tnUint uint) (accuracy, precision, recall, specificity, F1_score float32) {
	tp := float32(tpUint)
	fn := float32(fnUint)
	fp := float32(fpUint)
	tn := float32(tnUint)

	accuracy = (tp + tn) / (tp + tn + fp + fn)
	precision = tp / (tp + fp)
	recall = tp / (tp + fn)
	specificity = tn / (tn + fp)
	F1_score = (2 * (precision * recall)) / (precision + recall)
	if tp == 0 {
		precision = 0
		F1_score = 0
	}
	return
}

// printTrain prints metrics & confusion matrix for training & test sets
func printTrain(forest forest, train_set, test_set [][]float32) {
	tpTrain, fnTrain, fpTrain, tnTrain := predictTally(forest, train_set)
	tpTest, fnTest, fpTest, tnTest := predictTally(forest, test_set)

	accuracyTrain, precisionTrain, recallTrain, specificityTrain, F1_scoreTrain := getMetrics(tpTrain, fnTrain, fpTrain, tnTrain)
	accuracyTest, precisionTest, recallTest, specificityTest, F1_scoreTest := getMetrics(tpTest, fnTest, fpTest, tnTest)

	fmt.Printf("+-----------------+---------------+---------------+\n")
	fmt.Printf("|%v Metric          %v|%v Training Set  %v|%v Test Set      %v|\n", BOLD, RESET, BOLD, RESET, BOLD, RESET)
	fmt.Printf("+-----------------+---------------+---------------+\n")
	fmt.Printf("|%v        Accuracy %v| %-8f      | %-8f      |\n", BOLD, RESET, accuracyTrain, accuracyTest)
	fmt.Printf("|                 |               |               |\n")
	fmt.Printf("|%v       Precision %v| %-8f      | %-8f      |\n", BOLD, RESET, precisionTrain, precisionTest)
	fmt.Printf("|                 |               |               |\n")
	fmt.Printf("|%v          Recall %v| %-8f      | %-8f      |\n", BOLD, RESET, recallTrain, recallTest)
	fmt.Printf("|                 |               |               |\n")
	fmt.Printf("|%v     Specificity %v| %-8f      | %-8f      |\n", BOLD, RESET, specificityTrain, specificityTest)
	fmt.Printf("|                 |               |               |\n")
	fmt.Printf("|%v        F1_score %v| %-8f      | %-8f      |\n", BOLD, RESET, F1_scoreTrain, F1_scoreTest)
	fmt.Printf("+-----------------+---------------+---------------+\n\n")
	confusionMatrix2(tpTrain, fnTrain, fpTrain, tnTrain, tpTest, fnTest, fpTest, tnTest)
}

// confusionMatrix2 shows true & false, positives & negatives for training & test sets
func confusionMatrix2(tpTrain, fnTrain, fpTrain, tnTrain, tpTest, fnTest, fpTest, tnTest uint) {
	fmt.Printf("%vConfusion Matrix%v  +---------------+---------------+\n", BOLD, RESET)
	fmt.Printf("                  ǁ%v Ground Truth %v ǁ%v Ground Truth %v ǁ\n", BOLD, RESET, BOLD, RESET)
	fmt.Printf("                  ǁ-------+-------ǁ-------+-------ǁ\n")
	fmt.Printf("                  ǁ%v%v True %v |%v%v False %vǁ%v%v True %v |%v%v False %vǁ\n", BOLD, GREEN, RESET, BOLD, RED, RESET, BOLD, GREEN, RESET, BOLD, RED, RESET)
	fmt.Printf("+-----------------ǁ---------------ǁ---------------ǁ\n")
	fmt.Printf("|         |%v%v True %v ǁ%v %-5v %v|%v %-5v %vǁ%v %-5v %v|%v %-5v %vǁ\n", BOLD, GREEN, RESET, GREEN, tpTrain, RESET, RED, fpTrain, RESET, GREEN, tpTest, RESET, RED, fpTest, RESET)
	fmt.Printf("|%v Predict %v+-------ǁ-------+-------ǁ-------+-------ǁ\n", BOLD, RESET)
	fmt.Printf("|         |%v%v False %vǁ%v %-5v %v|%v %-5v %vǁ%v %-5v %v|%v %-5v %vǁ\n", BOLD, RED, RESET, RED, fnTrain, RESET, GREEN, tnTrain, RESET, RED, fnTest, RESET, GREEN, tnTest, RESET)
	fmt.Printf("+-----------------+---------------+---------------+\n\n")
}

// printMetrics shows the final metrics
func printMetrics(tp, fn, fp, tn uint) {
	accuracy, precision, recall, specificity, F1_score := getMetrics(tp, fn, fp, tn)

	fmt.Printf("+-------------+----------+-------------------------------------------------------------------------+\n")
	fmt.Printf("|%v Metric      %v|%v Value    %v|%v Description                                                             %v|\n", BOLD, RESET, BOLD, RESET, BOLD, RESET)
	fmt.Printf("+-------------+----------+-------------------------------------------------------------------------+\n")
	fmt.Printf("|%v    Accuracy %v| %-8f | proportion of predictions classified correctly                          |\n", BOLD, RESET, accuracy)
	fmt.Printf("|             |          |                                                                         |\n")
	fmt.Printf("|%v   Precision %v| %-8f | proportion of positive identifications correct                          |\n", BOLD, RESET, precision)
	fmt.Printf("|             |          |                                                                         |\n")
	fmt.Printf("|%v      Recall %v| %-8f | proportion of actual positives correctly identified. True Positive Rate |\n", BOLD, RESET, recall)
	fmt.Printf("|             |          |                                                                         |\n")
	fmt.Printf("|%v Specificity %v| %-8f | proportion of actual negatives correctly identified. True Negative Rate |\n", BOLD, RESET, specificity)
	fmt.Printf("|             |          |                                                                         |\n")
	fmt.Printf("|%v    F1_score %v| %-8f | harmonic mean of precision & recall. Max 1 (perfect), min 0             |\n", BOLD, RESET, F1_score)
	fmt.Printf("+-------------+----------+-------------------------------------------------------------------------+\n\n\n")
}

// confusionMatrix shows true & false, positives & negatives for the test set
func confusionMatrix(tp, fn, fp, tn uint) {
	fmt.Printf("%vConfusion Matrix%v  +---------------+\n", BOLD, RESET)
	fmt.Printf("                  ǁ%v Ground Truth %v ǁ\n", BOLD, RESET)
	fmt.Printf("Total: %-4v       ǁ-------+-------ǁ\n", (tp + fn + fp + tn))
	fmt.Printf("                  ǁ%v%v True %v |%v%v False %vǁ\n", BOLD, GREEN, RESET, BOLD, RED, RESET)
	fmt.Printf("+-----------------ǁ---------------ǁ\n")
	fmt.Printf("|         |%v%v True %v ǁ%v %-5v %v|%v %-5v %vǁ\n", BOLD, GREEN, RESET, GREEN, tp, RESET, RED, fp, RESET)
	fmt.Printf("|%v Predict %v+-------ǁ-------+-------ǁ\n", BOLD, RESET)
	fmt.Printf("|         |%v%v False %vǁ%v %-5v %v|%v %-5v %vǁ\n", BOLD, RED, RESET, RED, fn, RESET, GREEN, tn, RESET)
	fmt.Printf("+-----------------+---------------+\n\n")
}
