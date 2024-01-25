package exc001

import (
	"testing"
)

func TestExc_01(t *testing.T) {

	want := "21,28"
	got := Exc1(20, 35)

	if got != want {
		t.Errorf("exc1()=%v,want=%v", got, want)
	}
}
