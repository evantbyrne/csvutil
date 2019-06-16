package lib

import (
	"testing"
)

func TestOperationWhere(t *testing.T) {
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
	expected := [][]string{
		{"id", "name"},
		{"1", "Foo"},
		{"2", "Foo"},
	}
	operation := &OperationWhere{}
	if err, _ := operation.Construct(source, []string{"--where"}); err == nil {
		t.Fatal("Expected --where to fail without column name.")
	}
	if err, _ := operation.Construct(source, []string{"--where", "name"}); err == nil {
		t.Fatal("Expected --where to fail without operator.")
	}
	if err, _ := operation.Construct(source, []string{"--where", "name", "=="}); err == nil {
		t.Fatal("Expected --where to fail without value.")
	}
	if err, _ := operation.Construct(source, []string{"--where", "name", "==", "Foo"}); err != nil {
		t.Fatalf("Unexpected --where failure: %s", err)
	}
	if err := operation.Run(source); err != nil {
		t.Fatalf("Unexpected --where failure: %s", err)
	}
	if !rowsEqual(source.Rows, expected) {
		t.Fatalf("Unexpected --where results: %v", source.Rows)
	}

	// Invalid column.
	if err, _ := operation.Construct(source, []string{"--where", "blah", "==", "Foo"}); err != nil {
		t.Fatalf("Unexpected --where failure: %s", err)
	}
	if err := operation.Run(source); err == nil {
		t.Fatal("Expected --where to fail when run with invalid column.")
	}

	// Invalid operator.
	if err, _ := operation.Construct(source, []string{"--where", "name", "!!", "Foo"}); err == nil {
		t.Fatal("Expected --where to fail when run with invalid operator.")
	}
}
