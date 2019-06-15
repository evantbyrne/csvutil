package lib

import (
	"testing"
)

func TestOperationCount(t *testing.T) {
	source := &Source{
		Rows: [][]string{
			[]string{"id", "name"},
			[]string{"1", "Foo"},
			[]string{"2", "Bar"},
			[]string{"3", "Baz"},
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
		[]string{"total"},
		[]string{"3"},
	}
	if !rowsEqual(source.Rows, expected) {
		t.Fatalf("Unexpected --count results: %v", source.Rows)
	}
}
