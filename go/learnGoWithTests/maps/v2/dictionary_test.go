package v2

import "testing"

func TestSearch(t *testing.T)  {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("hello")

		if err == nil {
			t.Fatal("expected to get an error")
		}
		want := "word not found"
		assertStrings(t, err.Error(), want)
	})
}

func TestAdd(t *testing.T)  {
	t.Run("un exists word", func(t *testing.T) {
		dictionary := &Dictionary{}

		err := dictionary.Add("hello", "hello world")
		if err != nil {
			t.Fatal("add fail! expected no error")
		}
		got, _ := dictionary.Search("hello")
		want := "hello world"

		assertStrings(t, got, want)
	})

	t.Run("exists word", func(t *testing.T) {
		dictionary := &Dictionary{"test": "this is just a test"}
		err := dictionary.Add("test", "duplicate word test")
		if err == nil {
			t.Fatal("expected duplicate error")
		}
		want := "duplicate word error"

		assertStrings(t, err.Error(), want)
	})

}

func assertStrings(t *testing.T, got, want string)  {
	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func TestUpdate(t *testing.T)  {
	t.Run("un exists word", func(t *testing.T) {
		dictionary := &Dictionary{}

		err := dictionary.Update("test", "hello world")
		if err == nil {
			t.Fatal("expected update error")
		}
		want := "update not exists word error"

		assertStrings(t, err.Error(), want)
	})

	t.Run("exists word", func(t *testing.T) {
		dictionary := &Dictionary{"test": "this is just a test"}
		err := dictionary.Update("test", "update word test")
		if err != nil {
			t.Fatal("update fail! expected no error")
		}
		got, _ := dictionary.Search("test")
		want := "update word test"

		assertStrings(t, got, want)
	})

}

func TestDelete(t *testing.T)  {
	t.Run("un exists word", func(t *testing.T) {
		dictionary := &Dictionary{}

		err := dictionary.Delete("test")
		if err == nil {
			t.Fatal("expected delete error")
		}
		want := "delete not exists word error"

		assertStrings(t, err.Error(), want)
	})

	t.Run("exists word", func(t *testing.T) {
		dictionary := &Dictionary{"test": "this is just a test"}
		err := dictionary.Delete("test")
		if err != nil {
			t.Fatal("delete fail! expected no error")
		}
		_, err = dictionary.Search("test")

		if err != NotFoundError {
			t.Error("expected not found error")
		}
	})

}

