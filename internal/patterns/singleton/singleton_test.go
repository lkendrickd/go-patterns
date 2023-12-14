package singleton_test

import (
	"testing"

	"github.com/lkendrickd/patterns/internal/patterns/singleton"
)

func Test_Singleton(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
	}{
		{"TestSingleton"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			chanOperator := singleton.New()
			if chanOperator == nil {
				t.Error("chanOperator is nil")
			}

			id := chanOperator.ID

			chanOperator = singleton.New()
			if chanOperator.ID != id {
				t.Error("chanOperator is not a singleton as the id has changed")
			}
		})
	}
}
