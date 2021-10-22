package testproperties

import "testing"

func TestRomanNumerals(t *testing.T) {
	cases := []struct {
		Description string
		Arabic      int
		Want        string
	}{
		{"1 gets converted to I", 1, "I"},
		{"2 gets converted to II", 2, "II"},
	}

	for _, testCase := range cases {
		t.Run(testCase.Description, func(t *testing.T) {
			got := ConvertToRoman(testCase.Arabic)

			if got != testCase.Want {
				t.Errorf("got %q, want %q", got, testCase.Want)
			}
		})
	}
}
