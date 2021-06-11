package forest

import (
	"fmt"
	"math/rand"
	"sort"
)

type treeInfo struct {
	nodes       uint
	leafs       uint
	samples     int     // total
	samplesLeaf float32 // per leaf              min mean max?
	impurity    float32 // mean for leafs        min mean max?
	depth       int     // deepest                  min mean max?
}

type node struct {
	depth      int
	feature    int // data column // uint8????!!!!!!!!!!!!!
	impurity   float32
	split      float32
	childLeft  *node
	childRight *node
	data       [][]float32
	diagnosis  bool // majority vote
}

type forest struct {
	trees []node
}

func sortSet(set [][]float32, feature int) [][]float32 {
	var sorted [][]float32
	sorted = append(sorted, set...)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i][feature] < sorted[j][feature]
	})
	return sorted
}

func giniImpurity(set [][]float32) float32 {
	var diagnosis float32
	var i int
	for ; i < len(set); i++ {
		diagnosis += set[i][0]
	}
	diagnosis /= float32(i)
	antiDiagnosis := 1 - diagnosis
	impurity := 1 - ((diagnosis * diagnosis) + (antiDiagnosis * antiDiagnosis))
	return impurity
}

func splitFeature(set [][]float32, feature int) (float32, float32, [][]float32, [][]float32) {
	set = sortSet(set, feature)

	var bestImpurity float32 = 1
	var bestSplit float32 = 0
	var bestLeft [][]float32
	var bestRight [][]float32

	for i := 0; i < len(set)-1; i++ {
		dataLeft := set[:i+1]
		dataRight := set[i+1:]
		weightedImpurity := (giniImpurity(dataLeft) * (float32(i+1) / float32(len(set)))) + (giniImpurity(dataRight) * (float32(len(set)-(i+1)) / float32(len(set))))
		if weightedImpurity < bestImpurity {
			bestImpurity = weightedImpurity
			bestSplit = (set[i][feature] + set[i+1][feature]) / 2
			bestLeft = dataLeft
			bestRight = dataRight
		}
	}
	return bestImpurity, bestSplit, bestLeft, bestRight
}

func recordLeaf(current *node, treeInfo *treeInfo, flagF bool) {
	treeInfo.leafs += 1
	if current.depth > treeInfo.depth {
		treeInfo.depth = current.depth

	}
	treeInfo.samplesLeaf += float32(len(current.data))
	treeInfo.impurity += current.impurity
}

func diagnoseNode(current *node, currentDepth, depth int) {
	var sum float32
	for i := 0; i < len(current.data); i++ {
		sum += current.data[i][0]
	}
	if sum/float32(len(current.data)) > 0.5 {
		current.diagnosis = true
	}
}

func splitNode(current *node, currentDepth, depth int, flagF bool, treeInfo *treeInfo) {
	treeInfo.nodes += 1
	current.depth = currentDepth
	current.impurity = giniImpurity(current.data)

	diagnoseNode(current, currentDepth, depth)

	if currentDepth >= depth {
		// fmt.Printf("depth <= 0\n") ////////////
		recordLeaf(current, treeInfo, flagF)
		return
	}

	var bestFeature int
	var bestSplit float32
	var bestLeft [][]float32
	var bestRight [][]float32
	bestImpurity := current.impurity

	if bestImpurity == 0 {
		// fmt.Printf("current.impurity == 0\n") ////////////
		recordLeaf(current, treeInfo, flagF)
		return
	}
	for feature := 1; feature < len(current.data[0]); feature++ {
		impurity, split, left, right := splitFeature(current.data, feature)
		if impurity < bestImpurity {
			bestFeature = feature
			bestImpurity = impurity
			bestSplit = split
			bestLeft = left
			bestRight = right
		}
		if impurity == 0 {
			break
		}
	}

	if len(bestLeft) == 0 || len(bestRight) == 0 {
		// fmt.Printf("len(bestLeft): %v  len(bestRight): %v\n", len(bestLeft), len(bestRight)) ////////////
		recordLeaf(current, treeInfo, flagF)
		return
	}
	current.impurity = bestImpurity
	current.feature = bestFeature
	current.split = bestSplit

	current.childLeft = &node{}
	current.childRight = &node{}
	current.childLeft.data = bestLeft
	current.childRight.data = bestRight

	splitNode(current.childLeft, currentDepth+1, depth, flagF, treeInfo)
	splitNode(current.childRight, currentDepth+1, depth, flagF, treeInfo)
}

func splitSubset(forest forest, i int, train_set [][]float32, size int) {
	split := 0.5 // proportion of training set given to each tree
	if size == 1 {
		split = 1
	}
	rand.Shuffle(len(train_set), func(i, j int) { train_set[i], train_set[j] = train_set[j], train_set[i] })
	var subset [][]float32
	subset = append(subset, train_set[:int(float64(len(train_set))*split)]...)
	forest.trees[i].data = subset
}

func train(forest forest, train_set, test_set [][]float32, flags flags) {
	fmt.Printf("\n%v%vTrain Forest%v\n\n", BOLD, UNDERLINE, RESET)
	var treeInfos []treeInfo

	for i := 0; i < flags.size; i++ {
		fmt.Printf("%v Training tree %v %v/ %v\r", BOLD, i+1, RESET, flags.size) //////////////
		forest.trees = append(forest.trees, node{})                              // root
		treeInfo := treeInfo{}
		splitSubset(forest, i, train_set, flags.size)
		treeInfo.samples = len(forest.trees[i].data)
		splitNode(&forest.trees[i], 0, flags.depth, flags.flagF, &treeInfo)
		treeInfo.samplesLeaf /= float32(treeInfo.leafs)
		treeInfo.impurity /= float32(treeInfo.leafs) // wrong!! needs weighting by number of samples !!!!!!!!!!!!!
		treeInfos = append(treeInfos, treeInfo)
	}
	printForest(treeInfos)
	if flags.flagF {
		printTrees(forest.trees)
	}
	printTrain(forest, train_set, test_set)
}
