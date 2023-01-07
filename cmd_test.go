package tokscan

import "testing"

func TestFind(t *testing.T) {
	tests := []struct {
		Input    *Args
		Expected string
	}{
		{
			Input: &Args{
				Dirs: []string{"testdata/a"},
			},
			Expected: `{"testdata/a/hoge.tsx":["bg-shade-dark-default"]}`,
		},
	}

	for i, test := range tests {
		got, err := Find(test.Input)
		if err != nil {
			t.Error(err)
			continue
		}
		if got != test.Expected {
			t.Errorf("[%2d]: Expected %s but got %s", i, test.Expected, got)
			continue
		}
		t.Logf("[%2d]: Got %s", i, got)
	}
}
