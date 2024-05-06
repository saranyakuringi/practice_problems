package exc17

import (
	"testing"
	"time"
)

func Test_character_Counterpart(t *testing.T) {
	var want1 rune
	var size2 int
	want1, size2 = 65, 1
	got1, size1 := character_Counterpart("A")
	if got1 != want1 {
		t.Errorf("character_Counterpart()=%d %d,want=%d %d", got1, size1, want1, size2)
	}
}

func Test_calculateTimeDifference(t *testing.T) {
	cityA := "America/New_York"
	timestampA := "2024-01-29 12:34:56"
	cityB := "America/Los_Angeles"

	want1Layout := "2006-01-02 15:04:05"
	want1, _ := time.Parse(want1Layout, "2024-01-29 12:34:56")

	got1, got2 := calculateTimeDifference(cityA, timestampA, cityB)

	if !got1.Equal(want1) {
		t.Errorf("Time in %s: got=%s, want=%s", cityA, got1.Format(want1Layout), want1.Format(want1Layout))
	}

	want2 := got1.Add(-time.Hour * 3)

	// Remove want2Layout from the comparison
	if !got2.Equal(want2) {
		t.Errorf("Time in %s: got=%s, want=%s", cityB, got2.Format(want1Layout), want1.Format(want1Layout))
	}
}
