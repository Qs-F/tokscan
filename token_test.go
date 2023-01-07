package tokscan

import "testing"

func TestAllTokens(t *testing.T) {
	tests := []struct {
		ExpectedLen int
	}{
		{
			ExpectedLen: len(usages) * len(intentions) * len(levels) * len(states),
		},
	}

	for i, test := range tests {
		got := AllTokens()
		if len(got) != test.ExpectedLen {
			t.Errorf("[%2d]: Expected %d but got %d", i, test.ExpectedLen, len(got))
			continue
		}
		t.Logf("[%2d]: Got %d", i, len(got))
	}
}
