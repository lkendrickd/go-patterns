package singleton

import (
	"sync"

	"github.com/google/uuid"
)

// This is the singleton pattern. It is a creational pattern.
// It is used to ensure that only one instance of an object exists
// at a time. It is used to provide a global point of access to
// the object.

// NOTE: The full functionality for the Channel Operator will not be implemented for simplicity sake.
// the idea is to show how a singleton can be created using the sync package's sync.Once.

var (
	instance *ChannelOperator
	once     sync.Once
)

// ChannelOperator is the struct that is a singleton
// it is used as a single point of monitoring channels
type ChannelOperator struct {
	ID       string
	Channels map[string]chan interface{}
}

// New will return a new ChannelOperator struct using the go once.Do
// function to ensure that only one instance of the ChannelOperator
// struct is created.
func New() *ChannelOperator {
	// use the sync package's sync.Once to ensure that only one instance of the ChannelOperator
	// on repeated calls to the New function the ChannelOperator will not be created again
	// instead the existing ChannelOperator will be returned
	once.Do(func() {
		instance = &ChannelOperator{
			Channels: make(map[string]chan interface{}),
			ID:       uuid.NewString(), // assign a unique ID to easily show that only one instance is created
		}
	})
	return instance
}
