package model

import "testing"

func TestGetInValidTypeErrorMessage(t *testing.T) {
	t.Run("Test number", func(t *testing.T) {
		m := GetInValidTypeErrorMessage("number")
		if m != "Not a valid integer." {
			t.Errorf("Expected \"Not a valid integer.\" got %s", m)
		}
	})
	t.Run("Test string", func(t *testing.T) {
		m := GetInValidTypeErrorMessage("string")
		if m != "Not a valid string." {
			t.Errorf("Expected \"Not a valid string.\" got %s", m)
		}
	})
	t.Run("Test boolean", func(t *testing.T) {
		m := GetInValidTypeErrorMessage("boolean")
		if m != "Not a valid boolean." {
			t.Errorf("Expected \"Not a valid boolean.\" got %s", m)
		}
	})
}
