package model

import "testing"

func TestBuildStub(t *testing.T) {
	tests := []struct {
		str string
		exp string
	}{
		{"J.K. Rowling", "jk-rowling"},
		{"Robert E. Lee", "robert-e-lee"},
	}

	for i, test := range tests {
		if stub := BuildStub(test.str); stub != test.exp {
			t.Errorf("test[%d] expected: %s; got: %s\n", i, test.exp, stub)
		}
	}
}
