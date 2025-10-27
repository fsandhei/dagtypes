// Package targets provides the 'Target' enum to indicate a build target, derived from Rust target names.
package targets

import (
	"encoding/json"
	"fmt"
)

type Target int

const (
	X8664PcWindowsMsvc = iota
	Aarch64PcWindowsMsvc
	X8664UnknownLinuxGnu
	Aarch64UnknownLinuxGnu
	X8664AppleDarwin
	Aarch64AppleDarwin
)

var targetString = map[Target]string{
	X8664PcWindowsMsvc:     "x86_64-pc-windows-msvc",
	Aarch64PcWindowsMsvc:   "aarch64-pc-windows-msvc",
	X8664UnknownLinuxGnu:   "x86_64-unknown-linux-gnu",
	Aarch64UnknownLinuxGnu: "aarch64-unknown-linux-gnu",
	X8664AppleDarwin:       "x86_64-apple-darwin",
	Aarch64AppleDarwin:     "aarch64-apple-darwin",
}

var stringTarget = map[string]Target{
	"x86_64-pc-windows-msvc":    X8664PcWindowsMsvc,
	"aarch64-pc-windows-msvc":   Aarch64PcWindowsMsvc,
	"x86_64-unknown-linux-gnu":  X8664UnknownLinuxGnu,
	"aarch64-unknown-linux-gnu": Aarch64UnknownLinuxGnu,
	"x86_64-apple-darwin":       X8664AppleDarwin,
	"aarch64-apple-darwin":      Aarch64AppleDarwin,
}

// Creates a new Target and returns an error if the name isn't known
func New(name string) (Target, error) {
	target, ok := stringTarget[name]
	if !ok {
		return 0, fmt.Errorf("unknown or unsupported target: %s", name)
	}
	return target, nil
}

// Defaults returns the list of default targets that we want to build for
func Defaults() []Target {
	return []Target{
		X8664PcWindowsMsvc,
		X8664UnknownLinuxGnu,
		Aarch64UnknownLinuxGnu,
		X8664AppleDarwin,
		Aarch64AppleDarwin,
	}
}

// String returns the string form of the target
func (t Target) String() string {
	return targetString[t]
}

// ToStringArray converts an []Target to []string
func ToStringArray(targets []Target) []string {
	strings := make([]string, 0, len(targets))

	for _, target := range targets {
		strings = append(strings, target.String())
	}

	return strings
}

// FromStringArray converts an []string to []Target. An error is returned if a target is not recognized
func FromStringArray(names []string) ([]Target, error) {
	targets := make([]Target, 0, len(names))

	for _, name := range names {
		target, err := New(name)
		if err != nil {
			return nil, err
		}

		targets = append(targets, target)
	}

	return targets, nil
}

func (t *Target) UnmarshalText(text []byte) error {
	target, err := New(string(text))
	*t = target
	return err
}

// MarshalJSON implements the json.Marshaler interface
func (t Target) MarshalJSON() ([]byte, error) {
	str, ok := targetString[t]
	if !ok {
		return nil, fmt.Errorf("invalid target value: %d", t)
	}
	return json.Marshal(str)
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (t *Target) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return fmt.Errorf("target must be a string: %w", err)
	}

	target, ok := stringTarget[str]
	if !ok {
		return fmt.Errorf("unknown or unsupported target: %s", str)
	}

	*t = target
	return nil
}
