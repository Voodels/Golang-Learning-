package greetings

import (
	"regexp"
	"testing"
)

// t *testing.T - testing framework passes a pointer to testing.T type to the test function
func TestHello(t *testing.T) {
	//
	name := "luffy"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("luffy")
	if err != nil {
		t.Fatalf("Hello(%q) returned error: %v", name, err)
	}
	if !want.MatchString(msg) {
		t.Fatalf("Hello(%q) = %q, want match for %q", name, msg, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	_, err := Hello("")
	if err == nil {
		t.Fatal("Expected an error for empty name, but got none")
	}
}
