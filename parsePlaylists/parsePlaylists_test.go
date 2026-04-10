package parsePlaylists

import (
	"reflect"
	"testing"
)

func TestDefineSongOrder(t *testing.T) {
	t.Run("Adds the correct index", func(t *testing.T) {
		lines := []string{
			"folder/song",
			"folder/song2",
		}
		got, _ := DefineTrackOrder(lines)
		expected := []string{
			"0 - song",
			"1 - song2",
		}
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("expected %q but got %q", expected, got)
		}
	})

	t.Run("Adds the correct index with 2 digits", func(t *testing.T) {
		lines := []string{
			"folder/song",
			"folder/song",
			"folder/song",
			"folder/song",
			"folder/song",
			"folder/song",
			"folder/song",
			"folder/song",
			"folder/song",
			"folder/song",
			"folder/song",
			"folder/song",
			"folder/song",
			"folder/song",
			"folder/song",
			"folder/song",
			"folder/song",
			"folder/song",
			"folder/song",
			"folder/song",
			"folder/song",
			"folder/song",
		}
		got, _ := DefineTrackOrder(lines)
		expected := []string{
			"00 - song",
			"01 - song",
			"02 - song",
			"03 - song",
			"04 - song",
			"05 - song",
			"06 - song",
			"07 - song",
			"08 - song",
			"09 - song",
			"10 - song",
			"11 - song",
			"12 - song",
			"13 - song",
			"14 - song",
			"15 - song",
			"16 - song",
			"17 - song",
			"18 - song",
			"19 - song",
			"20 - song",
			"21 - song",
		}
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("expected %q but got %q", expected, got)
		}
	})

	t.Run("Deals with empty input", func(t *testing.T) {

		lines := []string{}
		_, err := DefineTrackOrder(lines)
		// expected := [][]string{
		// 	{"folder", "0 - song"},
		// 	{"folder", "1 - song2"},
		// }
		if err != nil {
			t.Error("Didn't expect an error.")
		}
	})
}

func TestParsePlayLists(t *testing.T) {
	t.Run("Parse a valid playlist", func(t *testing.T) {
		path := "./fixtures/test.m3u"
		got, _ := ParsePlaylist(path)
		expected := Output{
			Source: []string{
				"artist1/album1/01 track.flac",
				"artist1/album1/02 track.flac",
				"artist1/album1/03 track.flac",
			},
			Target: []string{"0 - 01 track.flac",
				"1 - 02 track.flac",
				"2 - 03 track.flac",
			},
		}
		assertCorrect(t, got, expected)
	})

	t.Run("Panic with wrong path", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()
		path := "./fixtures/wrong_test.m3u"
		ParsePlaylist(path)
	})
	t.Run("Panic with empty playlist", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()
		path := "./fixtures/empty.m3u"
		ParsePlaylist(path)
	})
}
func assertCorrect(t testing.TB, got, want Output) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("expected %q but got %q", want, got)
	}
}
