package forest

import (
	"os"
	"strconv"
	"time"
)

// flags contains all flags & arguments
type flags struct {
	dataPath    string // data.csv
	dataPathSet bool   // default?
	depth       int    // tree depth
	size        int    // number of trees
	seed        int64  // randomize seed
	flagS       bool   // seed set?
	flagF       bool   // print forest
	flagQ       bool   // quiet
}

// defaultConfig initializes default values
func defaultConfig() flags {
	flags := flags{}
	flags.dataPath = "data.csv"
	flags.seed = time.Now().UnixNano()
	flags.depth = 5 // best number ??
	flags.size = 2  // best number ??
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
func parseDepth(args []string, i int) (int, int) {
	i++
	if i >= len(args) {
		usageError("No depth integer provided after -d", "")
	}
	depth, err := strconv.Atoi(args[i])
	if err != nil {
		usageError("Bad depth: ", args[i])
	}
	if depth <= 0 || depth >= 100000 {
		usageError("depth must be between 0 & 100000, given: ", args[i])
	}
	return depth, i
}

// parseSize parses string to int, must be between 0 & 100000
func parseSize(args []string, i int) (int, int) {
	i++
	if i >= len(args) {
		usageError("No forest size integer provided after -t", "")
	}
	size, err := strconv.Atoi(args[i])
	if err != nil {
		usageError("Bad forest size: ", args[i])
	}
	if size <= 0 || size >= 100000 {
		usageError("forest size must be between 0 & 100000, given: ", args[i])
	}
	return size, i
}

// parseSeed converts arg string to int
func parseSeed(args []string, i int, flags flags) (flags, int) {
	i++
	if i >= len(args) {
		usageError("No seed provided after -s", "")
	}
	seed, err := strconv.Atoi(args[i])
	if err != nil {
		usageError("Bad seed: ", args[i])
	}
	flags.seed = int64(seed)
	flags.flagS = true
	return flags, i
}

// parseArg parses & returns arguments for flags
func parseArg() flags {
	flags := defaultConfig()

	args := os.Args[1:]
	if len(args) == 0 {
		return flags
	} else if len(args) > 8 {
		usageError("Too many arguments: ", strconv.Itoa(len(args)))
	}

	for i := 0; i < len(args); i++ {
		if args[i] == "-h" || args[i] == "--help" {
			printUsage()
		} else if args[i] == "-d" || args[i] == "--depth" {
			flags.depth, i = parseDepth(args, i)
		} else if args[i] == "-t" || args[i] == "--trees" {
			flags.size, i = parseSize(args, i)
		} else if args[i] == "-s" || args[i] == "--seed" {
			flags, i = parseSeed(args, i, flags)
		} else if args[i] == "-f" || args[i] == "--forest" {
			flags.flagF = true
		} else if args[i] == "-q" || args[i] == "--quiet" {
			flags.flagQ = true
		} else {
			flags = parseDataPath(args[i], flags)
		}
	}
	return flags
}
