package is_test

import (
	"fmt"
	"testing"

	"github.com/stanleynguyen/is-thirteen"
)

func TestNumberThirteen(t *testing.T) {
	var tests = []struct {
		input float64
		want  bool
	}{
		{13, true},
		{14, false},
		{13.1, false},
		{12.8, false},
	}

	for _, tc := range tests {
		testname := fmt.Sprintf("is.Number(%v).Thirteen()", tc.input)
		t.Run(testname, func(t *testing.T) {
			actual := is.Number(tc.input).Thirteen()
			if actual != tc.want {
				t.Errorf("got %v, want %v", actual, tc.want)
			}
		})
	}
}

func TestNumberRoughlyThirteen(t *testing.T) {
	var tests = []struct {
		input float64
		want  bool
	}{
		{13, true},
		{14, false},
		{13.1, true},
		{13.6, false},
		{12.8, true},
		{12.4, false},
	}

	for _, tc := range tests {
		testname := fmt.Sprintf("is.Number(%v).Roughly.Thirteen()", tc.input)
		t.Run(testname, func(t *testing.T) {
			actual := is.Number(tc.input).Roughly.Thirteen()
			if actual != tc.want {
				t.Errorf("got %v, want %v", actual, tc.want)
			}
		})
	}
}

func TestNumberNotThirteen(t *testing.T) {
	var tests = []struct {
		input float64
		want  bool
	}{
		{13, false},
		{14, true},
		{13.1, true},
		{12.8, true},
	}

	for _, tc := range tests {
		testname := fmt.Sprintf("is.Number(%v).Not.Thirteen()", tc.input)
		t.Run(testname, func(t *testing.T) {
			actual := is.Number(tc.input).Not.Thirteen()
			if actual != tc.want {
				t.Errorf("got %v, want %v", actual, tc.want)
			}
		})
	}
}

func TestNumberDivisibleByThirteen(t *testing.T) {
	var tests = []struct {
		input float64
		want  bool
	}{
		{13, true},
		{26, true},
		{39, true},
		{0, true},
		{-13, true},
		{14, false},
		{58, false},
	}

	for _, tc := range tests {
		testname := fmt.Sprintf("is.Number(%v).DivisibleBy.Thirteen()", tc.input)
		t.Run(testname, func(t *testing.T) {
			actual := is.Number(tc.input).DivisibleBy.Thirteen()
			if actual != tc.want {
				t.Errorf("got %v, want %v", actual, tc.want)
			}
		})
	}
}

func TestNumberSquareOfThirteen(t *testing.T) {
	var tests = []struct {
		input float64
		want  bool
	}{
		{13, false},
		{169, true},
		{121, false},
		{144, false},
	}

	for _, tc := range tests {
		testname := fmt.Sprintf("is.Number(%v).SquareOf.Thirteen()", tc.input)
		t.Run(testname, func(t *testing.T) {
			actual := is.Number(tc.input).SquareOf.Thirteen()
			if actual != tc.want {
				t.Errorf("got %v, want %v", actual, tc.want)
			}
		})
	}
}

func TestNumberGreaterThanThirteen(t *testing.T) {
	var tests = []struct {
		input float64
		want  bool
	}{
		{13, false},
		{14, true},
		{13.1, true},
		{12.8, false},
	}

	for _, tc := range tests {
		testname := fmt.Sprintf("is.Number(%v).GreaterThan.Thirteen()", tc.input)
		t.Run(testname, func(t *testing.T) {
			actual := is.Number(tc.input).GreaterThan.Thirteen()
			if actual != tc.want {
				t.Errorf("got %v, want %v", actual, tc.want)
			}
		})
	}
}

func TestNumberLessThanThirteen(t *testing.T) {
	var tests = []struct {
		input float64
		want  bool
	}{
		{13, false},
		{14, false},
		{13.1, false},
		{12.8, true},
	}

	for _, tc := range tests {
		testname := fmt.Sprintf("is.Number(%v).LessThan.Thirteen()", tc.input)
		t.Run(testname, func(t *testing.T) {
			actual := is.Number(tc.input).LessThan.Thirteen()
			if actual != tc.want {
				t.Errorf("got %v, want %v", actual, tc.want)
			}
		})
	}
}

func TestNumberWithinOfThirteen(t *testing.T) {
	var tests = []struct {
		input    float64
		numRange float64
		want     bool
	}{
		{13, 1, true},
		{14, 1.2, true},
		{13.1, 0.05, false},
		{12.8, 0.1, false},
	}

	for _, tc := range tests {
		testname := fmt.Sprintf("is.Number(%v).Within(%v).Of.Thirteen()", tc.input, tc.numRange)
		t.Run(testname, func(t *testing.T) {
			actual := is.Number(tc.input).Within(tc.numRange).Of.Thirteen()
			if actual != tc.want {
				t.Errorf("got %v, want %v", actual, tc.want)
			}
		})
	}
}

func TestNumberPlusThirteen(t *testing.T) {
	var tests = []struct {
		input float64
		add   float64
		want  bool
	}{
		{12, 1, true},
		{14, 1, false},
		{13.1, 0.9, false},
		{12.8, 0.2, true},
	}

	for _, tc := range tests {
		testname := fmt.Sprintf("is.Number(%v).Plus(%v).Thirteen()", tc.input, tc.add)
		t.Run(testname, func(t *testing.T) {
			actual := is.Number(tc.input).Plus(tc.add).Thirteen()
			if actual != tc.want {
				t.Errorf("got %v, want %v", actual, tc.want)
			}
		})
	}
}

func TestNumberMinusThirteen(t *testing.T) {
	var tests = []struct {
		input float64
		sub   float64
		want  bool
	}{
		{12, 1, false},
		{14, 1, true},
		{13.1, 0.1, true},
		{12.8, 0.2, false},
	}

	for _, tc := range tests {
		testname := fmt.Sprintf("is.Number(%v).Minus(%v).Thirteen()", tc.input, tc.sub)
		t.Run(testname, func(t *testing.T) {
			actual := is.Number(tc.input).Minus(tc.sub).Thirteen()
			if actual != tc.want {
				t.Errorf("got %v, want %v", actual, tc.want)
			}
		})
	}
}

func TestNumberTimesThirteen(t *testing.T) {
	var tests = []struct {
		input float64
		times float64
		want  bool
	}{
		{13, 1, true},
		{26, 0.5, true},
		{130, 0.1, true},
		{12.8, 2, false},
	}

	for _, tc := range tests {
		testname := fmt.Sprintf("is.Number(%v).Times(%v).Thirteen()", tc.input, tc.times)
		t.Run(testname, func(t *testing.T) {
			actual := is.Number(tc.input).Times(tc.times).Thirteen()
			if actual != tc.want {
				t.Errorf("got %v, want %v", actual, tc.want)
			}
		})
	}
}

func TestNumberDividesThirteen(t *testing.T) {
	var tests = []struct {
		input float64
		div   float64
		want  bool
	}{
		{13, 1, true},
		{26, 0.5, false},
		{130, 10, true},
		{12.8, 2, false},
	}

	for _, tc := range tests {
		testname := fmt.Sprintf("is.Number(%v).Divides(%v).Thirteen()", tc.input, tc.div)
		t.Run(testname, func(t *testing.T) {
			actual := is.Number(tc.input).Divides(tc.div).Thirteen()
			if actual != tc.want {
				t.Errorf("got %v, want %v", actual, tc.want)
			}
		})
	}
}

func TestStringThirteen(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"1️⃣3️⃣", true},
		{"13", true},
		{"13+0i", true},
		{"14", false},
		{"fourteen", false},
	}

	for _, tc := range tests {
		testname := fmt.Sprintf("is.String(%v).Thirteen()", tc.input)
		t.Run(testname, func(t *testing.T) {
			actual := is.String(tc.input).Thirteen()
			if actual != tc.want {
				t.Errorf("got %v, want %v", actual, tc.want)
			}
		})
	}
}

func TestStringBaseThirteen(t *testing.T) {
	var tests = []struct {
		input string
		base  int
		want  bool
	}{
		{"D", 16, true},
		{"13", 10, true},
		{"1101", 2, true},
		{"14", 8, false},
		{"1001", 2, false},
	}

	for _, tc := range tests {
		testname := fmt.Sprintf("is.String(%v).Base(%v).Thirteen()", tc.input, tc.base)
		t.Run(testname, func(t *testing.T) {
			actual := is.String(tc.input).Base(tc.base).Thirteen()
			if actual != tc.want {
				t.Errorf("got %v, want %v", actual, tc.want)
			}
		})
	}
}

func TestStringBasePanic(t *testing.T) {
	defer func() {
		r := recover()
		if r == nil {
			t.Error("got nil, want error")
		}
	}()

	is.String("123").Base(2).Thirteen()
}

func TestStringCanSpellThirteen(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"thirteen", true},
		{"tHirTeEn", true},
		{"fidEsttnnshenr", true},
		{"fourteen", false},
		{"thirteem", false},
	}

	for _, tc := range tests {
		testname := fmt.Sprintf("is.String(%v).CanSpell.Thirteen()", tc.input)
		t.Run(testname, func(t *testing.T) {
			actual := is.String(tc.input).CanSpell.Thirteen()
			if actual != tc.want {
				t.Errorf("got %v, want %v", actual, tc.want)
			}
		})
	}
}

func TestStringAnagramOfThirteen(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"thirteen", true},
		{"tHirTeEn", true},
		{"nethtire", true},
		{"fourteen", false},
		{"thirteem", false},
	}

	for _, tc := range tests {
		testname := fmt.Sprintf("is.String(%v).AnagramOf.Thirteen()", tc.input)
		t.Run(testname, func(t *testing.T) {
			actual := is.String(tc.input).AnagramOf.Thirteen()
			if actual != tc.want {
				t.Errorf("got %v, want %v", actual, tc.want)
			}
		})
	}
}

func TestStringBackwardsThirteen(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"neetriht", true},
		{"nEeTriHt", true},
		{"thirteen", false},
		{"tHirTeEn", false},
		{"nethtire", false},
	}

	for _, tc := range tests {
		testname := fmt.Sprintf("is.String(%v).Backwards.Thirteen()", tc.input)
		t.Run(testname, func(t *testing.T) {
			actual := is.String(tc.input).Backwards.Thirteen()
			if actual != tc.want {
				t.Errorf("got %v, want %v", actual, tc.want)
			}
		})
	}
}

func TestStringAtomicNumberThirteen(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"aluminum", true},
		{"aLumInUm", true},
		{"fndUsuauMimslear", true},
		{"calcium", false},
		{"uranium", false},
	}

	for _, tc := range tests {
		testname := fmt.Sprintf("is.String(%v).AtomicNumber.Thirteen()", tc.input)
		t.Run(testname, func(t *testing.T) {
			actual := is.String(tc.input).AtomicNumber.Thirteen()
			if actual != tc.want {
				t.Errorf("got %v, want %v", actual, tc.want)
			}
		})
	}
}
