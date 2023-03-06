package custom

import (
	"testing"
)

func TestBool_Scan(t *testing.T) {
	var tests = []struct {
		name     string
		current  Bool
		input    string
		expected bool
	}{
		{"true - \x01", Bool(false), "\x01", true},
		{"false - \x01", Bool(true), "\x00", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.current.Scan(tt.input); err != nil {
				t.Error("Bool.Scan error:", err)
			}

			if !tt.current.Equals(tt.expected) {
				t.Errorf("Bool.Scan gave %t, expected %t", tt.current, tt.expected)
			}
		})
	}
}

func TestBool_Equals(t *testing.T) {
	var tests = []struct {
		name     string
		current  Bool
		input    bool
		expected bool
	}{
		{"Matching true", Bool(true), true, true},
		{"Matching false", Bool(false), false, true},
		{"Not matching true/false", Bool(true), false, false},
		{"Not matching false/true", Bool(false), true, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.current.Equals(tt.input); got != tt.expected {
				t.Errorf("current: %t, expected: %t, got: %t", tt.current, tt.expected, got)
			}
		})
	}
}

func TestBool_Boolean(t *testing.T) {
	var tests = []struct {
		name     string
		current  Bool
		expected bool
	}{
		{"True", Bool(true), true},
		{"False", Bool(false), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.current.Boolean(); got != tt.expected {
				t.Errorf("current: %t, expected: %t, got: %t", tt.current, tt.expected, got)
			}
		})
	}
}
