package tokscan

import (
	"bufio"
	"errors"
	"io"
	"math"
	"strings"
	"testing"
)

func TestScan(t *testing.T) {
	tests := []struct {
		InputReader io.ReadCloser
		InputTarget string
		Expected    bool
		ExpectedErr error
	}{
		{
			InputReader: io.NopCloser(strings.NewReader("hello world")),
			InputTarget: "world",
			Expected:    true,
		},
		{
			InputReader: io.NopCloser(strings.NewReader("hello world again")),
			InputTarget: "world",
			Expected:    true,
		},
		{
			InputReader: io.NopCloser(strings.NewReader(`hello
			world`)),
			InputTarget: "world",
			Expected:    true,
		},
		{
			InputReader: io.NopCloser(strings.NewReader(`hello
			world
			again`)),
			InputTarget: "world",
			Expected:    true,
		},
		{
			InputReader: io.NopCloser(strings.NewReader(`hello w
orld`)),
			InputTarget: "world",
			Expected:    false,
		},
		{
			InputReader: io.NopCloser(strings.NewReader(`hello w orld`)),
			InputTarget: "world",
			Expected:    false,
		},
		{
			InputReader: io.NopCloser(strings.NewReader(`hello w orld`)),
			InputTarget: "world",
			Expected:    false,
		},
		{
			InputReader: io.NopCloser(strings.NewReader(`hello world`)),
			InputTarget: "",
			Expected:    true,
		},
		{
			InputReader: io.NopCloser(strings.NewReader(strings.Repeat(" ", math.MaxInt32))),
			InputTarget: "",
			Expected:    false,
			ExpectedErr: bufio.ErrTooLong,
		},
	}

	for i, test := range tests {
		got, err := Scan(test.InputReader, test.InputTarget)
		if errors.Is(err, test.ExpectedErr) {
			t.Logf("[%2d]: Got err: %v", i, err)
			continue
		}
		if err != nil {
			t.Errorf("[%2d] err: %v", i, err)
			continue
		}
		if got != test.Expected {
			t.Errorf("[%2d]: Expected %v but got %v", i, test.Expected, got)
			continue
		}
		t.Logf("[%2d]: Got %v", i, got)
	}
}
