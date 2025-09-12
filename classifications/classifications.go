// Package classifications provides an abstraction of Nordic Semiconductor's information classification levels.
package classifications

import (
	"fmt"
)

type Classification int

const (
	NrfInternalConfidential = iota
	NrfInternal
	NrfExternalConfidential
	NrfExternal
)

var classificationStrings = map[Classification]string{
	NrfInternalConfidential: "nrf-internal-confidential",
	NrfInternal:             "nrf-internal",
	NrfExternalConfidential: "nrf-external-confidential",
	NrfExternal:             "nrf-external",
}

// New creates a new Classification and returns an error if the name isn't known
func New(name string) (*Classification, error) {
	for _, classification := range All() {
		if name == classification.String() {
			return &classification, nil
		}
	}

	return nil, fmt.Errorf("unrecognised classification: %s", name)
}

func All() []Classification {
	return []Classification{
		NrfInternalConfidential,
		NrfInternal,
		NrfExternalConfidential,
		NrfExternal,
	}
}

// Defaults eturns the list of default classifications that we want to build for
func Defaults() []Classification {
	return []Classification{
		NrfInternal,
		NrfExternal,
	}
}

// String Returns the string form of the classification
func (c Classification) String() string {
	return classificationStrings[c]
}

// ToStringArray converts an []Classification to []string
func ToStringArray(classifications []Classification) []string {
	strings := make([]string, 0, len(classifications))

	for _, classification := range classifications {
		strings = append(strings, classification.String())
	}

	return strings
}

// FromStringArray converts an []string to []Classification and panics if one isn't recognised
//
// We panic so that we don't have to return an error so that the function can be used inline more easily
func FromStringArray(names []string) []Classification {
	classifications := make([]Classification, 0, len(names))

	for _, name := range names {
		classification, err := New(name)
		if err != nil {
			panic(err)
		}

		classifications = append(classifications, *classification)
	}

	return classifications
}
