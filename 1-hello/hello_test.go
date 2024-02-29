package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	verifiesCorrectMessage := func(t *testing.T, result, expected string) {
		t.Helper()

		if result != expected {
			t.Errorf("result '%s', expected '%s'", result, expected)
		}
	}

	t.Run("says the default greeting to the people", func(t *testing.T) {
		result := Hello("Joao", "")

		expected := "Hello, Joao"

		verifiesCorrectMessage(t, result, expected)
	})

	t.Run("says 'Hello, world' when an empty string is passed", func(t *testing.T) {
		result := Hello("", "")

		expected := "Hello, world"

		verifiesCorrectMessage(t, result, expected)
	})

	t.Run("says hello in portuguese", func(t *testing.T) {
		result := Hello("Joao", "portuguese")

		expected := "Olá, Joao"

		verifiesCorrectMessage(t, result, expected)
	})

	t.Run("says hello in spanish", func(t *testing.T) {
		result := Hello("Joao", "spanish")

		expected := "Hola, Joao"

		verifiesCorrectMessage(t, result, expected)
	})
}
