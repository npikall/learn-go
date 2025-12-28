package roman

import "testing"

func TestRomanNumerals(t *testing.T) {
	cases := []struct {
		name string
		got  int
		want string
	}{
		{"1 gets converted to I", 1, "I"},
		{"2 gets converted to II", 2, "II"},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := ConvertToRoman(tt.got)

			if got != tt.want {
				t.Errorf("got %q want %q", got, tt.want)
			}
		})
	}
}
