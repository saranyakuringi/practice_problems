package exc14

import (
	"testing"
)

func Test_Exc14(t *testing.T) {
	str1 := "ABCABC"
	str2 := "ABC"
	want := "ABC"
	got := GcdOfStrings(str1, str2)
	if got != want {
		t.Errorf("Exc14()=%s,want=%s", got, want)
	}
}
