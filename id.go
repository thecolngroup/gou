package util

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid/v2"
)

// ID is a unique identifier.
// Intended to wrap a ULID which is lexicographically sortable.
type ID string

// NewID returns a new ID with seed as time now.
func NewID() ID {
	return NewIDWithTime(time.Now())
}

// NewIDWithTime returns a new ID with the given time as a seed.
func NewIDWithTime(t time.Time) ID {
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	return ID(ulid.MustNew(ulid.Timestamp(t), entropy).String())
}
