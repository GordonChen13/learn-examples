package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got,want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}

	t.Run("Chinese: saying hello to people", func(t *testing.T) {
		got := Hello("小明", "chinese")
		want := "你好, 小明"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Spanish: saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "spanish")
		want := "Hola, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "english")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
}