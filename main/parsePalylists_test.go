package parsePlaylists

import (
	"reflect"
	"testing"
)

func TestDefineSongOrder(t *testing.T) {
	lines := [][]string{
		{"folder", "song"},
		{"folder", "song2"},
	}
	got, _ := DefineSongOrder(lines)
	expected := [][]string{
		{"folder", "0 - song"},
		{"folder", "1 - song2"},
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("expected %q but got %q", expected, got)
	}
}
