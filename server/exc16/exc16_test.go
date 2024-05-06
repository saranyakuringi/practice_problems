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

func Test_Min_to_Sec_converter(t *testing.T) {
	want := 300
	got := Min_to_Sec_converter(5)
	if got != want {
		t.Errorf("Min_to_Sec_converter()=%d,want=%d", got, want)
	}
}

func Test_Perimeter_of_Rectange(t *testing.T) {
	want := 26
	got := Perimeter_of_Rectange(6, 7)
	if got != want {
		t.Errorf("Perimeter_of_Rectange()=%d,want=%d", got, want)
	}
}

func Test_Remainder(t *testing.T) {
	want := 0
	got := Remainder(5, 5)
	if got != want {
		t.Errorf("Remainder()=%d,want=%d", got, want)
	}
}

func Test_Area_of_Triangle(t *testing.T) {
	want := float32(14)
	got := Area_of_the_triange(7, 4)
	if got != want {
		t.Errorf("Area_of_triangle()=%f,want=%f", got, want)
	}

}
