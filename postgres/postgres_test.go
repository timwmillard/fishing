package postgres

import (
	"database/sql"
	"testing"
)

func TestNullIntP(t *testing.T) {
	tests := []struct {
		arg      sql.NullInt32
		expected *int
	}{
		{
			arg:      sql.NullInt32{Int32: 3, Valid: true},
			expected: intP(3),
		},
		{
			arg:      sql.NullInt32{Int32: 0, Valid: true},
			expected: intP(0),
		},
		{
			arg:      sql.NullInt32{Int32: 0, Valid: false},
			expected: intP(0),
		},
	}

	for _, test := range tests {
		got := nullIntP(test.arg)
		if *got != *test.expected {
			t.Errorf("nullIntP got %v but expected %v", got, test.expected)
		}
	}
}

func intP(i int) *int {
	r := new(int)
	*r = i
	return r
}

func TestNullInt(t *testing.T) {
	tests := []struct {
		arg      sql.NullInt32
		expected int
	}{
		{
			arg:      sql.NullInt32{Int32: 3, Valid: true},
			expected: 3,
		},
		{
			arg:      sql.NullInt32{Int32: 0, Valid: true},
			expected: 0,
		},
		{
			arg:      sql.NullInt32{Int32: 0, Valid: false},
			expected: 0,
		},
	}

	for _, test := range tests {
		got := nullInt(test.arg)
		if got != test.expected {
			t.Errorf("nullIntP got %v but expected %v", got, test.expected)
		}
	}
}

func TestNullStringP(t *testing.T) {
	tests := []struct {
		arg      sql.NullString
		expected *string
	}{
		{
			arg:      sql.NullString{String: "test", Valid: true},
			expected: stringP("test"),
		},
		{
			arg:      sql.NullString{String: "", Valid: true},
			expected: stringP(""),
		},
		{
			arg:      sql.NullString{String: "", Valid: false},
			expected: stringP(""),
		},
	}

	for _, test := range tests {
		got := nullStringP(test.arg)
		if *got != *test.expected {
			t.Errorf("nullIntP got %v but expected %v", got, test.expected)
		}
	}
}

func stringP(s string) *string {
	r := new(string)
	*r = s
	return r
}

func TestNullString(t *testing.T) {
	tests := []struct {
		arg      sql.NullString
		expected string
	}{
		{
			arg:      sql.NullString{String: "test", Valid: true},
			expected: "test",
		},
		{
			arg:      sql.NullString{String: "", Valid: true},
			expected: "",
		},
		{
			arg:      sql.NullString{String: "", Valid: false},
			expected: "",
		},
	}

	for _, test := range tests {
		got := nullString(test.arg)
		if got != test.expected {
			t.Errorf("nullIntP got %v but expected %v", got, test.expected)
		}
	}
}
