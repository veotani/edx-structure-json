package parser

import "testing"

func TestParseStructure(t *testing.T) {
	course, err := ParseCourse()
	if err != nil {
		t.Fatal(err)
	}

	if course.DisplayName != "Академическое русское письмо" {
		t.Errorf("Incorrect course name")
	}
}
