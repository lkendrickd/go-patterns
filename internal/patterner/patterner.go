package patterner

// This is the primary pattern executor. The implementation itself is a pattern.
// It follows the command pattern.

import (
	"errors"
	"fmt"
	"log/slog"
)

// This file will launch the requested pattern function
// based on the command line arguments or by environment variables
// even with flags passed the environment variables should win

// Patterner is the interface for the pattern functions
type Patterner interface {
	Run() error
}

// PaternOperator is the struct that holds the pattern functions
type PatternOperator struct {
	Patterns map[string]Patterner // Patterns are a map of type string to Patterner
	Types    []string             // Types are the types of patterns that can be run by string name
	Logger   *slog.Logger
}

// NewPatternOperator will return a new PatternOperator struct
func NewPatternOperator(patternTypes []string, logger *slog.Logger) *PatternOperator {
	return &PatternOperator{
		Logger:   logger,
		Patterns: make(map[string]Patterner),
		Types:    patternTypes,
	}
}

// AddPattern will add a pattern to the Patterns map
func (p *PatternOperator) AddPattern(pattern Patterner) {
	p.Patterns[pattern.(*Pattern).Pattern] = pattern
}

// RemovePattern will remove a pattern from the Patterns map
func (p *PatternOperator) RemovePattern(pattern string) {
	// check if the pattern even exists
	if !p.isPatternExist(pattern) {
		return
	}

	delete(p.Patterns, pattern)
}

// RunPattern will run the pattern function
func (p *PatternOperator) RunPattern(pattern string) error {
	if !p.isPatternExist(pattern) {
		return fmt.Errorf("pattern with the name %s does not exist", pattern)
	}
	if err := p.Patterns[pattern].Run(); err != nil {
		return err
	}

	return nil
}

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

// New will return a new pattern struct
func NewPattern(pattern string, patternFunc func() error) *Pattern {
	return &Pattern{
		Pattern:     pattern,
		PatternFunc: patternFunc,
	}
}

// isPatternExist will check if the pattern exists
func (p *PatternOperator) isPatternExist(pattern string) bool {
	// check if the pattern even exists
	if _, ok := p.Patterns[pattern]; !ok {
		return false
	}

	return true
}
