package gospin_test

import (
	"fmt"

	"github.com/domdavis/gospin"
)

func ExampleNew() {
	// This is the same as gospin.Basic()
	gospin.New()

	// New spinner with custom frames
	gospin.New("-", "+", "*", "+", "-", " ")

	// Basic ascii spinner using |,/,-, and \
	gospin.Basic()

	// Spinner with a single dot that moves in a circle
	gospin.Dot()

	// Spinner with multiple dots moving in a circle
	gospin.Dots()

	// Spinner that uses the frames ".", "..", "...", ""
	gospin.Ellipses()

	// Spinner that uses scrolling ellipses.
	gospin.Scrolling()

	// Output:
}

func ExampleSpinner_Writer() {
	w := &mock{}

	// This is the same as gospin.New(gospin.Basic...)
	s := gospin.New()
	s.Writer(w)
	s.Advance()
	s.Advance()
	s.Advance()
	s.Advance()
	s.Advance()
	s.Done()

	// We output the output byte slice rather than the string, because the
	// control codes result in something that can't be represented in a comment.
	fmt.Println(w.output)

	// Output:
	// [27 91 63 50 53 108 124 27 91 68 47 27 91 68 45 27 91 68 92 27 91 68 124 27 91 68 27 91 63 50 53 104]
}

func ExampleSpinner_Porcelain() {
	s := gospin.New()
	s.Porcelain()
	s.Advance()
	s.Done()

	// Output:
}

type mock struct {
	buffer []byte
	output []byte
}

func (m *mock) Write(b []byte) (int, error) {
	m.buffer = append(m.buffer, b...)
	return len(b), nil
}

func (m *mock) Sync() error {
	m.output = append(m.output, m.buffer...)
	m.buffer = []byte{}
	return nil
}
