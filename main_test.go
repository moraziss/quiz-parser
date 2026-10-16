package main

import "testing"

func TestParseLines(t *testing.T) {
	got, err := parseLines([][]string{
		{"5+5", " 10 "},
		{"1+1", "2"},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got) != 2 {
		t.Fatalf("got %d problems, want 2", len(got))
	}
	if got[0].q != "5+5" || got[0].a != "10" {
		t.Errorf("first problem = %+v, want q=5+5 a=10 (answer trimmed)", got[0])
	}
	if got[1].q != "1+1" || got[1].a != "2" {
		t.Errorf("second problem = %+v, want q=1+1 a=2", got[1])
	}
}

func TestParseLinesEmpty(t *testing.T) {
	got, err := parseLines(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got) != 0 {
		t.Errorf("got %d problems, want 0", len(got))
	}
}

func TestParseLinesRejectsRecordWithoutAnswer(t *testing.T) {
	_, err := parseLines([][]string{{"5+5"}})
	if err == nil {
		t.Fatal("expected an error for a record without an answer, got nil")
	}
}
