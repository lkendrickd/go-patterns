package adapter

// LegacyAPI the legacy API interface this represents the legacy API
type LegacyAPI interface {
	Records() []string
}

// Records is the struct that holds the records amd implements the LegacyAPI interface
type RecordsAPI struct {
	records []string
}

// NewRecordsAPI will return a new RecordsAPI struct
func NewRecordsAPI() *RecordsAPI {
	return &RecordsAPI{
		// inset some preexisting records symbolizing the legacy API data
		records: []string{"foo", "bar", "baz"},
	}
}

// Records will return the records
func (r *RecordsAPI) Records() []string {
	return r.records
}
