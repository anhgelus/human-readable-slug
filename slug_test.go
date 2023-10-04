package main

import (
	"strings"
	"testing"
	"time"
)

func TestGenerateSlug(t *testing.T) {
	for i := 0; i < 10; i++ {
		slug := GenerateSlug(time.Now().Unix(), 7)
		t.Log(slug)
		if !IsHumanReadable(slug) {
			t.Error("slug", slug, "is not human readable")
		}
	}
}

func TestNoNumbersSlug(t *testing.T) {
	options := CustomSlug{
		HasCapitalLetters:        true,
		HasNumbers:               false,
		CanBe2ConsecutiveLetters: false,
	}
	for i := 0; i < 10; i++ {
		slug := options.GenerateSlug(time.Now().Unix(), 7)
		t.Log(slug)
		if strings.ContainsAny(slug, numbers) {
			t.Error("slug", slug, "contains number")
		}
		if !options.IsHumanReadable(slug) {
			t.Error("slug", slug, "is not human readable")
		}
	}
}

func TestNoUpperSlug(t *testing.T) {
	options := CustomSlug{
		HasCapitalLetters:        false,
		HasNumbers:               true,
		CanBe2ConsecutiveLetters: false,
	}
	for i := 0; i < 10; i++ {
		slug := options.GenerateSlug(time.Now().Unix(), 7)
		t.Log(slug)
		if strings.ContainsAny(slug, strings.ToUpper(minLetters)) {
			t.Error("slug", slug, "contains upper case")
		}
		if !options.IsHumanReadable(slug) {
			t.Error("slug", slug, "is not human readable")
		}
	}
}

func TestNoUpperNoNumbersSlug(t *testing.T) {
	options := CustomSlug{
		HasCapitalLetters:        false,
		HasNumbers:               false,
		CanBe2ConsecutiveLetters: false,
	}
	for i := 0; i < 10; i++ {
		slug := options.GenerateSlug(time.Now().Unix(), 7)
		t.Log(slug)
		if strings.ContainsAny(slug, strings.ToUpper(minLetters)+numbers) {
			t.Error("slug", slug, "contains upper case")
		}
		if !options.IsHumanReadable(slug) {
			t.Error("slug", slug, "is not human readable")
		}
	}
}
