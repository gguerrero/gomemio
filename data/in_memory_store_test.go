package data

import (
	"testing"
)

func TestList(t *testing.T) {
	ds := NewInMemoryStore()
	ds.data["foo"] = "bar"

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
	ds := NewInMemoryStore()
	ds.data["foo"] = "bar"

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

func TestAdd(t *testing.T) {
	ds := NewInMemoryStore()

	got := ds.data["foo"]
	want := ""
	if got != want {
		t.Errorf("%q = %q, want %q", got, want, want)
	}

	ds.Add("foo", "bar")
	got = ds.data["foo"]
	want = "bar"
	if got != want {
		t.Errorf("%q = %q, want %q", got, want, want)
	}
}

func TestDelete(t *testing.T) {
	ds := NewInMemoryStore()
	ds.data["foo"] = "bar"

	got := ds.data["foo"]
	want := "bar"
	if got != want {
		t.Errorf("%q = %q, want %q", got, want, want)
	}

	ds.Delete("foo")
	got = ds.data["foo"]
	want = ""
	if got != want {
		t.Errorf("%q = %q, want %q", got, want, want)
	}
}
