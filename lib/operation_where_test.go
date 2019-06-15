package lib

import (
	"testing"
)

func TestOperationWhere(t *testing.T) {
	source := &Source{
		Rows: [][]string{
			[]string{"id", "name"},
			[]string{"1", "Foo"},
			[]string{"2", "Foo"},
			[]string{"3", "Bar"},
			[]string{"4", "Baz"},
			[]string{"5", "Foobar"},
			[]string{"4", "Baz"},
		},
	}
	expected := [][]string{
		[]string{"id", "name"},
		[]string{"1", "Foo"},
		[]string{"2", "Foo"},
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
