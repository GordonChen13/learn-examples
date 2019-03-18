package v1

import "testing"

func TestSearch(t *testing.T)  {
	t.Run("found", func(t *testing.T) {
		dictionary := &dictionary{"test": "this is just a test"}

		want := "this is just a test"
		assertKeyMatch(t, dictionary, want, "test")
	})

	t.Run("not found", func(t *testing.T) {
		dictionary := &dictionary{"test": "this is just a test"}

		want := ""
		assertKeyMatch(t, dictionary, want, "hello")
	})
}

func assertKeyMatch(t *testing.T, dictionary *dictionary, want, key string)  {
	t.Helper()

	got := dictionary.Search(key)

	if got != want {
		t.Errorf("got '%s' want '%s' given, '%s'", got, want, "test")
	}
}
