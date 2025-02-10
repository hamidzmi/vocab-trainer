package review

import (
	"reflect"
	"testing"
	"time"
)

func TestUpdateReview_NotRemembered_FirstTime(t *testing.T) {
	got := Word{
		Term:       "Hallo",
		Definition: "Hello",
		Example:    "Hallo, wie geht es dir?",
	}

	UpdateReview(&got, false)
	want := Word{
		Term:       "Hallo",
		Definition: "Hello",
		Example:    "Hallo, wie geht es dir?",
		Level:      0,
		NextReview: time.Now().Add(24 * time.Hour).Truncate(time.Second),
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestUpdateReview_NotRemembered_SecondTime(t *testing.T) {
	got := Word{
		Term:       "Hallo",
		Definition: "Hello",
		Example:    "Hallo, wie geht es dir?",
		Level:      1,
	}

	UpdateReview(&got, false)
	want := Word{
		Term:       "Hallo",
		Definition: "Hello",
		Example:    "Hallo, wie geht es dir?",
		Level:      0,
		NextReview: time.Now().Add(24 * time.Hour).Truncate(time.Second),
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestUpdateReview_Remembered_FirstTime(t *testing.T) {
	got := Word{
		Term:       "Hallo",
		Definition: "Hello",
		Example:    "Hallo, wie geht es dir?",
	}

	UpdateReview(&got, true)
	want := Word{
		Term:       "Hallo",
		Definition: "Hello",
		Example:    "Hallo, wie geht es dir?",
		Level:      1,
		NextReview: time.Now().Add(24 * time.Hour).Truncate(time.Second),
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestUpdateReview_Remembered_LastTime(t *testing.T) {
	got := Word{
		Term:       "Hallo",
		Definition: "Hello",
		Example:    "Hallo, wie geht es dir?",
		Level:      2,
	}

	UpdateReview(&got, true)
	want := Word{
		Term:       "Hallo",
		Definition: "Hello",
		Example:    "Hallo, wie geht es dir?",
		Level:      3,
		NextReview: time.Now().Add(3 * 24 * time.Hour).Truncate(time.Second),
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestUpdateReview_Remembered_MoreThanThreeTimes(t *testing.T) {
	got := Word{
		Term:       "Hallo",
		Definition: "Hello",
		Example:    "Hallo, wie geht es dir?",
		Level:      3,
	}

	UpdateReview(&got, true)
	want := Word{
		Term:       "Hallo",
		Definition: "Hello",
		Example:    "Hallo, wie geht es dir?",
		Level:      3,
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestGetWordsDueForReview(t *testing.T) {
	words := []Word{
		{
			Term:       "Hallo",
			Definition: "Hello",
			Example:    "Hallo, wie geht es dir?",
			Level:      1,
			NextReview: time.Now().AddDate(0, 0, -1).Truncate(time.Second),
		},
		{
			Term:       "Tschüss",
			Definition: "Goodbye",
			Example:    "Tschüss Beispiel",
			Level:      2,
			NextReview: time.Now().AddDate(0, 0, 1).Truncate(time.Second),
		},
	}

	got := GetWordsDueForReview(words)

	want := []Word{
		{
			Term:       "Hallo",
			Definition: "Hello",
			Example:    "Hallo, wie geht es dir?",
			Level:      1,
			NextReview: time.Now().AddDate(0, 0, -1).Truncate(time.Second),
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
