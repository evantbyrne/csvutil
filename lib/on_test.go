package lib

import (
	"testing"
)

func TestOn(t *testing.T) {
	source := &Source{
		Previous: &Source{
			Rows: [][]string{
				{"username", "id", "group_id"},
				{"foo", "1", "10"},
				{"bar", "2", ""},
				{"baz", "3", "20"},
				{"four", "4", "20"},
				{"five", "5", ""},
			},
		},
		Rows: [][]string{
			{"name", "id"},
			{"Admin", "10"},
			{"Moderator", "20"},
		},
	}

	// Prepare match.
	on := &On{
		ColumnLeft:  "foo",
		Operator:    "==",
		ColumnRight: "id",
	}
	if err := on.PrepareMatch(source); err == nil {
		t.Fatal("Expected on.PrepareMatch to fail on invalid left column.")
	}
	on = &On{
		ColumnLeft:  "group_id",
		Operator:    "==",
		ColumnRight: "foo",
	}
	if err := on.PrepareMatch(source); err == nil {
		t.Fatal("Expected on.PrepareMatch to fail on invalid right column.")
	}
	on = &On{
		ColumnLeft:  "group_id",
		Operator:    "==",
		ColumnRight: "id",
	}
	if err := on.PrepareMatch(source); err != nil {
		t.Fatal("Unexpected on.PrepareMatch failure: ", err)
	}

	// Match ==.
	if !on.Match([]string{"foo", "1", "10"}, []string{"Admin", "10"}) {
		t.Fatalf("Expected comparison match: %v %v %v", on, []string{"foo", "1", "10"}, []string{"Admin", "10"})
	}
	if on.Match([]string{"bar", "2", ""}, []string{"Admin", "10"}) {
		t.Fatalf("Unexpected comparison match: %v %v %v", on, []string{"bar", "2", ""}, []string{"Admin", "10"})
	}
	if on.Match([]string{"baz", "3", "20"}, []string{"Admin", "10"}) {
		t.Fatalf("Unexpected comparison match: %v %v %v", on, []string{"baz", "3", "20"}, []string{"Admin", "10"})
	}
	if !on.Match([]string{"baz", "3", "20"}, []string{"Moderator", "20"}) {
		t.Fatalf("Expected comparison match: %v %v %v", on, []string{"baz", "3", "20"}, []string{"Moderator", "20"})
	}

	// Match !=.
	on = &On{
		ColumnLeft:  "group_id",
		Operator:    "!=",
		ColumnRight: "id",
	}
	if err := on.PrepareMatch(source); err != nil {
		t.Fatal("Unexpected on.PrepareMatch failure: ", err)
	}
	if on.Match([]string{"foo", "1", "10"}, []string{"Admin", "10"}) {
		t.Fatalf("Unexpected comparison match: %v %v %v", on, []string{"foo", "1", "10"}, []string{"Admin", "10"})
	}
	if !on.Match([]string{"bar", "2", ""}, []string{"Admin", "10"}) {
		t.Fatalf("Expected comparison match: %v %v %v", on, []string{"bar", "2", ""}, []string{"Admin", "10"})
	}
	if !on.Match([]string{"baz", "3", "20"}, []string{"Admin", "10"}) {
		t.Fatalf("Expected comparison match: %v %v %v", on, []string{"baz", "3", "20"}, []string{"Admin", "10"})
	}
	if on.Match([]string{"baz", "3", "20"}, []string{"Moderator", "20"}) {
		t.Fatalf("Unexpected comparison match: %v %v %v", on, []string{"baz", "3", "20"}, []string{"Moderator", "20"})
	}
}
