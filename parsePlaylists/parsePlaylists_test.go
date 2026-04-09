package parsePlaylists

import (
	"reflect"
	"testing"
)

func TestDefineSongOrder(t *testing.T) {
	t.Run("Adds the correct index", func(t *testing.T) {
		lines := [][]string{
			{"folder", "song"},
			{"folder", "song2"},
		}
		got, _ := DefineSongOrder(lines)
		expected := [][]string{
			{"folder", "0 - song"},
			{"folder", "1 - song2"},
		}
		assertCorrectMatrix(t, got, expected)
	})

	t.Run("Adds the correct index with 2 digits", func(t *testing.T) {
		lines := [][]string{
			{"folder", "song"},
			{"folder", "song"},
			{"folder", "song"},
			{"folder", "song"},
			{"folder", "song"},
			{"folder", "song"},
			{"folder", "song"},
			{"folder", "song"},
			{"folder", "song"},
			{"folder", "song"},
			{"folder", "song"},
			{"folder", "song"},
			{"folder", "song"},
			{"folder", "song"},
			{"folder", "song"},
			{"folder", "song"},
			{"folder", "song"},
			{"folder", "song"},
			{"folder", "song"},
			{"folder", "song"},
			{"folder", "song"},
			{"folder", "song"},
		}
		got, _ := DefineSongOrder(lines)
		expected := [][]string{
			{"folder", "00 - song"},
			{"folder", "01 - song"},
			{"folder", "02 - song"},
			{"folder", "03 - song"},
			{"folder", "04 - song"},
			{"folder", "05 - song"},
			{"folder", "06 - song"},
			{"folder", "07 - song"},
			{"folder", "08 - song"},
			{"folder", "09 - song"},
			{"folder", "10 - song"},
			{"folder", "11 - song"},
			{"folder", "12 - song"},
			{"folder", "13 - song"},
			{"folder", "14 - song"},
			{"folder", "15 - song"},
			{"folder", "16 - song"},
			{"folder", "17 - song"},
			{"folder", "18 - song"},
			{"folder", "19 - song"},
			{"folder", "20 - song"},
			{"folder", "21 - song"},
		}
		assertCorrectMatrix(t, got, expected)
	})

	t.Run("Deals with empty input", func(t *testing.T) {

		lines := [][]string{}
		_, err := DefineSongOrder(lines)
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
		expected := [][]string{
			{"artist1", "album1", "0 - 01 track.flac"},
			{"artist1", "album1", "1 - 02 track.flac"},
			{"artist1", "album1", "2 - 03 track.flac"},
		}
		assertCorrectMatrix(t, got, expected)
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
func assertCorrectMatrix(t testing.TB, got, want [][]string) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("expected %q but got %q", want, got)
	}
}
