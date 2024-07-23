package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

type UniqueOptions struct {
	countOccurrences   bool
	printRepeatedLines bool
	printUniqueLines   bool
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func isDirectory(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

func getOutput(outputDestination string) (*os.File, error) {
	var output = os.Stdout
	var outputError error = nil

	if outputDestination != "-" && !fileExists(outputDestination) && !isDirectory(outputDestination) {
		output, outputError = os.Create(outputDestination)
	}

	return output, outputError
}

func getInput(inputSource string) (*os.File, error) {
	var input = os.Stdin
	var inputError error = nil

	if inputSource != "-" && fileExists(inputSource) && !isDirectory(inputSource) {
		input, inputError = os.OpenFile(inputSource, os.O_RDONLY, 0644)
	}

	return input, inputError
}

func uniqueLines(lineMap map[string]int, orderedLineKeys []string, countOccurrences bool, out *os.File) {
	for index := range orderedLineKeys {
		if countOccurrences {
			_, _ = fmt.Fprintf(out, "%d %s\n", lineMap[orderedLineKeys[index]], orderedLineKeys[index])
		} else {
			_, _ = fmt.Fprintln(out, orderedLineKeys[index])
		}
	}
}

func repeatedLines(lineMap map[string]int, orderedLineKeys []string, countOccurrences bool, out *os.File) {
	for index := range orderedLineKeys {
		if lineMap[orderedLineKeys[index]] > 1 {
			if countOccurrences {
				_, _ = fmt.Fprintf(out, "%d %s\n", lineMap[orderedLineKeys[index]], orderedLineKeys[index])
			} else {
				_, _ = fmt.Fprintln(out, orderedLineKeys[index])
			}
		}
	}
}

func process(in *os.File, out *os.File, options UniqueOptions) {
	lineMap := make(map[string]int)
	var orderedLineKeys []string

	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		scannedLine := scanner.Text()

		if value, exists := lineMap[scannedLine]; exists {
			lineMap[scannedLine] = value + 1
		} else {
			lineMap[scannedLine] = 1
			orderedLineKeys = append(orderedLineKeys, scannedLine)
		}
	}

	if options.printUniqueLines {
		uniqueLines(lineMap, orderedLineKeys, options.countOccurrences, out)
	} else if options.printRepeatedLines {
		repeatedLines(lineMap, orderedLineKeys, options.countOccurrences, out)
	}

	if out != os.Stdout {
		_ = out.Close()
	}
}

func main() {
	countOccurrencesFlag := flag.Bool("c", false, "Count occurrences")
	countOccurrencesFlagAlias := flag.Bool("count", false, "Count occurrences")
	printRepeatedLinesFlag := flag.Bool("d", false, "Print only repeated lines")
	printRepeatedLinesFlagAlias := flag.Bool("repeated", false, "Print only repeated lines")
	printUniqueLinesFlag := flag.Bool("u", false, "Print only unique lines")
	printUniqueLinesFlagAlias := flag.Bool("unique", false, "Print only unique lines")

	flag.Usage = func() {
		_, _ = fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [OPTIONS] [inputfile] [outputfile]\n", os.Args[0])
		flag.PrintDefaults()
		_, _ = fmt.Fprintln(flag.CommandLine.Output(), "  [inputfile] can either be the path to an input file or '-' for stdin")
		_, _ = fmt.Fprintln(flag.CommandLine.Output(), "  [outputfile] can either be the path to an output file or '-' for stdout")
	}

	flag.Parse()

	if *countOccurrencesFlagAlias {
		*countOccurrencesFlag = true
	}
	if *printRepeatedLinesFlagAlias {
		*printRepeatedLinesFlag = true
	}
	if *printUniqueLinesFlagAlias {
		*printUniqueLinesFlag = true
	}

	if !*printUniqueLinesFlag && !*printRepeatedLinesFlag {
		*printUniqueLinesFlag = true
	}

	if *printUniqueLinesFlag && *printRepeatedLinesFlag {
		_, _ = fmt.Fprintln(os.Stderr, "Specify either -u OR -d")
		os.Exit(1)
	}

	remainingArgs := flag.Args()

	if len(remainingArgs) > 0 {
		var input, output = "-", "-"

		if len(remainingArgs) > 0 {
			input = remainingArgs[0]
		}
		if len(remainingArgs) > 1 {
			output = remainingArgs[1]
		}

		options := UniqueOptions{
			countOccurrences:   *countOccurrencesFlag,
			printRepeatedLines: *printRepeatedLinesFlag,
			printUniqueLines:   *printUniqueLinesFlag,
		}

		inputFile, inputFileError := getInput(input)
		outputFile, outputFileError := getOutput(output)

		if inputFileError == nil && outputFileError == nil {
			process(inputFile, outputFile, options)
		}
	} else {
		_, _ = fmt.Fprintln(os.Stderr, "No input or output specified")
		os.Exit(1)
	}

	os.Exit(0)
}
