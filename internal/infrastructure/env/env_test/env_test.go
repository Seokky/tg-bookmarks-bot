package env_test

import (
	"testing"
	"tg-bookmarks-bot/internal/infrastructure/env"
)

func TestGet(t *testing.T) {
	t.Parallel()

	if v := env.Get("A"); v != "a" {
		t.Errorf("expected: %s, got: %s", "A", v)
	}

	if v := env.Get("B"); v != "b" {
		t.Errorf("expected: %s, got: %s", "B", v)
	}

	if v := env.Get("C"); v != "c" {
		t.Errorf("expected: %s, got: %s", "C", v)
	}

	if v := env.Get("D"); v != "" {
		t.Errorf("expected: %s, got: %s", "D", v)
	}
}

func TestGetAfterSet(t *testing.T) {
	t.Setenv("Name", "John")
	if v := env.Get("Name"); v != "John" {
		t.Errorf("expected: %s, got: %s", "John", v)
	}

	t.Setenv("Age", "50")
	if v := env.Get("Age"); v != "50" {
		t.Errorf("expected: %s, got: %s", "Age", v)
	}
}
