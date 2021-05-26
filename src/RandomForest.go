package forest

import "fmt"

// RandomForest is the main & only exposed function
func RandomForest() {
	// flags := parseArg()
	printHeader( /*flags*/ )

	// Data
	// data := preprocess(flags.dataPath)
	data := preprocess("data.csv")
	train_set, test_set := splitData(data /*, flags*/)

	fmt.Printf("len(train_set): %v\n", len(train_set)) ///////////////////
	fmt.Printf("len(test_set): %v\n", len(test_set))   ///////////////////
	fmt.Printf("Oh Hi!!\n")                            ///////////////////
}
