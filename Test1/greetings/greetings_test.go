package greetings

import (
	"regexp"
	"testing"
)

func TestTotoName(t *testing.T) {
	name := "Agathe"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Toto("Agathe")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Toto("Agathe") = %q, %v; want match for %#q, nil`, msg, err, want)
	}
}

func TestTotoEmpty(t *testing.T) {
	msg, err := Toto("")
	if msg != "" || err == nil {
		t.Fatalf(`Toto("") = %q, %v; want "", nil`, msg, err)
	}
}
