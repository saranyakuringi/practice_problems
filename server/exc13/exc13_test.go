package exc13

import (
	"testing"
)

func Test_Exc13(t *testing.T) {
	word1 := "ab"
	word2 := "pqrs"
	want := "apbqrs"
	got := MergeStrings(word1, word2)
	if got != want {
		t.Errorf("Exc13()=%s,want=%s", got, want)
	}
}
