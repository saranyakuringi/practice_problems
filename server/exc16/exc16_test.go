package exc16

import "testing"

func Test_Add(t *testing.T) {
	want2 := 10
	b := 9
	got2 := Add(b)

	if got2 != want2 {
		t.Errorf("Add()=%d,want1=%d", got2, want2)
	}
}

func Test_Power(t *testing.T) {
	want := 330
	voltage, current := 110, 3
	got := Power(voltage, current)
	if got != want {
		t.Errorf("Power()=%d,want=%d", got, want)
	}
}

func Test_Triangle_thirdEdge(t *testing.T) {
	want := 17
	side1 := 8
	side2 := 10
	got := Triangle_thirdEdge(side1, side2)
	if got != want {
		t.Errorf("Triangle_thirdEdge()=%d,want=%d", got, want)
	}
}
