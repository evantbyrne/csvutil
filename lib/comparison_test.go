package lib

import (
	"testing"
)

func TestComparison(t *testing.T) {
	source := &Source{
		Rows: [][]string{
			{"id", "name"},
			{"1", "Foo"},
			{"2", "Foo"},
			{"3", "Bar"},
			{"4", "Baz"},
			{"5", "Foobar"},
			{"4", "Baz"},
		},
	}

	// Prepare match.
	comparison := &Comparison{
		Column:   "blah",
		Operator: "==",
		Values:   []string{"Bar"},
	}
	if err := comparison.PrepareMatch(source); err == nil {
		t.Fatal("Expected comparison.PrepareMatch to fail on invalid column.")
	}
	comparison = &Comparison{
		Column:   "name",
		Operator: "==",
		Values:   []string{"Bar"},
	}
	if err := comparison.PrepareMatch(source); err != nil {
		t.Fatalf("Unexpected comparison.PrepareMatch failure: %s", err)
	}

	// Match ==.
	if comparison.Match([]string{"1", "Foo"}) {
		t.Fatalf("Unexpected comparison match: %v %v", comparison, []string{"1", "Foo"})
	}
	if comparison.Match([]string{"Bar", "Baz"}) {
		t.Fatalf("Unexpected comparison match: %v %v", comparison, []string{"Bar", "Baz"})
	}
	if !comparison.Match([]string{"3", "Bar"}) {
		t.Fatalf("Expected comparison match: %v %v", comparison, []string{"3", "Bar"})
	}

	// Match !=.
	comparison = &Comparison{
		Column:   "name",
		Operator: "!=",
		Values:   []string{"Bar"},
	}
	if err := comparison.PrepareMatch(source); err != nil {
		t.Fatalf("Unexpected comparison.PrepareMatch failure: %s", err)
	}
	if !comparison.Match([]string{"1", "Foo"}) {
		t.Fatalf("Expected comparison match: %v %v", comparison, []string{"1", "Foo"})
	}
	if !comparison.Match([]string{"Bar", "Baz"}) {
		t.Fatalf("Expected comparison match: %v %v", comparison, []string{"Bar", "Baz"})
	}
	if comparison.Match([]string{"3", "Bar"}) {
		t.Fatalf("Unexpected comparison match: %v %v", comparison, []string{"3", "Bar"})
	}

	// Match IN.
	comparison = &Comparison{
		Column:   "id",
		Operator: "IN",
		Values:   []string{"2,5"},
	}
	if err := comparison.PrepareMatch(source); err != nil {
		t.Fatalf("Unexpected comparison.PrepareMatch failure: %s", err)
	}
	if comparison.Match([]string{"1", "Foo"}) {
		t.Fatalf("Unexpected comparison match: %v %v", comparison, []string{"1", "Foo"})
	}
	if comparison.Match([]string{"2,5", "Foo"}) {
		t.Fatalf("Unexpected comparison match: %v %v", comparison, []string{"2,5", "Foo"})
	}
	if !comparison.Match([]string{"2", "Foo"}) {
		t.Fatalf("Expected comparison match: %v %v", comparison, []string{"2", "Foo"})
	}
	if !comparison.Match([]string{"5", "Foo"}) {
		t.Fatalf("Expected comparison match: %v %v", comparison, []string{"2", "Foo"})
	}

	// Match NOT_IN.
	comparison = &Comparison{
		Column:   "id",
		Operator: "NOT_IN",
		Values:   []string{"2,5"},
	}
	if err := comparison.PrepareMatch(source); err != nil {
		t.Fatalf("Unexpected comparison.PrepareMatch failure: %s", err)
	}
	if !comparison.Match([]string{"1", "Foo"}) {
		t.Fatalf("Expected comparison match: %v %v", comparison, []string{"1", "Foo"})
	}
	if !comparison.Match([]string{"2,5", "Foo"}) {
		t.Fatalf("Expected comparison match: %v %v", comparison, []string{"2,5", "Foo"})
	}
	if comparison.Match([]string{"2", "Foo"}) {
		t.Fatalf("Unexpected comparison match: %v %v", comparison, []string{"2", "Foo"})
	}
	if comparison.Match([]string{"5", "Foo"}) {
		t.Fatalf("Unexpected comparison match: %v %v", comparison, []string{"2", "Foo"})
	}
}
