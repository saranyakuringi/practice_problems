package exc2

import (
	"testing"
)

func TestExc_02(t *testing.T) {
	want1 := 0
	got1 := Exc2(-1)
	if got1 != want1 {
		t.Errorf("Exc2()=%v,want=%v", got1, want1)
	}

	want2 := 1
	got2 := Exc2(1)
	if got2 != want2 {
		t.Errorf("Exc2()=%v,want=%v", got2, want2)
	}

	want3 := 6
	got3 := Exc2(3)
	if got3 != want3 {
		t.Errorf("Exc2()=%v,want=%v", got3, want3)
	}

}
