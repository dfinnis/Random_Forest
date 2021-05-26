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
}

// defaultConfig initializes default values
func defaultConfig() flags {
	flags := flags{}
	flags.dataPath = "data.csv"
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

// parseArg parses & returns arguments for flags
func parseArg() flags {
	flags := defaultConfig()

	args := os.Args[1:]
	if len(args) == 0 {
		return flags
	} else if len(args) > 2 {
		usageError("Too many arguments: ", strconv.Itoa(len(args)))
	}

	for i := 0; i < len(args); i++ {
		if args[i] == "-h" || args[i] == "--help" {
			printUsage()
		} else if args[i] == "-q" || args[i] == "--quiet" {
			flags.flagQ = true
		} else {
			flags = parseDataPath(args[i], flags)
		}
	}
	return flags
}
