package exc15

import (
	"testing"
)

func Test_Exc15(t *testing.T) {
	candies1 := []int{2, 3, 5, 1, 3}
	extraCandies1 := 3
	want := []bool{true, true, true, false, true}
	//var got []bool
	got := KidsWithCandies(candies1, extraCandies1)
	for i := range got {
		if got[i] != want[i] {
			t.Errorf("Exc15()=%t,want=%t", got[i], want[i])
		}
	}

}
