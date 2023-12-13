package adapter_test

import (
	"testing"

	"github.com/lkendrickd/patterns/internal/patterns/adapter"
)

func TestRecordsAPI(t *testing.T) {
	tests := []struct {
		name   string
		expect []string
	}{
		{"DefaultRecords", []string{"foo", "bar", "baz"}},
	}

	api := adapter.NewRecordsAPI()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := api.Records(); !equalSlice(got, tt.expect) {
				t.Errorf("Records() = %v, want %v", got, tt.expect)
			}
		})
	}
}

func TestEntriesAPI(t *testing.T) {
	tests := []struct {
		name     string
		key      string
		value    string
		wantErr  bool
		expected map[string]string
	}{
		{"AddValidEntry", "key1", "value1", false, map[string]string{"key1": "value1"}},
		{"AddEmptyKey", "", "value1", true, map[string]string{}},
		{"AddEmptyValue", "key1", "", true, map[string]string{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := adapter.NewEntriesAPI() // Create a new instance for each test
			err := api.AddEntry(tt.key, tt.value)

			if (err != nil) != tt.wantErr {
				t.Errorf("AddEntry() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !equalMap(api.Entries(), tt.expected) {
				t.Errorf("Entries() = %v, want %v", api.Entries(), tt.expected)
			}
		})
	}
}

func TestAdapter(t *testing.T) {
	legacy := adapter.NewRecordsAPI()
	modern := adapter.NewEntriesAPI()
	adapter := adapter.NewAdapter(legacy, modern)

	tests := []struct {
		name string
		test func(t *testing.T)
	}{
		{"NewAdapter", func(t *testing.T) {
			if adapter == nil {
				t.Errorf("NewAdapter() returned nil")
			}
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, tt.test)
	}
}

func TestAdapter_ConvertRecords(t *testing.T) {
	legacyAPI := adapter.NewRecordsAPI()
	modernAPI := adapter.NewEntriesAPI()
	adap := adapter.NewAdapter(legacyAPI, modernAPI)

	t.Run("ConvertRecords", func(t *testing.T) {
		// Convert the records from the legacy API to the modern API
		err := adap.ConvertRecords()
		if err != nil {
			t.Errorf("ConvertRecords() returned an error: %v", err)
		}

		// Check if the number of entries in the modern API matches the number of records in the legacy API
		if len(modernAPI.Entries()) != len(legacyAPI.Records()) {
			t.Errorf("Expected %d entries in modern API, got %d", len(legacyAPI.Records()), len(modernAPI.Entries()))
		}
	})
}

func TestAdapter_ListEntries(t *testing.T) {
	legacyAPI := adapter.NewRecordsAPI()
	modernAPI := adapter.NewEntriesAPI()
	adap := adapter.NewAdapter(legacyAPI, modernAPI)

	adap.ConvertRecords() // Convert records before listing

	entries := adap.ListEntries()
	if len(entries) != len(legacyAPI.Records()) {
		t.Errorf("ListEntries() expected %d entries, got %d", len(legacyAPI.Records()), len(entries))
	}
}

// Helper functions for comparing slices and maps.
func equalSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func equalMap(a, b map[string]string) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if val, ok := b[k]; !ok || val != v {
			return false
		}
	}
	return true
}
