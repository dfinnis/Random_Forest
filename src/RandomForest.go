package forest

import (
	"fmt"
	"math/rand"
)

// seedRandom initializes rand with time or -s SEED
func seedRandom(flags flags) {
	rand.Seed(flags.seed)
	if !(flags.flagS || flags.flagQ) {
		fmt.Printf("Random seed: %d\n\n", flags.seed)
	}
}

// RandomForest is the main & only exposed function
func RandomForest() {
	flags := parseArg()
	printHeader(flags)
	seedRandom(flags)

	// Data
	data := preprocess(flags.dataPath)
	train_set, test_set := splitData(data /*, flags*/)

	// Train
	forest := train(train_set, test_set, flags)

	// Predict
	predict(forest, test_set)

	// printTree(&forest.trees[0], 0)
}
