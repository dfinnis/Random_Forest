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

// printTrees prints all trees in given forest
func printTrees(trees []node) {
	for _, tree := range trees {
		printTree(&tree)
		fmt.Printf("\n\n")
	}
}

// rangeDepth finds min, mean & max depth for given forest
func rangeDepth(treeInfos []treeInfo) (int, float32, int) {
	depthMin := treeInfos[0].depth
	depthMax := treeInfos[0].depth
	var depthTotal int
	for _, tree := range treeInfos {
		depthTotal += tree.depth
		if tree.depth < depthMin {
			depthMin = tree.depth
		}
		if tree.depth > depthMax {
			depthMax = tree.depth
		}
	}
	depthMean := float32(depthTotal) / float32(len(treeInfos))
	return depthMin, depthMean, depthMax
}

// rangeNodes finds min, mean & max nodes for given forest
func rangeNodes(treeInfos []treeInfo) (uint, float32, uint) {
	nodesMin := treeInfos[0].nodes
	nodesMax := treeInfos[0].nodes
	var nodesTotal uint
	for _, tree := range treeInfos {
		nodesTotal += tree.nodes
		if tree.nodes < nodesMin {
			nodesMin = tree.nodes
		}
		if tree.nodes > nodesMax {
			nodesMax = tree.nodes
		}
	}
	nodesMean := float32(nodesTotal) / float32(len(treeInfos))
	return nodesMin, nodesMean, nodesMax
}

// rangeLeafs finds min, mean & max leafs for given forest
func rangeLeafs(treeInfos []treeInfo) (uint, float32, uint) {
	leafsMin := treeInfos[0].leafs
	leafsMax := treeInfos[0].leafs
	var leafsTotal uint
	for _, tree := range treeInfos {
		leafsTotal += tree.leafs
		if tree.leafs < leafsMin {
			leafsMin = tree.leafs
		}
		if tree.leafs > leafsMax {
			leafsMax = tree.leafs
		}
	}
	leafsMean := float32(leafsTotal) / float32(len(treeInfos))
	return leafsMin, leafsMean, leafsMax
}

// rangeSamples finds min, mean & max samples per leaf for given forest
func rangeSamples(treeInfos []treeInfo) (float32, float32, float32) {
	samplesMin := treeInfos[0].samplesLeaf
	samplesMax := treeInfos[0].samplesLeaf
	var samplesTotal float32
	for _, tree := range treeInfos {
		samplesTotal += tree.samplesLeaf
		if tree.samplesLeaf < samplesMin {
			samplesMin = tree.samplesLeaf
		}
		if tree.samplesLeaf > samplesMax {
			samplesMax = tree.samplesLeaf
		}
	}
	samplesMean := float32(samplesTotal) / float32(len(treeInfos))
	return samplesMin, samplesMean, samplesMax
}

// rangeImpurity finds min, mean & max gini impurity for given forest
func rangeImpurity(treeInfos []treeInfo) (float32, float32, float32) {
	impurityMin := treeInfos[0].impurity
	impurityMax := treeInfos[0].impurity
	var impurityTotal float32
	for _, tree := range treeInfos {
		impurityTotal += tree.impurity
		if tree.impurity < impurityMin {
			impurityMin = tree.impurity
		}
		if tree.impurity > impurityMax {
			impurityMax = tree.impurity
		}
	}
	impurityMean := float32(impurityTotal) / float32(len(treeInfos))
	return impurityMin, impurityMean, impurityMax
}

// printForest prints info about trees
func printForest(treeInfos []treeInfo) {
	depthMin, depthMean, depthMax := rangeDepth(treeInfos)
	nodesMin, nodesMean, nodesMax := rangeNodes(treeInfos)
	leafsMin, leafsMean, leafsMax := rangeLeafs(treeInfos)
	samplesMin, samplesMean, samplesMax := rangeSamples(treeInfos)
	impurityMin, impurityMean, impurityMax := rangeImpurity(treeInfos)

	// fmt.Printf("|%v Samples/Tree %v| %-18v |\n", BOLD, RESET, treeInfos[0].samples)
	fmt.Printf("%vTrees: %v%-7v +------+------+------+\n", BOLD, RESET, len(treeInfos))
	fmt.Printf("               |%v Min  %v|%v Mean %v|%v Max  %v|\n", BOLD, RESET, BOLD, RESET, BOLD, RESET)
	fmt.Printf("+--------------+------+------+------+\n")
	if depthMin == depthMax {
		fmt.Printf("|%v        Depth %v| %-4v | %-4v | %-4v |\n", BOLD, RESET, depthMin, depthMean, depthMax)
	} else {
		fmt.Printf("|%v        Depth %v| %-4v | %-4.1f | %-4v |\n", BOLD, RESET, depthMin, depthMean, depthMax)
	}
	fmt.Printf("|              |      |      |      |\n")
	fmt.Printf("|%v        Nodes %v| %-4v | %-4.1f | %-4v |\n", BOLD, RESET, nodesMin, nodesMean, nodesMax)
	fmt.Printf("|              |      |      |      |\n")
	fmt.Printf("|%v        Leafs %v| %-4v | %-4.1f | %-4v |\n", BOLD, RESET, leafsMin, leafsMean, leafsMax)
	fmt.Printf("|              |      |      |      |\n")
	if depthMax == 1 {
		fmt.Printf("|%v Samples/Leaf %v| %-4.0f | %-4.0f | %-4.0f |\n", BOLD, RESET, samplesMin, samplesMean, samplesMax)
	} else {
		fmt.Printf("|%v Samples/Leaf %v| %-4.1f | %-4.1f | %-4.1f |\n", BOLD, RESET, samplesMin, samplesMean, samplesMax)
	}
	fmt.Printf("|              |      |      |      |\n")
	fmt.Printf("|%v    Gini mean %v| %-4.2f | %-4.2f | %-4.2f |\n", BOLD, RESET, impurityMin, impurityMean, impurityMax)
	fmt.Printf("+--------------+--------------------+\n\n")
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
