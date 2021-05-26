package forest

import "fmt"

type node struct {
	feature   uint8
	split     float32
	diagnosis bool
	childA    *node
	childB    *node
}

type forest struct {
	// depth uint
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
	forest.trees = append(forest.trees, node{})
	return forest
}

// RandomForest is the main & only exposed function
func RandomForest() {
	flags := parseArg()
	printHeader(flags)

	// Data
	// data := preprocess(flags.dataPath)
	data := preprocess(flags.dataPath)
	train_set, test_set := splitData(data /*, flags*/)
	fmt.Printf("len(train_set): %v\n", len(train_set)) ///////////////////
	fmt.Printf("len(test_set): %v\n", len(test_set))   ///////////////////

	// Initialize
	forest := initForest()
	fmt.Printf("forest: %v\n", forest) ///////////////////

	fmt.Printf("Oh Hi!!\n") ///////////////////
}
