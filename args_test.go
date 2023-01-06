package tokscan

import (
	"errors"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		Input        []string
		Expected     *Args
		ExpectedErr  error
		ExpectedFail bool
	}{
		{
			Input: []string{"-d", ".", "-t", "bg-shade-dark-default"},
			Expected: &Args{
				Dirs:   []string{"."},
				Tokens: []string{"bg-shade-dark-default"},
			},
		},
		{
			Input: []string{"-d", ".", "-d", "../test", "-t", "bg-shade-dark-default"},
			Expected: &Args{
				Dirs:   []string{".", "../test"},
				Tokens: []string{"bg-shade-dark-default"},
			},
		},
		{
			Input: []string{"-d", ".", "-t", "bg-shade-dark-default", "-d", "."},
			Expected: &Args{
				Dirs:   []string{".", "."},
				Tokens: []string{"bg-shade-dark-default"},
			},
		},
		{
			Input: []string{"-d", "./*.go", "-t", "bg-shade-dark-default"},
			Expected: &Args{
				Dirs:   []string{"./*.go"},
				Tokens: []string{"bg-shade-dark-default"},
			},
		},
		{
			Input: []string{"-s", "."},
			Expected: &Args{
				Dirs:   nil,
				Tokens: nil,
			},
			ExpectedFail: true,
		},
	}

	for i, test := range tests {
		got, err := Parse(test.Input)
		if err != nil && (test.ExpectedFail || errors.Is(err, test.ExpectedErr)) {
			t.Logf("[%2d]: Got err: %v", i, err)
			continue
		}
		if err != nil {
			t.Errorf("[%2d] err: %v", i, err)
			continue
		}
		if !reflect.DeepEqual(got, test.Expected) {
			t.Errorf("[%2d]: Expected %#v but got %#v", i, test.Expected, got)
			continue
		}
		t.Logf("[%2d]: Got %v", i, got)
	}
}
