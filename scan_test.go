package tokscan

import (
	"bufio"
	"errors"
	"io"
	"math"
	"reflect"
	"strings"
	"testing"
)

func TestScan(t *testing.T) {
	tests := []struct {
		InputReader io.ReadCloser
		InputTarget []string
		Expected    []string
		ExpectedErr error
	}{
		{
			InputReader: io.NopCloser(strings.NewReader("hello world")),
			InputTarget: []string{"world"},
			Expected:    []string{"world"},
		},
		{
			InputReader: io.NopCloser(strings.NewReader("hello world again")),
			InputTarget: []string{"world"},
			Expected:    []string{"world"},
		},
		{
			InputReader: io.NopCloser(strings.NewReader(`hello
			world`)),
			InputTarget: []string{"world"},
			Expected:    []string{"world"},
		},
		{
			InputReader: io.NopCloser(strings.NewReader(`hello
			world
			again`)),
			InputTarget: []string{"world"},
			Expected:    []string{"world"},
		},
		{
			InputReader: io.NopCloser(strings.NewReader(`hello w
orld`)),
			InputTarget: []string{"world"},
			Expected:    nil,
		},
		{
			InputReader: io.NopCloser(strings.NewReader(`hello w orld`)),
			InputTarget: []string{"world"},
			Expected:    nil,
		},
		{
			InputReader: io.NopCloser(strings.NewReader(`hello w orld`)),
			InputTarget: []string{"world"},
			Expected:    nil,
		},
		{
			InputReader: io.NopCloser(strings.NewReader(`hello world`)),
			InputTarget: nil,
			Expected:    nil,
		},
		{
			InputReader: io.NopCloser(strings.NewReader(strings.Repeat(" ", math.MaxInt32))),
			InputTarget: nil,
			Expected:    nil,
			ExpectedErr: bufio.ErrTooLong,
		},
	}

	for i, test := range tests {
		got, err := Scan(test.InputReader, test.InputTarget)
		if err != nil && errors.Is(err, test.ExpectedErr) {
			t.Logf("[%2d]: Got err: %v", i, err)
			continue
		}
		if err != nil {
			t.Errorf("[%2d] err: %v", i, err)
			continue
		}
		if !reflect.DeepEqual(got, test.Expected) {
			t.Errorf("[%2d]: Expected %v but got %v", i, test.Expected, got)
			continue
		}
		t.Logf("[%2d]: Got %v", i, got)
	}
}
