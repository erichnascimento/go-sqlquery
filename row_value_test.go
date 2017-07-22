package query

import (
	"testing"
	"time"
)

type expect struct {
	given interface{}
	want  interface{}
	valid bool
}

func TestRowValueAsString(t *testing.T) {
	expecs := []expect{
		{"foo", "foo", true},
		{nil, "", false},
	}

	for _, e := range expecs {
		rv := newRowValue(&e.given)

		if got := rv.AsString().String; got != e.want {
			t.Fatalf(`given = %#v, want %#v, got %#v`, e.given, e.want, got)
		}

		if got := rv.AsString().Valid; got != e.valid {
			t.Fatalf(`given = %#v, want valid %#v, got valid %#v`, e.given, e.want, got)
		}
	}
}

func TestRowValueAsInt64(t *testing.T) {
	expecs := []expect{
		{"5", int64(5), true},
		{"NULL", int64(0), true},
		{nil, int64(0), false},
	}
	for _, e := range expecs {
		rv := newRowValue(&e.given)

		if got := rv.AsInt64().Int64; got != e.want {
			t.Fatalf(`given = %#v, want %#v, got %#v`, e.given, e.want, got)
		}

		if got := rv.AsInt64().Valid; got != e.valid {
			t.Fatalf(`given = %#v, want valid %#v, got valid %#v`, e.given, e.want, got)
		}
	}
}

func TestRowValueAsBool(t *testing.T) {
	expecs := []expect{
		{"true", true, true},
		{"t", true, true},
		{"T", true, true},
		{"1", true, true},
		{"True", true, true},
		{"TRUE", true, true},
		{nil, false, false},
	}
	for _, e := range expecs {
		rv := newRowValue(&e.given)

		if got := rv.AsBool().Bool; got != e.want {
			t.Fatalf(`given = %#v, want %#v, got %#v`, e.given, e.want, got)
		}

		if got := rv.AsBool().Valid; got != e.valid {
			t.Fatalf(`given = %#v, want valid %#v, got valid %#v`, e.given, e.want, got)
		}
	}
}

func TestRowValueAsTime(t *testing.T) {
	dateTime, _ := time.Parse("2006-01-02 15:04:05", "2016-11-23 02:30:59")
	date, _ := time.Parse("2006-01-02", "2016-11-23")
	expecs := []expect{
		{"2016-11-23 02:30:59", dateTime, true},
		{"2016-11-23", date, true},
		{"NULL", time.Time{}, false},
	}

	for _, e := range expecs {
		rv := newRowValue(&e.given)
		if got := rv.AsTime().Time; !got.Equal(e.want.(time.Time)) {
			t.Fatalf(`given = %#v, want %#v, got %#v`, e.given, e.want, got)
		}

		if got := rv.AsTime().Valid; got != e.valid {
			t.Fatalf(`given = %#v, want valid %#v, got valid %#v`, e.given, e.want, got)
		}
	}
}
