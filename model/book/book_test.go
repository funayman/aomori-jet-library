package book

import "testing"

func TestIsValidIsbn(t *testing.T) {

	tests := []struct {
		isbn  string
		valid bool
	}{
		{"9780545010221", true},
		{"aaaaaaaaaaaaa", false},
		{"", false},
		{"0545010225", true},
	}

	for i, test := range tests {
		if isValid := IsValidIsbn(test.isbn); test.valid != isValid {
			t.Errorf("test[%d] - isbn: %s; expected: %t; got: %t\n", i, test.isbn, test.valid, isValid)
		}
	}
}
