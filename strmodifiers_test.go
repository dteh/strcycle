package strcycle

import "testing"

func TestStringModifiers(t *testing.T) {
	base := "a"

	if UPPER(base) != "A" {
		t.Fail()
	}

	if LOWER(base) != "a" {
		t.Fail()
	}

	if SingleEncodeLower(base) != "%61" {
		t.Fail()
	}

	if SingleEncodeUpper(base) != "%41" {
		t.Fail()
	}

	if DoubleEncodeLower(base) != "%2561" {
		t.Fail()
	}

	if DoubleEncodeUpper(base) != "%2541" {
		t.Fail()
	}
}
