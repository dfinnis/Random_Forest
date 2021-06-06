package forest

import (
	"os"
	"strconv"
)

// flags contains all flags & arguments
type flags struct {
	dataPath    string
	dataPathSet bool
	flagQ       bool
	flagF       bool
	depth       int // tree depth
	// size        int    // number of trees
}

// defaultConfig initializes default values
func defaultConfig() flags {
	flags := flags{}
	flags.dataPath = "data.csv"
	flags.depth = 2 // best number ??
	return flags
}

// parseFilepath checks if filepath exists
func parseFilepath(filepath string) string {
	_, err := os.Stat(filepath)
	if err != nil {
		usageError("Invalid filepath: ", filepath)
	}
	return filepath
}

// parseDataPath sets path to data, default: data.csv. Catches all bad arguments
func parseDataPath(filepath string, flags flags) flags {
	if flags.dataPathSet {
		usageError("Invalid argument: ", filepath)
	}
	flags.dataPath = parseFilepath(filepath)
	flags.dataPathSet = true
	return flags
}

// parseDepth parses string to int, must be between 0 & 100000
func parseDepth(i int, args []string) int {
	if i >= len(args) {
		usageError("No depth number provided after -ep", "")
	}
	depth, err := strconv.Atoi(args[i])
	if err != nil {
		usageError("Bad depth: ", args[i])
	}
	if depth <= 0 || depth >= 100000 {
		usageError("depth must be between 0 & 100000, given: ", args[i])
	}
	return depth
}

// parseArg parses & returns arguments for flags
func parseArg() flags {
	flags := defaultConfig()

	args := os.Args[1:]
	if len(args) == 0 {
		return flags
	} else if len(args) > 3 {
		usageError("Too many arguments: ", strconv.Itoa(len(args)))
	}

	for i := 0; i < len(args); i++ {
		if args[i] == "-h" || args[i] == "--help" {
			printUsage()
		} else if args[i] == "-q" || args[i] == "--quiet" {
			flags.flagQ = true
		} else if args[i] == "-f" || args[i] == "--forest" {
			flags.flagF = true
		} else if args[i] == "-d" || args[i] == "--depth" {
			i++
			flags.depth = parseDepth(i, args)
		} else {
			flags = parseDataPath(args[i], flags)
		}
	}
	return flags
}
