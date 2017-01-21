package util

import (
	"testing"
)

func TestSlugify(t *testing.T) {
	s := "test->àèâ<-test"
	slug := Slugify(s)
	expected := "test-aea-test"
	if slug != expected {
		t.Fatal("Return string is not slugified as expected", expected, slug)
	}
}

func TestLowerOption(t *testing.T) {
	s := "Test->àèâ<-Test"
	slug := Slugify(s, true)
	expected := "test-aea-test"
	if slug != expected {
		t.Error("Return string is not slugified as expected", expected, slug)
	}
	slug = Slugify(s, false)
	expected = "Test-aea-Test"
	if slug != expected {
		t.Error("Return string is not slugified as expected", expected, slug)
	}
	slug = Slugify(s)
	expected = "Test-aea-Test"
	if slug != expected {
		t.Error("Return string is not slugified as expected", expected, slug)
	}
}
