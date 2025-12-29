package clockface_test

import (
	"testing"
	"time"

	clockface "github.com/npikall/learn-go/maths"
)

func TestSecondHand(t *testing.T) {
	t.Run("Second Hand at Midnight", func(t *testing.T) {
		tm := time.Date(1492, time.January, 1, 0, 0, 0, 0, time.UTC)

		want := clockface.Point{X: 150, Y: 150 - 90}
		got := clockface.SecondHand(tm)

		assertEqual(t, got, want)
	})
	t.Run("Second Hand at 30 seconds", func(t *testing.T) {
		tm := time.Date(1492, time.January, 1, 0, 0, 30, 0, time.UTC)

		want := clockface.Point{X: 150, Y: 150 + 90}
		got := clockface.SecondHand(tm)

		assertEqual(t, got, want)
	})
}

func assertEqual(t *testing.T, got, want clockface.Point) {
	t.Helper()
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
