package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"

	"github.com/lkendrickd/patterns/internal/pattern"
	"github.com/lkendrickd/patterns/internal/patterns/adapter"
)

// Developer Notes:  This uses the pattern operator to run a specified pattern
// The pattern operator is a struct that holds the patterns and the types of patterns
// that can be run.  The pattern operator is also responsible for running the pattern
// functions when requested.

// The pattern application itself runs on a version of the Command Pattern.
// The Command Pattern is a behavioral design pattern in which an object is used to
// encapsulate all information needed to perform an action or trigger an event at a
// later time. This information includes the method name, the object that owns the
// method and values for the method parameters.
// https://en.wikipedia.org/wiki/Command_pattern

var (
	// fPattern is the string flag pattern to be used to execute the pattern by name if it exists
	fPattern = flag.String("pattern", "", "pattern name to execute")
)

func main() {
	// create a new global slog.Logger - this is done for dependency injection purposes
	// and to maintain a single logger throughout the application
	logger := slog.New(
		slog.NewJSONHandler(os.Stdout, nil),
	)

	// Parse the flags
	flag.Parse()

	// check if the environment variable exists and is not empty
	// if the environment variable exists and is not empty then override the flag value
	// this app favors environment variables over flags
	if value := getEnv("PATTERN", ""); value != "" {
		// set the flag value to the environment variable value
		*fPattern = value
	}

	// Create a new PatternOperator
	patternOperator := pattern.NewPatternOperator([]string{"foo"}, logger)
	// Add a new pattern to the PatternOperator this adds a default pattern called foo
	// so the application can be called.
	patternOperator.AddPattern(pattern.NewPattern(
		"foo",
		func() error {
			fmt.Println("foo")
			return nil
		},
	))

	// Add the adapter pattern to the PatternOperator
	patternOperator.AddPattern(pattern.NewPattern(
		"adapter",
		adapterExecutor,
	))

	// Run the pattern
	if err := patternOperator.Run(*fPattern); err != nil {
		// log the error and exit with an exit code of 1 so this can be checked
		// such as in a ci/cd pipeline or a bash script evocation.
		logger.Error(err.Error())
		os.Exit(1)
	}

	// Remove the pattern
	patternOperator.RemovePattern(*fPattern)
}

/*##################################################################################
# Pattern Functions
##################################################################################*/

// adapterExecutor is the pattern function for the adapter pattern
func adapterExecutor() error {
	// Create a legacy read only API representing a legacy API
	legacyAPI := adapter.NewRecordsAPI()
	// Create a modern read/write API that represents a modern API
	// it will convert the records from the legacy API to the modern API
	// it stores the old Records in a new format called Entries
	modernAPI := adapter.NewEntriesAPI()

	// Create a new adapter to wrapper both the legacy and modern APIs
	adapter := adapter.NewAdapter(legacyAPI, modernAPI)

	// Convert the records from the legacy API to the modern API
	if err := adapter.ConvertRecords(); err != nil {
		return err
	}

	// List the entries from the modern API
	fmt.Println(adapter.ListEntries())

	return nil
}

/*##################################################################################
# Helper Functions
##################################################################################*/

// getEnv is a helper function to retrieve the environment variable for the pattern if the environment variable exists
// is not empty then this will override the flag value
func getEnv(envVar string, fallback string) string {
	if value, ok := os.LookupEnv(envVar); ok {
		if value != "" {
			return value
		}
	}
	return fallback
}
