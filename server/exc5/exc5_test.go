package exc5

import (
	"testing"
)

func TestExc_05(t *testing.T) {
	want := "HELLO WORLD!"

	// Capture the result of Exc5_ReadString
	input := Exc5_ReadString()

	// Pass the result to PrintString
	got := PrintString(input)

	if got != want {
		t.Errorf("PrintString(Exc5_ReadString()) = %v, want %v", got, want)
	}
}
