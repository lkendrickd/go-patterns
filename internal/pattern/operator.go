package pattern

import (
	"log/slog"

	"github.com/pkg/errors"
)

var (
	ErrAddPattern = errors.New("could not add as pattern name is not found")
)

// PatternOperator is the struct that holds the pattern functions
type PatternOperator struct {
	Patterns map[string]Pattern // Patterns are a map of type string to Pattern
	Types    []string           // Types are the types of patterns that can be run by string name
	Logger   *slog.Logger
}

// NewPatternOperator will return a new PatternOperator struct
func NewPatternOperator(patternTypes []string, logger *slog.Logger) *PatternOperator {
	return &PatternOperator{
		Logger:   logger,
		Patterns: make(map[string]Pattern),
		Types:    patternTypes,
	}
}

// AddPattern will add a pattern to the Patterns map
func (p *PatternOperator) AddPattern(pattern Pattern) error {
	// if the pattern name is not found then return an error
	if pattern.Pattern == "" {
		return errors.New("pattern name is not found")
	}

	// if the patter function is missing then return an error
	if pattern.PatternFunc == nil {
		return ErrAddPattern
	}

	// add pattern to the map
	p.Patterns[pattern.Pattern] = pattern

	// no error to return
	return nil

}

// RemovePattern will remove a pattern from the Patterns map
func (p *PatternOperator) RemovePattern(pattern string) error {
	if _, ok := p.Patterns[pattern]; ok {
		delete(p.Patterns, pattern)
		return nil
	}
	return errors.New("pattern does not exist")
}

// Run will run the pattern function
func (p *PatternOperator) Run(pattern string) error {
	// check the Patterns map to see if the requested pattern exists
	if _, ok := p.Patterns[pattern]; !ok {
		return errors.New("pattern does not exist")
	}

	// run the pattern function
	if err := p.Patterns[pattern].PatternFunc(); err != nil {
		return err
	}
	return nil
}
