package forest

import (
	"fmt"
)

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
	train(forest, train_set, flags)

	// Predict
	predict(forest, test_set, flags)

	// printTree(&forest.trees[0], 0)
	fmt.Printf("Oh Hi!!\n") ///////////////////
}
