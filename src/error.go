package forest

import (
	"fmt"
	"os"
)

// printUsage prints usage & quits
func printUsage() {
	fmt.Printf("\nUsage:\tgo build; ./Random_Forest [DATA.CSV] [-h]\n\n")
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
