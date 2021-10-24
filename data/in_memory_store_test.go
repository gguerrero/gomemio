package data_test

import (
	"testing"

	"github.com/gguerrero/gomemio/data"
)

func TestList(t *testing.T) {
	ds := data.NewInMemoryStore()
	ds.Add("foo", "bar")

	want := map[string]string{
		"foo": "bar",
	}

	for k, got := range ds.List() {
		if got != want[k] {
			t.Errorf("%q = %q, want %q", got, want[k], want[k])
		}
	}
}

func TestFind(t *testing.T) {
	ds := data.NewInMemoryStore()
	ds.Add("foo", "bar")

	got := ds.Find("foo")
	want := "bar"
	if got != want {
		t.Errorf("%q = %q, want %q", got, want, want)
	}

	got = ds.Find("baz")
	want = ""
	if got != want {
		t.Errorf("%q = %q, want %q", got, want, want)
	}
}
