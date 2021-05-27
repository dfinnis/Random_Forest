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

func splitFeature(set [][]float32, feature uint) (float32, float32) {
	set = sortSet(set, feature)
	bestImpurity := (giniImpurity(set[:1]) * (float32(1) / float32(len(set)))) + (giniImpurity(set[1:]) * (float32(len(set)-(1)) / float32(len(set))))
	bestSplit := (set[0][feature] + set[1][feature]) / 2
	for i := 1; i < len(set)-1; i++ {
		// fmt.Printf("%3v. sample[feature]: %-13v, diagnosis: %v\n", i, set[i][feature], set[i][0]) /////////////

		weightedImpurity := (giniImpurity(set[:i+1]) * (float32(i+1) / float32(len(set)))) + (giniImpurity(set[i+1:]) * (float32(len(set)-(i+1)) / float32(len(set))))
		if weightedImpurity < bestImpurity {
			bestImpurity = weightedImpurity
			bestSplit = (set[i][feature] + set[i+1][feature]) / 2
		}
		// fmt.Printf("weightedImpurity: %v\n", weightedImpurity) //////////////
		// break                                                  /////////
	}
	fmt.Printf("bestImpurity: %v\n", bestImpurity) //////////////
	fmt.Printf("bestSplit: %v\n", bestSplit)       //////////////
	return bestImpurity, bestSplit
}

func train(forest forest, train_set [][]float32) {
	fmt.Printf("\n%v%vTrain Forest%v\n\n", BOLD, UNDERLINE, RESET)
	impurity := giniImpurity(train_set) // impurity of root, daiagnose all Benign
	fmt.Printf("root impurity: %v\n\n", impurity)
	splitFeature(train_set, 1)
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
