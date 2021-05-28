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

func printTree(current *node, depth int) {
	fmt.Printf("Depth: %v\n\n", depth)

	index := [31]string{"None (leaf)", "Radius Mean", "Texture Mean", "Perimeter Mean", "Area Mean", "Smoothness Mean", "Compactness Mean", "Concavity Mean", "Concave points Mean", "Symmetry Mean", "Fractal dimension Mean", "Radius se", "Texture se", "Perimeter se", "Area se", "Smoothness se", "Compactness se", "Concavity se", "Concave points se", "Symmetry se", "Fractal dimension se", "Radius Worst", "Texture Worst", "Perimeter Worst", "Area Worst", "Smoothness Worst", "Compactness Worst", "Concavity Worst", "Concave points Worst", "Symmetry Worst", "Fractal dimension Worst"} // data.csv column titles

	var sum float32
	for i := 0; i < len(current.data); i++ {
		sum += current.data[i][0]
	}
	var diagnosis bool
	if sum/float32(len(current.data)) > 0.5 {
		diagnosis = true
	}

	fmt.Printf("+-----------+-------------------------+\n")
	fmt.Printf("| Feature   | %-23v |\n", index[current.feature])
	if current.split == 0 {
		fmt.Printf("| Split     | None (leaf)             |\n")
	} else {
		fmt.Printf("| Split     | %-23v |\n", current.split)
	}
	fmt.Printf("| Gini      | %-23v |\n", giniImpurity(current.data))
	fmt.Printf("| Samples   | %-23v |\n", len(current.data))
	fmt.Printf("| Value     | %-3v, %-18v |\n", len(current.data)-int(sum), sum)
	if diagnosis {
		fmt.Printf("| Diagnosis |%v Malignant               %v|\n", RED, RESET)
	} else {
		fmt.Printf("| Diagnosis |%v Benign                  %v|\n", GREEN, RESET)
	}
	fmt.Printf("+-----------+-------------------------+\n\n")

	if current.childLeft != nil {
		printTree(current.childLeft, depth+1)
	}
	if current.childRight != nil {
		printTree(current.childRight, depth+1)
	}
}

func printNode(current *node, depth int) {

	index := [31]string{"None (leaf)", "Radius Mean", "Texture Mean", "Perimeter Mean", "Area Mean", "Smoothness Mean", "Compactness Mean", "Concavity Mean", "Concave points Mean", "Symmetry Mean", "Fractal dimension Mean", "Radius se", "Texture se", "Perimeter se", "Area se", "Smoothness se", "Compactness se", "Concavity se", "Concave points se", "Symmetry se", "Fractal dimension se", "Radius Worst", "Texture Worst", "Perimeter Worst", "Area Worst", "Smoothness Worst", "Compactness Worst", "Concavity Worst", "Concave points Worst", "Symmetry Worst", "Fractal dimension Worst"} // data.csv column titles

	var sum float32
	for i := 0; i < len(current.data); i++ {
		sum += current.data[i][0]
	}
	var diagnosis bool
	if sum/float32(len(current.data)) > 0.5 {
		diagnosis = true
	}

	fmt.Printf("Depth: %v\n", depth)
	fmt.Printf("+-----------+-------------------------+\n")
	fmt.Printf("| Feature   | %-23v |\n", index[current.feature])
	if current.split == 0 {
		fmt.Printf("| Split     | None (leaf)             |\n")
	} else {
		fmt.Printf("| Split     | %-23v |\n", current.split)
	}
	fmt.Printf("| Gini      | %-23v |\n", giniImpurity(current.data))
	fmt.Printf("| Samples   | %-23v |\n", len(current.data))
	fmt.Printf("| Value     | %-3v, %-18v |\n", len(current.data)-int(sum), sum)
	if diagnosis {
		fmt.Printf("| Diagnosis |%v Malignant               %v|\n", RED, RESET)
	} else {
		fmt.Printf("| Diagnosis |%v Benign                  %v|\n", GREEN, RESET)
	}
	fmt.Printf("+-----------+-------------------------+\n\n")
}
