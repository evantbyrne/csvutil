package lib

import (
	"testing"
)

func TestOperationDistinct(t *testing.T) {
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
	operation := &OperationDistinct{}
	if err, _ := operation.Construct(source, []string{"--distinct"}); err == nil {
		t.Fatal("Expected --distinct to fail without a comma-separated list of column names")
	}

	expected := [][]string{
		[]string{"id", "name"},
		[]string{"1", "Foo"},
		[]string{"2", "Foo"},
		[]string{"3", "Bar"},
		[]string{"4", "Baz"},
		[]string{"5", "Foobar"},
	}
	if err, _ := operation.Construct(source, []string{"--distinct", "*"}); err != nil {
		t.Fatalf("Unexpected --count failure: %s", err)
	}
	if err := operation.Run(source); err != nil {
		t.Fatalf("Unexpected --distinct failure: %s", err)
	}
	if !rowsEqual(source.Rows, expected) {
		t.Fatalf("Unexpected --distinct results: %v", source.Rows)
	}
	if err, _ := operation.Construct(source, []string{"--distinct", "name,id"}); err != nil {
		t.Fatalf("Unexpected --count failure: %s", err)
	}
	if err := operation.Run(source); err != nil {
		t.Fatalf("Unexpected --distinct failure: %s", err)
	}
	if !rowsEqual(source.Rows, expected) {
		t.Fatalf("Unexpected --distinct results: %v", source.Rows)
	}

	expected = [][]string{
		[]string{"id", "name"},
		[]string{"1", "Foo"},
		[]string{"3", "Bar"},
		[]string{"4", "Baz"},
		[]string{"5", "Foobar"},
	}
	if err, _ := operation.Construct(source, []string{"--distinct", "name"}); err != nil {
		t.Fatalf("Unexpected --count failure: %s", err)
	}
	if err := operation.Run(source); err != nil {
		t.Fatalf("Unexpected --distinct failure: %s", err)
	}
	if !rowsEqual(source.Rows, expected) {
		t.Fatalf("Unexpected --distinct results: %v", source.Rows)
	}
}
