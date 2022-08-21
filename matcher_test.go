package main

import (
	"regexp"
	"testing"
)

func TestMatch(t *testing.T) {
	tests := []struct {
		re       string
		txt      string
		expected bool
	}{
		{"", "", true},
		{"", "foo", true},
		{"foo", "", false},
		{"foo", "foo", true},
		{"foo", "A food truck", true},
		{"foo", "foes", false},
		{"^", "", true},
		{"^", "foo", true},
		{"^foo", "foo", true},
		{"^foo", "food", true},
		{"^foo", "fo", false},
		{"^$", "", true},
		{"^$", "x", false},
		{"^.$", "", false},
		{"^.$", "x", true},
		{"^.$", "xy", false},
		{"$", "", true},
		{"$", "foo", true},
		{"foo$", "foo", true},
		{"foo$", "xfoo", true},
		{"foo$", "oo", false},
		{"foo$", "food", false},
		{"foo$", "A food truck", false},
		{"ab$", "abcab", true},
		{"a*", "", true},
		{"a*", "a", true},
		{"a*", "aaaa", true},
		{"fo*d", "fd", true},
		{"fo*d", "fod", true},
		{"fo*d", "food", true},
		{"fo*d", "fooooood", true},
		{"fo*d", "A food truck", true},
		{"fo*d", "", false},
		{"fo*d", "f", false},
		{"fo*d", "fx", false},
		{"fo*d", "fox", false},
		{"foo.*bar", "foobar", true},
		{"foo.*bar", "foodbar", true},
		{"foo.*bar", "food and bar", true},
		{"foo.*bar", "The food bar.", true},
		{"foo.*bar", "", false},
		{"foo.*bar", "foo", false},
		{"foo.*bar", "bar", false},
		{"foo.*bar", "fooar", false},
		{"foo.*bar", "fobar", false},
		{"foo.*", "foo", true},
		{"foo.*", "A food truck", true},
		{"^foo.*$", "foodie", true},
		{"^foo.*$", "A food truck", false},
		{".*foo", "foo", true},
		{".*foo", "A food truck", true},
		{"^.*foo$", "A foo", true},
		{"^.*foo$", "A food truck", false},
		{".*", "foo", true},
		{".*", "A food truck", true},
		{"^.*$", "A foo", true},
		{"abc..", "abcde", true},
		{"abc..", "abcd", false},
		{"abc..", "abcdefghijklm", true},
		{"abc..", "_abcde_", true},
		{"abc..", "vwxyz", false},
		{"f.o", "fao", true},
		{"f.o", "fzo", true},
		{"f.o", "f.o", true},
		{"f.o", "A fxod truck", true},
		{"f.o", "fo", false},
		{"f.o", "fxy", false},
		{".dog", "The dog", true},
		{".dog", "_dog", true},
		{".dog", "doggy", false},
		{".", "", false},
		{".", "a", true},
		{".", "abcdef", true},
		{"^.$", "a", true},
		{"^.$", "ab", false},
		{".a", "xa", true},
		{".a", "_ya_", true},
		{".a", "xb", false},
		{"a.", "ax", true},
		{"a.", "_ay_", true},
		{"a.", "bx", false},
	}

	for i, tt := range tests {
		match := Match(tt.re, tt.txt)
		if match != tt.expected {
			t.Errorf("test=%d failed. got=%v, want=%v", i+1, match, tt.expected)
		}

		goMatch, err := regexp.MatchString(tt.re, tt.txt)
		if err != nil {
			t.Fatalf("error while compiling re: %v", err)
		}

		if goMatch != tt.expected {
			t.Errorf("test=%d using go's regexp failed. got=%v, want=%v", i, goMatch, tt.expected)
		}
	}
}

func BenchmarkMatch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Match("foo.*bar", "food and bar")
	}
}
