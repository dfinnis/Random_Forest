package forest

// RandomForest is the main & only exposed function
func RandomForest() {
	flags := parseArg()
	printHeader(flags)

	// Data
	data := preprocess(flags.dataPath)
	train_set, test_set := splitData(data /*, flags*/)

	// Initialize
	forest := initForest()

	// Train
	train(forest, train_set, test_set, flags)

	// Predict
	// predict(forest, test_set)

	// printTree(&forest.trees[0], 0)
}
