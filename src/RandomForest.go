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
		return
	}

	current.childLeft = &node{}
	current.childRight = &node{}

	current.impurity = giniImpurity(current.data)
	// fmt.Printf("root Impurity: %v\n\n", current.impurity) /////////////
	for feature := 1; feature < len(current.data[0]); feature++ {
		// fmt.Printf("feature: %v\n", feature) ////////////
		impurity, split, left, right := splitFeature(current.data, feature)
		if impurity < current.impurity {
			current.impurity = impurity
			current.split = split
			current.feature = feature
			current.childLeft.data = left
			current.childRight.data = right
		}
	}
	fmt.Printf("bestFeature: %v\n", current.feature)   /////////
	fmt.Printf("bestImpurity: %v\n", current.impurity) /////////
	fmt.Printf("bestSplit: %v\n\n", current.split)     /////////

	splitNode(current.childLeft, depth-1)
	splitNode(current.childRight, depth-1)
}

func train(forest forest, train_set [][]float32, depth int) {
	fmt.Printf("\n%v%vTrain Forest%v\n\n", BOLD, UNDERLINE, RESET)
	forest.trees[0].data = train_set
	splitNode(&forest.trees[0], depth)
}

// RandomForest is the main & only exposed function
func RandomForest() {
	flags := parseArg()
	printHeader(flags)

	// Data
	data := preprocess(flags.dataPath)
	train_set, test_set := splitData(data /*, flags*/)
	fmt.Printf("len(train_set): %v\n", len(train_set)) ///////////////////
	fmt.Printf("len(test_set): %v\n", len(test_set))   ///////////////////

	// Initialize
	forest := initForest()

	// Train
	train(forest, train_set, 2)

	// Predict

	fmt.Printf("forest.trees[0].feature: %v\n", forest.trees[0].feature)                     ///////////////////
	fmt.Printf("forest.trees[0].impurity: %v\n", forest.trees[0].impurity)                   ///////////////////
	fmt.Printf("forest.trees[0].split: %v\n", forest.trees[0].split)                         ///////////////////
	fmt.Printf("len(forest.trees[0].data: %v\n", len(forest.trees[0].data))                  ///////////////////
	fmt.Printf("len(forest.trees[0].childLeft: %v\n", len(forest.trees[0].childLeft.data))   ///////////////////
	fmt.Printf("len(forest.trees[0].childRight: %v\n", len(forest.trees[0].childRight.data)) ///////////////////

	fmt.Printf("Oh Hi!!\n") ///////////////////
}
