package is

import (
	"fmt"
	"testing"
	"time"
)

type mockClock struct{}

func (c mockClock) Now() time.Time {
	return time.Date(2000, time.June, 0, 0, 0, 0, 0, time.UTC)
}

func TestNumberYearOfBirthThirteen(t *testing.T) {
	var tests = []struct {
		input float64
		want  bool
	}{
		{2013, true},
		{2012, false},
		{1996, false},
		{2005, false},
	}

	for _, tc := range tests {
		testname := fmt.Sprintf("is.Number(%v).YearOfBirth.Thirteen()", tc.input)
		t.Run(testname, func(t *testing.T) {
			matcher := Number(tc.input)
			matcher.YearOfBirth.clock = mockClock{}
			actual := matcher.YearOfBirth.Thirteen()
			if actual != tc.want {
				t.Errorf("got %v, want %v", actual, tc.want)
			}
		})
	}
}
