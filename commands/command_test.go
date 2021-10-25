package commands

import (
	"errors"
	"testing"
)

func Test_NewCommand(t *testing.T) {
	tests := []struct {
		name         string
		commands     []string
		wantedAction action
		wantedKey    string
		wantedValue  string
		wantedErr    error
	}{
		{
			name:         "GET Uppercase",
			commands:     []string{"GET", "foo"},
			wantedAction: GET,
			wantedKey:    "foo",
			wantedValue:  "",
			wantedErr:    nil,
		},
		{
			name:         "GET Lowercase",
			commands:     []string{"get", "foo"},
			wantedAction: GET,
			wantedKey:    "foo",
			wantedValue:  "",
			wantedErr:    nil,
		},
		{
			name:         "SET Uppercase",
			commands:     []string{"SET", "foo", "bar", "baz"},
			wantedAction: SET,
			wantedKey:    "foo",
			wantedValue:  "bar baz",
			wantedErr:    nil,
		},
		{
			name:         "SET Lowercase",
			commands:     []string{"set", "foo", "bar", "baz"},
			wantedAction: SET,
			wantedKey:    "foo",
			wantedValue:  "bar baz",
			wantedErr:    nil,
		},
		{
			name:         "DEL Uppercase",
			commands:     []string{"DEL", "foo"},
			wantedAction: DEL,
			wantedKey:    "foo",
			wantedValue:  "",
			wantedErr:    nil,
		},
		{
			name:         "DEL Lowercase",
			commands:     []string{"del", "foo"},
			wantedAction: DEL,
			wantedKey:    "foo",
			wantedValue:  "",
			wantedErr:    nil,
		},
		{
			name:         "LIST Uppercase",
			commands:     []string{"LIST"},
			wantedAction: LIST,
			wantedKey:    "",
			wantedValue:  "",
			wantedErr:    nil,
		},
		{
			name:         "LIST Lowercase",
			commands:     []string{"list"},
			wantedAction: LIST,
			wantedKey:    "",
			wantedValue:  "",
			wantedErr:    nil,
		},
		{
			name:         "EXIT Uppercase",
			commands:     []string{"EXIT"},
			wantedAction: EXIT,
			wantedKey:    "",
			wantedValue:  "",
			wantedErr:    nil,
		},
		{
			name:         "EXIT Lowercase",
			commands:     []string{"exit"},
			wantedAction: EXIT,
			wantedKey:    "",
			wantedValue:  "",
			wantedErr:    nil,
		},
		{
			name:         "An unknown action",
			commands:     []string{"foobar"},
			wantedAction: 0,
			wantedKey:    "",
			wantedValue:  "",
			wantedErr:    errors.New("commands parseAction: Unknown action foobar"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd, err := NewCommand(tt.commands)
			if err != nil && err.Error() != tt.wantedErr.Error() {
				t.Errorf("%q = %q, want %q", err, tt.wantedErr, tt.wantedErr)
			}

			if cmd.action != tt.wantedAction {
				t.Error("unwanted command action")
			}

			if cmd.key != tt.wantedKey {
				t.Errorf("%q = %q, want %q", cmd.key, tt.wantedKey, tt.wantedKey)
			}

			if cmd.value != tt.wantedValue {
				t.Errorf("%q = %q, want %q", cmd.value, tt.wantedValue, tt.wantedValue)
			}
		})
	}
}

func Test_IsExit(t *testing.T) {
	commands := []string{"GET"}
	cmd, _ := NewCommand(commands)

	if cmd.IsExit() {
		t.Errorf("unwanted IsExit(), wanted false")
	}

	commands = []string{"exit"}
	cmd, _ = NewCommand(commands)

	if !cmd.IsExit() {
		t.Errorf("unwanted IsExit(), wanted true")
	}
}

func Test_ExecuteGET(t *testing.T) {
	cmd := &command{
		action: GET,
		key:    "foo",
	}

	ds.Add("foo", "bar") // This is coupling with the current package dataStore. We can define a test interface impl. in the future.

	got, err := cmd.Execute()
	want := "bar"
	if err != nil {
		t.Fatal(err)
	}

	if got != want {
		t.Errorf("%s = %q, want %q", got, want, want)
	}
}

func Test_ExecuteSET(t *testing.T) {
	cmd := &command{
		action: SET,
		key:    "foo",
		value:  "bar",
	}

	got, err := cmd.Execute()
	want := "OK"
	if err != nil {
		t.Fatal(err)
	}

	if got != want {
		t.Errorf("%s = %q, want %q", got, want, want)
	}

	got = ds.Find("foo")
	want = "bar"
	if got != want {
		t.Errorf("%s = %q, want %q", got, want, want)
	}
}

func Test_ExecuteDEL(t *testing.T) {
	cmd := &command{
		action: DEL,
		key:    "foo",
	}

	ds.Add("foo", "bar")

	got, err := cmd.Execute()
	want := "DELETED"
	if err != nil {
		t.Fatal(err)
	}

	if got != want {
		t.Errorf("%s = %q, want %q", got, want, want)
	}

	got = ds.Find("foo")
	want = ""
	if got != want {
		t.Errorf("%s = %q, want %q", got, want, want)
	}
}

func Test_ExecuteLIST(t *testing.T) {
	cmd := &command{
		action: LIST,
		key:    "foo",
	}

	ds.Add("foo", "bar")

	got, err := cmd.Execute()
	want := "map[foo:bar]"
	if err != nil {
		t.Fatal(err)
	}

	if got != want {
		t.Errorf("%s = %q, want %q", got, want, want)
	}
}

func Test_ExecuteUNKNOWN(t *testing.T) {
	cmd := &command{
		action: action(100),
		key:    "foo",
	}

	_, err := cmd.Execute()
	wantedErr := errors.New("commands excute: cannot excute action 100")
	if err == nil || err.Error() != wantedErr.Error() {
		t.Errorf("%q = %q, want %q", err, wantedErr, wantedErr)
	}
}
