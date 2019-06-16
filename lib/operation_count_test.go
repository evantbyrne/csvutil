package lib

import (
	"testing"
)

func TestOperationCount(t *testing.T) {
	source := &Source{
		Rows: [][]string{
			{"id", "name"},
			{"1", "Foo"},
			{"2", "Bar"},
			{"3", "Baz"},
		},
	}
	operation := &OperationCount{}
	if err, _ := operation.Construct(source, []string{"--count"}); err == nil {
		t.Fatal("Expected --count to fail without column name.")
	}
	if err, _ := operation.Construct(source, []string{"--count", "total"}); err != nil {
		t.Fatalf("Unexpected --count failure: %s", err)
	}
	if err := operation.Run(source); err != nil {
		t.Fatalf("Unexpected --count failure: %s", err)
	}

	expected := [][]string{
		{"total"},
		{"3"},
	}
	if !rowsEqual(source.Rows, expected) {
		t.Fatalf("Unexpected --count results: %v", source.Rows)
	}
}
