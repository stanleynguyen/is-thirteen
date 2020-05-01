package is

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/stanleynguyen/is-thirteen/internal/anagram"
)

type numberMatcher struct {
	value       float64
	Roughly     roughlyMatcher
	Not         invertedMatcher
	DivisibleBy divisibilityMatcher
	SquareOf    squareMatcher
	GreaterThan greaterMatcher
	LessThan    lessMatcher
	YearOfBirth birthMatcher
}

// Number Instanitiate a matcher for number type
func Number(value float64) numberMatcher {
	return numberMatcher{
		value:       value,
		Roughly:     roughlyMatcher{value},
		Not:         invertedMatcher{value},
		DivisibleBy: divisibilityMatcher{value},
		SquareOf:    squareMatcher{value},
		GreaterThan: greaterMatcher{value},
		LessThan:    lessMatcher{value},
		YearOfBirth: birthMatcher{value, realClock{}},
	}
}

// Thirteen Check if the number is Thirteen
func (m numberMatcher) Thirteen() bool {
	return m.value == Thirteen
}

type ofWithinMatcher struct {
	Of withinMatcher
}

// Within Instanitiate a matcher for number within a range around Thirtheen
func (m numberMatcher) Within(numberRange float64) ofWithinMatcher {
	return ofWithinMatcher{
		Of: withinMatcher{
			value:       m.value,
			numberRange: numberRange,
		},
	}
}

// Plus Instanitiate a number matcher with a new value with an addition
func (m numberMatcher) Plus(number float64) numberMatcher {
	return Number(m.value + number)
}

// Minus Instanitiate a number matcher with a new value with an subtraction
func (m numberMatcher) Minus(number float64) numberMatcher {
	return Number(m.value - number)
}

// Times Instanitiate a number matcher with a new value with a product
func (m numberMatcher) Times(number float64) numberMatcher {
	return Number(m.value * number)
}

// Divides Instanitiate a number matcher with a new value with a division
func (m numberMatcher) Divides(number float64) numberMatcher {
	return Number(m.value / number)
}

type stringMatcher struct {
	value        string
	CanSpell     spellMatcher
	AnagramOf    anagramMatcher
	Backwards    backwardsMatcher
	AtomicNumber atomicMatcher
}

// String Instanitiate a matcher for string type
func String(value string) stringMatcher {
	return stringMatcher{
		value:        value,
		CanSpell:     spellMatcher{value},
		AnagramOf:    anagramMatcher{value},
		Backwards:    backwardsMatcher{value},
		AtomicNumber: atomicMatcher{value},
	}
}

// Thirteen Check whether the string is 13 in base 10 or is one of ThirteenStrings
func (m stringMatcher) Thirteen() bool {
	for _, magicStr := range ThirteenStrings {
		if magicStr == m.value {
			return true
		}
	}
	baseTenInt, err := strconv.ParseInt(m.value, 10, 64)
	if err != nil {
		return false
	}
	return baseTenInt == Thirteen
}

// Base Instanitiate a number matcher with a new value of base b
func (m stringMatcher) Base(b int) numberMatcher {
	parsedInt, err := strconv.ParseInt(m.value, b, 64)
	if err != nil {
		panic(fmt.Errorf("is.String.Base: %s", err.Error()))
	}
	return Number(float64(parsedInt))
}

type roughlyMatcher struct {
	value float64
}

// Thirteen Check if the number is roughly Thirteen
func (m roughlyMatcher) Thirteen() bool {
	return math.Round(m.value) == Thirteen
}

type invertedMatcher struct {
	value float64
}

// Thirteen Check if the number is not Thirteen
func (m invertedMatcher) Thirteen() bool {
	return m.value != Thirteen
}

type divisibilityMatcher struct {
	value float64
}

// Thirteen Check if the number is divisible by Thirteen
func (m divisibilityMatcher) Thirteen() bool {
	return math.Mod(m.value, 13) == 0
}

type squareMatcher struct {
	value float64
}

// Thirteen Check if the number is square of Thirteen
func (m squareMatcher) Thirteen() bool {
	return m.value == Thirteen*Thirteen
}

type greaterMatcher struct {
	value float64
}

// Thirteen Check if the number is greater than Thirteen
func (m greaterMatcher) Thirteen() bool {
	return m.value > Thirteen
}

type lessMatcher struct {
	value float64
}

// Thirteen Check if the number is less than Thirteen
func (m lessMatcher) Thirteen() bool {
	return m.value < Thirteen
}

type withinMatcher struct {
	value       float64
	numberRange float64
}

// Thirteen Check if the number is within a range of Thirteen
func (m withinMatcher) Thirteen() bool {
	return m.value > (Thirteen-m.numberRange) && m.value < (Thirteen+m.numberRange)
}

type birthMatcher struct {
	value float64
	clock interface {
		Now() time.Time
	}
}

type realClock struct{}

func (realClock) Now() time.Time { return time.Now() }

// Thirteen Check if person with this year of birth is Thirteen year-old
func (m birthMatcher) Thirteen() bool {
	return m.value-float64(m.clock.Now().Year()) == 13
}

type spellMatcher struct {
	value string
}

// Thirteen Check if a string has enough character to spells Thirteen
func (m spellMatcher) Thirteen() bool {
	return hasCharacters(m.value, map[string]int{
		"t": 2,
		"h": 1,
		"i": 1,
		"r": 1,
		"e": 2,
		"n": 1,
	})
}

type anagramMatcher struct {
	value string
}

// Thirteen Check if a string is an anagram of Thirteen
func (m anagramMatcher) Thirteen() bool {
	return anagram.AreAnagram(strings.ToLower(m.value), "thirteen")
}

type backwardsMatcher struct {
	value string
}

// Thirteen Check if a string is thirteen spelled backwards
func (m backwardsMatcher) Thirteen() bool {
	return reverse(strings.ToLower(m.value)) == "thirteen"
}
func reverse(in string) string {
	var sb strings.Builder
	runes := []rune(in)
	for i := len(runes) - 1; 0 <= i; i-- {
		sb.WriteRune(runes[i])
	}
	return sb.String()
}

type atomicMatcher struct {
	value string
}

// Thirteen Check if the string can spell Aluminum - the 13th chemical element
func (m atomicMatcher) Thirteen() bool {
	return hasCharacters(m.value, map[string]int{
		"a": 1,
		"l": 1,
		"u": 2,
		"m": 2,
		"i": 1,
		"n": 1,
	})
}

func hasCharacters(s string, occurrenceCount map[string]int) bool {
	for char, count := range occurrenceCount {
		if strings.Count(strings.ToLower(s), char) < count {
			return false
		}
	}

	return true
}
