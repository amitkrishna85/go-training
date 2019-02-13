package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

var numeronymInputData = []struct {
	n        []string // input
	expected []string // expected result
}{
	{[]string{"accessibility", "Kubernetes", "abc"}, []string{"a11y", "K8s", "abc"}},
	{[]string{"a", "b", "abc", "abcd"}, []string{"a", "b", "abc", "a2d"}},
}

func TestNumeronymOutput(t *testing.T) {
	r := require.New(t)
	for _, tt := range numeronymInputData {
		actual := numeronyms(tt.n...)
		r.Equalf(tt.expected, actual, "Unexpected output in main()")
	}
}

func TestLettersMainOutput(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	// When
	main()

	// Then
	expected := "[a11y K8s abc]"
	actual := buf.String()
	r.Equalf(expected, actual, "Unexpected output in main()")
}
