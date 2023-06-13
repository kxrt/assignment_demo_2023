package main

import (
	"testing"
)

func TestParseChatLessThan(t *testing.T) {
	a := "a"
	b := "b"
	expected := "a:b"

	result := parseChat(a, b)
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

func TestParseChatGreaterThan(t *testing.T) {
	a := "b"
	b := "a"
	expected := "a:b"

	result := parseChat(a, b)
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

func TestParseChatEqual(t *testing.T) {
	a := "a"
	b := "a"
	expected := "a:a"

	result := parseChat(a, b)
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

func TestParseChatPull(t *testing.T) {
	chat := "a:b"
	expected := "a:b"

	result := parseChatPull(chat)
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

func TestParseChatPullReverse(t *testing.T) {
	chat := "b:a"
	expected := "a:b"

	result := parseChatPull(chat)
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

func TestParseChatPullSameParticipant(t *testing.T) {
	chat := "a:a"
	expected := "a:a"

	result := parseChatPull(chat)
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}
