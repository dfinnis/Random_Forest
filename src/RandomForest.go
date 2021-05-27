package forest

import (
	"fmt"
	"sort"
)

type node struct {
	feature    uint8 // data column
	split      float32
	childLeft  *node
	childRight *node
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

func sortSet(set [][]float32, feature uint) [][]float32 {
	sort.SliceStable(set, func(i, j int) bool {
		return set[i][feature] < set[j][feature]
	})
	return set
}

func giniImpurity(set [][]float32, feature uint, split float32) float32 {
	var impurity float32
	set = sortSet(set, feature)
	for i, sample := range set {
		fmt.Printf("%3v. sample[feature]: %-13v, diagnosis: %v\n", i, sample[feature], sample[0])
	}
	return impurity
}

func train(forest forest, train_set [][]float32) {
	fmt.Printf("\n%v%vTrain Forest%v\n\n", BOLD, UNDERLINE, RESET)
	impurity := giniImpurity(train_set, 1, 0)
	fmt.Printf("impurity: %v\n\n", impurity)
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
	// fmt.Printf("forest: %v\n", forest) ///////////////////

	// Train
	train(forest, train_set)

	// Predict

	fmt.Printf("Oh Hi!!\n") ///////////////////
}
