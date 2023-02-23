package funfacts

import (
	"reflect"
	"testing"
)

/*
*

	Mal for TestGetFunFacts funksjonen.
	Definer korrekte typer for input og want,
	og sette korrekte testverdier i slice tests.
*/
func TestGetFunFacts(t *testing.T) {
	type test struct {
		input string
		want  []string
	}

	tests := []test{
		{input: "sun", want: []string{
			"The sun is a star that is located at the center of the solar system.",
			"It is estimated to be about 4.6 billion years old.",
			"The sun makes up about 99.86% of the total mass of the solar system.",
			"The temperature at the sun's core is estimated to be about 15 million Kelvin.",
		}},
		{input: "luna", want: []string{
			"The moon is the Earth's only natural satellite.",
			"The moon is approximately 238,855 miles away from Earth.",
			"The moon is approximately 1/6th the size of the Earth.",
			"The moon is believed to have formed around 4.5 billion years ago.",
		}},
		{input: "terra", want: []string{
			"Earth is the third planet from the sun.",
			"Earth has a diameter of about 7,926 miles.",
			"Earth is the only known planet to have liquid water on its surface.",
			"Earth is estimated to be about 4.54 billion years old.",
		}},
	}

	for _, tc := range tests {
		got := GetFunFacts(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}
