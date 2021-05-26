package forest

import "fmt"

const RESET = "\x1B[0m"
const BOLD = "\x1B[1m"
const UNDERLINE = "\x1B[4m"
const RED = "\x1B[31m"
const GREEN = "\x1B[32m"

// printHeader prints intro
func printHeader(flags flags) {
	if !flags.flagQ {
		fmt.Printf("\033[H\033[2J") // Clear screen
	} else {
		fmt.Printf("\n")
	}
	fmt.Printf("%v%vLaunching Random Forest%v\n\n", BOLD, UNDERLINE, RESET)
}

// printSplit shows how the data is split between training & test set
func printSplit(train_set, test_set int) {
	fmt.Printf("+--------------+---------+\n")
	fmt.Printf("|%v Data Split   %v|%v Samples %v|\n", BOLD, RESET, BOLD, RESET)
	fmt.Printf("+--------------+---------+\n")
	fmt.Printf("| Training set | %-7v |\n", train_set)
	fmt.Printf("| Test set     | %-7v |\n", test_set)
	fmt.Printf("+--------------+---------+\n\n")
}
