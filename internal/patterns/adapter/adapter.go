package adapter

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

// APIAdapter is the interface that the adapter will implement
type APIAdapter interface {
	ConvertRecords() error
	ListEntries() map[string]string
}

// Adapter is the struct that implements the APIAdapter interface
// this adapter is one of the key points of the adapter pattern
// it wraps both the legacy and modern APIs so they can be used
// interchangeably
type Adapter struct {
	legacy LegacyAPI
	modern ModernAPI
}

// NewAdapter will return a new Adapter struct
func NewAdapter(legacy LegacyAPI, modern ModernAPI) *Adapter {
	return &Adapter{
		legacy: legacy,
		modern: modern,
	}
}

// ConvertRecords will convert the records from the legacy API to the modern API
func (a *Adapter) ConvertRecords() error {
	if len(a.legacy.Records()) == 0 {
		return errors.New("no records to convert")
	}

	for _, record := range a.legacy.Records() {
		if err := a.modern.AddEntry(uuid.NewString(), record); err != nil {
			return err
		}
	}
	return nil
}

// ListEntries will list the entries from the modern API
func (a *Adapter) ListEntries() map[string]string {
	return a.modern.Entries()
}
