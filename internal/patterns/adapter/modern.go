package adapter

import "fmt"

// ModernAPI the modern API interface this represents the modern API
// note that this is a different interface than the legacy API and allows writes
type ModernAPI interface {
	Entries() map[string]string
	AddEntry(key string, value string) error
}

// EntriesAPI is the struct that holds the entries amd implements the ModernAPI interface
type EntriesAPI struct {
	entries map[string]string
}

// NewEntriesAPI will return a new EntriesAPI struct
func NewEntriesAPI() *EntriesAPI {
	return &EntriesAPI{
		// initialize the entries map that is empty as it will hold the
		// converted records from the legacy API
		entries: make(map[string]string),
	}
}

// Entries will return the entries and implements the ModernAPI interface
func (e *EntriesAPI) Entries() map[string]string {
	return e.entries
}

// AddEntry will add an entry to the entries and implements the ModernAPI interface
func (e *EntriesAPI) AddEntry(key string, value string) error {
	if key == "" || value == "" {
		return fmt.Errorf("key and value must not be empty")
	}

	e.entries[key] = value

	return nil
}
