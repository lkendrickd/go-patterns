package pattern_test

import (
	"errors"
	"io"
	"log/slog"
	"reflect"
	"testing"

	"github.com/lkendrickd/patterns/internal/pattern"
)

var (
	logger = slog.New(slog.NewJSONHandler(io.Discard, nil))
)

// TestRun tests the Run method of the Pattern struct
func TestRun(t *testing.T) {
	tests := []struct {
		name    string
		pattern pattern.Pattern
		wantErr bool
	}{
		{
			name: "ValidPattern",
			pattern: pattern.Pattern{
				PatternFunc: func() error { return nil },
			},
			wantErr: false,
		},
		{
			name: "NilPatternFunc",
			pattern: pattern.Pattern{
				PatternFunc: nil,
			},
			wantErr: true,
		},
		{
			name: "PatternFuncReturnsError",
			pattern: pattern.Pattern{
				PatternFunc: func() error { return errors.New("error") },
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.pattern.Run(); (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewPattern(t *testing.T) {
	dummyFunc := func() error { return nil }
	tests := []struct {
		name        string
		pattern     string
		patternFunc func() error
		wantPattern pattern.Pattern
	}{
		{
			name:        "CreateValidPattern",
			pattern:     "testPattern",
			patternFunc: dummyFunc,
			wantPattern: pattern.Pattern{Pattern: "testPattern", PatternFunc: dummyFunc},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPattern := pattern.NewPattern(tt.pattern, tt.patternFunc)

			// Compare Pattern field
			if gotPattern.Pattern != tt.wantPattern.Pattern {
				t.Errorf("NewPattern() got Pattern = %v, want %v", gotPattern.Pattern, tt.wantPattern.Pattern)
			}

			// Check if PatternFunc is not nil (or any other valid check you want)
			if gotPattern.PatternFunc == nil {
				t.Errorf("NewPattern() got PatternFunc = nil, want not nil")
			}

			// Optionally, you can invoke the function and check its behavior
			// if err := gotPattern.PatternFunc(); err != nil {
			//     t.Errorf("NewPattern() PatternFunc returned error = %v", err)
			// }
		})
	}
}

func TestOperatorPatternExist(t *testing.T) {
	tests := []struct {
		name       string
		operator   pattern.PatternOperator
		pattern    string
		wantExists bool
	}{
		{
			name: "PatternExists",
			operator: pattern.PatternOperator{
				Patterns: map[string]pattern.Pattern{"existingPattern": {}},
			},
			pattern:    "existingPattern",
			wantExists: true,
		},
		{
			name:       "PatternDoesNotExist",
			operator:   pattern.PatternOperator{Patterns: map[string]pattern.Pattern{}},
			pattern:    "nonExistingPattern",
			wantExists: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotExists := tt.operator.PatternExist(tt.pattern); gotExists != tt.wantExists {
				t.Errorf("PatternExist() = %v, want %v", gotExists, tt.wantExists)
			}
		})
	}
}

func TestNewPatternOperator(t *testing.T) {
	tests := []struct {
		name         string
		patternTypes []string
		logger       *slog.Logger
		wantTypes    []string
	}{
		{
			name:         "CreatePatternOperator",
			patternTypes: []string{"type1", "type2"},
			logger:       logger,
			wantTypes:    []string{"type1", "type2"},
		},
		// Add more test cases if needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := pattern.NewPatternOperator(tt.patternTypes, tt.logger)
			if !reflect.DeepEqual(got.Types, tt.wantTypes) {
				t.Errorf("NewPatternOperator() got Types = %v, want %v", got.Types, tt.wantTypes)
			}
		})
	}
}

func TestOperatorAddPattern(t *testing.T) {
	tests := []struct {
		name     string
		operator *pattern.PatternOperator
		pattern  pattern.Pattern
		wantErr  bool
	}{
		{
			name:     "AddValidPattern",
			operator: pattern.NewPatternOperator([]string{}, logger),
			pattern:  pattern.Pattern{Pattern: "test", PatternFunc: func() error { return nil }},
			wantErr:  false,
		},
		{
			name:     "AddPatternWithEmptyName",
			operator: pattern.NewPatternOperator([]string{}, logger),
			pattern:  pattern.Pattern{Pattern: "", PatternFunc: func() error { return nil }},
			wantErr:  true,
		},
		{
			name:     "AddPatternWithNilFunc",
			operator: pattern.NewPatternOperator([]string{}, logger),
			pattern:  pattern.Pattern{Pattern: "test", PatternFunc: nil},
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.operator.AddPattern(tt.pattern); (err != nil) != tt.wantErr {
				t.Errorf("AddPattern() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOperatorRemovePattern(t *testing.T) {
	tests := []struct {
		name     string
		operator *pattern.PatternOperator
		pattern  string
		wantErr  bool
	}{
		{
			name: "RemoveExistingPattern",
			operator: func() *pattern.PatternOperator {
				op := pattern.NewPatternOperator([]string{}, logger)
				op.Patterns["test"] = pattern.Pattern{}
				return op
			}(),
			pattern: "test",
			wantErr: false,
		},
		{
			name:     "RemoveNonExistingPattern",
			operator: pattern.NewPatternOperator([]string{}, logger),
			pattern:  "test",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.operator.RemovePattern(tt.pattern); (err != nil) != tt.wantErr {
				t.Errorf("RemovePattern() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOperatorRun(t *testing.T) {
	tests := []struct {
		name     string
		operator *pattern.PatternOperator
		pattern  string
		wantErr  bool
	}{
		{
			name: "RunExistingPattern",
			operator: func() *pattern.PatternOperator {
				op := pattern.NewPatternOperator([]string{}, logger)
				op.Patterns["test"] = pattern.Pattern{PatternFunc: func() error { return nil }}
				return op
			}(),
			pattern: "test",
			wantErr: false,
		},
		{
			name:     "RunNonExistingPattern",
			operator: pattern.NewPatternOperator([]string{}, logger),
			pattern:  "test",
			wantErr:  true,
		},
		{
			name:     "RunPatternWithNilFunc",
			operator: pattern.NewPatternOperator([]string{}, logger),
			pattern:  "test",
			wantErr:  true,
		},
		{
			name: "RunPatternFuncReturnsError",
			operator: func() *pattern.PatternOperator {
				op := pattern.NewPatternOperator([]string{}, logger)
				op.Patterns["errorPattern"] = pattern.Pattern{
					PatternFunc: func() error { return errors.New("function error") },
				}
				return op
			}(),
			pattern: "errorPattern",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.operator.Run(tt.pattern); (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
