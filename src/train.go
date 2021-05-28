package forest

import (
	"fmt"
	"sort"
)

type node struct {
	feature    int // data column // uint8????!!!!!!!!!!!!!
	impurity   float32
	split      float32
	childLeft  *node
	childRight *node
	data       [][]float32
	// diagnosis  bool // above or below split = Malignant
}

type forest struct {
	// depth uint
	// size  uint
	trees []node
}

// // newNode initializes a node
// func newNode(forest forest) forest {
// 	node := node{}
// 	forest.trees = append(forest.trees, node)
// 	return forest
// }

func initForest() forest {
	forest := forest{}
	forest.trees = append(forest.trees, node{}) // root
	return forest
}

func sortSet(set [][]float32, feature int) [][]float32 {
	sort.SliceStable(set, func(i, j int) bool {
		return set[i][feature] < set[j][feature]
	})
	return set
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
	bestImpurity := (giniImpurity(set[:1]) * (float32(1) / float32(len(set)))) + (giniImpurity(set[1:]) * (float32(len(set)-(1)) / float32(len(set))))
	bestSplit := (set[0][feature] + set[1][feature]) / 2
	var bestLeft [][]float32
	var bestRight [][]float32
	for i := 1; i < len(set)-1; i++ {
		// fmt.Printf("%3v. sample[feature]: %-13v, diagnosis: %v\n", i, set[i][feature], set[i][0]) /////////////
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
	// fmt.Printf("bestImpurity: %v\n", bestImpurity) //////////////
	// fmt.Printf("bestSplit: %v\n\n", bestSplit)     //////////////
	return bestImpurity, bestSplit, bestLeft, bestRight
}

func splitNode(current *node, depth int) {
	if depth <= 0 {
		fmt.Printf("depth <= 0\n") ////////////
		printNode(current, depth)  ////////////////////////????????
		return
	}

	var bestFeature int
	var bestSplit float32
	var bestLeft [][]float32
	var bestRight [][]float32

	current.impurity = giniImpurity(current.data)
	bestImpurity := current.impurity
	if bestImpurity == 0 {
		fmt.Printf("current.impurity == 0\n") ////////////
		printNode(current, depth)             ////////////////////////????????
		return
	}
	// fmt.Printf("root Impurity: %v\n\n", current.impurity) /////////////
	for feature := 1; feature < len(current.data[0]); feature++ {
		impurity, split, left, right := splitFeature(current.data, feature)
		// fmt.Printf("feature: %v, impurity: %v\n", feature, impurity) ////////////
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
	// fmt.Printf("bestFeature: %v\n", bestFeature)   /////////
	// fmt.Printf("bestImpurity: %v\n", bestImpurity) /////////
	// fmt.Printf("bestSplit: %v\n\n", bestSplit)     /////////

	if len(bestLeft) == 0 || len(bestRight) == 0 {
		fmt.Printf("len(bestLeft) == 0 || len(bestRight) == 0\n") ////////////
		printNode(current, depth)                                 ////////////////////////????????
		return
	}
	current.impurity = bestImpurity
	current.feature = bestFeature
	current.split = bestSplit
	printNode(current, depth) ////////////////////////????????

	current.childLeft = &node{}
	current.childRight = &node{}
	current.childLeft.data = bestLeft
	current.childRight.data = bestRight
	splitNode(current.childLeft, depth-1)
	splitNode(current.childRight, depth-1)
}

func train(forest forest, train_set [][]float32, flags flags) {
	fmt.Printf("\n%v%vTrain Forest%v\n\n", BOLD, UNDERLINE, RESET)
	forest.trees[0].data = train_set
	splitNode(&forest.trees[0], flags.depth)
}
