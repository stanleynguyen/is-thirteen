package anagram_test

import (
	"testing"

	"github.com/stanleynguyen/is-thirteen/internal/anagram"
)

func TestAreAnagram(t *testing.T) {
	if anagram.AreAnagram("listen", "silent") != true {
		t.Error(`"listen", "silent"`)
	}
	if anagram.AreAnagram("test", "ttew") != false {
		t.Error(`"test", "ttew"`)
	}
	if anagram.AreAnagram("geeksforgeeks", "forgeeksgeeks") != true {
		t.Error(`"geeksforgeeks", "forgeeksgeeks"`)
	}
	if anagram.AreAnagram("triangle", "integral") != true {
		t.Error(`"triangle", "integral"`)
	}
	if anagram.AreAnagram("abd", "acb") != false {
		t.Error(`"abd", "acb"`)
	}
}
