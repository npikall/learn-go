package iteration

import "testing"

func TestRepeatSimple(t *testing.T) {
	repeated := RepeatSimple("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func TestRepeatImproved(t *testing.T) {
	repeated := RepeatImproved("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeatSimple(b *testing.B) {
	for b.Loop() {
		RepeatSimple("a")
	}
}
func BenchmarkRepeatImproved(b *testing.B) {
	for b.Loop() {
		RepeatImproved("a")
	}
}
