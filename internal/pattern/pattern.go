package pattern

// This is the primary pattern executor. The implementation itself is a pattern.
// It follows the command pattern.

import (
	"errors"
)

// Pattern is the struct that holds the pattern functions
type Pattern struct {
	// Pattern is the name of the pattern to run
	Pattern string
	// PatternFunc is the function to run
	PatternFunc func() error
}

// Run will run the pattern function
func (p *Pattern) Run() error {
	if p.PatternFunc == nil {
		return errors.New("pattern function is nil")
	}

	if err := p.PatternFunc(); err != nil {
		return err
	}

	return nil
}

// NewPattern will return a new pattern struct
func NewPattern(pattern string, patternFunc func() error) Pattern {
	return Pattern{
		Pattern:     pattern,
		PatternFunc: patternFunc,
	}
}

// PatternExist will check if the pattern exists
func (p *PatternOperator) PatternExist(pattern string) bool {
	// check if the pattern even exists
	if _, ok := p.Patterns[pattern]; !ok {
		return false
	}

	return true
}
