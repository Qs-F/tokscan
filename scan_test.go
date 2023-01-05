package tokscan

import (
	"io"
	"strings"
	"testing"
)

func TestScan(t *testing.T) {
	tests := []struct {
		InputReader io.Reader
		InputTarget string
		Expected    bool
	}{
		{
			InputReader: strings.NewReader("hello world"),
			InputTarget: "world",
			Expected:    true,
		},
		{
			InputReader: strings.NewReader("hello world again"),
			InputTarget: "world",
			Expected:    true,
		},
		{
			InputReader: strings.NewReader(`hello
			world`),
			InputTarget: "world",
			Expected:    true,
		},
		{
			InputReader: strings.NewReader(`hello
			world
			again`),
			InputTarget: "world",
			Expected:    true,
		},
		{
			InputReader: strings.NewReader(`hello w
orld`),
			InputTarget: "world",
			Expected:    false,
		},
		{
			InputReader: strings.NewReader(`hello w orld`),
			InputTarget: "world",
			Expected:    false,
		},
		{
			InputReader: strings.NewReader(`hello w orld`),
			InputTarget: "world",
			Expected:    false,
		},
	}

	for i, test := range tests {
		got, err := Scan(test.InputReader, test.InputTarget)
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
