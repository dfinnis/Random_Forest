package forest

import (
	"fmt"
	"os"
)

// printUsage prints usage & quits
func printUsage() {
	fmt.Printf("\nUsage:\tgo build; ./Random_Forest [DATA.CSV] [-d DEPTH] [-t SIZE] [-s SEED] [-f] [-q] [-h]\n\n")
	fmt.Printf("    [-d]    (--depth)        Provide DEPTH integer, tree depth\n")
	fmt.Printf("    [-t]    (--trees)        Provide SIZE integer, number of trees in forest\n")
	fmt.Printf("    [-s]    (--seed)         Provide SEED integer for randomization e.g. -s 42\n")
	fmt.Printf("    [-f]    (--forest)       Print forest, trees, node by node\n")
	fmt.Printf("    [-q]    (--quiet)        Quiet that shit down\n") //// shit!!!
	fmt.Printf("    [-h]    (--help)         Show usage\n\n")
	os.Exit(1)
}

// usageError prints error message & usage, then quits
func usageError(message, err string) {
	fmt.Printf("%vERROR %v %v%v\n", RED, message, err, RESET)
	printUsage()
}

// errorExit prints error message & quits
func errorExit(message string) {
	fmt.Printf("%vERROR %v%v\n", RED, message, RESET)
	os.Exit(1)
}

// checkError prints error message & quits if error
func checkError(message string, err error) {
	if err != nil {
		fmt.Printf("%vERROR %v %v%v\n", RED, message, err, RESET)
		os.Exit(1)
	}
}
