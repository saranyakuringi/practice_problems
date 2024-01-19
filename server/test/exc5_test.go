package main

import (
	"practice/server/prog"
	"testing"
)

func TestExc_05(t *testing.T) {
	want := "HELLO WORLD!"

	// Capture the result of Exc5_ReadString
	input := prog.Exc5_ReadString()

	// Pass the result to PrintString
	got := prog.PrintString(input)

	if got != want {
		t.Errorf("PrintString(Exc5_ReadString()) = %v, want %v", got, want)
	}
}
