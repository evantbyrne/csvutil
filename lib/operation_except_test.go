package lib

import (
	"testing"
)

func TestOperationExcept(t *testing.T) {
	source := &Source{
		Rows: [][]string{
			{"id", "name"},
			{"1", "Oof"},
			{"2", "Bar"},
			{"3", "Zab"},
			{"4", "Foobar"},
			{"5", "Foo"},
		},
	}
	operation := &OperationExcept{}
	if err, _ := operation.Construct(source, []string{"--except"}); err == nil {
		t.Fatal("Expected --except to fail without a comma-separated list of column names.")
	}
	if err, _ := operation.Construct(source, []string{"--except", "*"}); err != nil {
		t.Fatalf("Unexpected --except failure: %s", err)
	}
	if err := operation.Run(source); err == nil {
		t.Fatal("Expected --except to fail when run on first source.")
	}

	expected := [][]string{
		{"id", "name"},
		{"1", "Foo"},
		{"3", "Baz"},
	}
	source.Previous = &Source{
		Rows: [][]string{
			{"id", "name"},
			{"1", "Foo"},
			{"2", "Bar"},
			{"3", "Baz"},
			{"4", "Foobar"},
		},
	}
	if err := operation.Run(source); err != nil {
		t.Fatalf("Unexpected --except failure: %s", err)
	}
	if !rowsEqual(source.Rows, expected) {
		t.Fatalf("Unexpected --except results: %v", source.Rows)
	}

	source = &Source{
		Previous: &Source{
			Rows: [][]string{
				{"id", "name"},
				{"1", "Foo"},
				{"2", "Bar"},
				{"3", "Baz"},
				{"4", "Foobar"},
			},
		},
		Rows: [][]string{
			{"id", "name"},
			{"1", "Oof"},
			{"2", "Bar"},
			{"3", "Zab"},
			{"4", "Foobar"},
			{"5", "Foo"},
		},
	}
	expected = [][]string{
		{"id", "name"},
	}
	if err, _ := operation.Construct(source, []string{"--except", "id"}); err != nil {
		t.Fatalf("Unexpected --except failure: %s", err)
	}
	if err := operation.Run(source); err != nil {
		t.Fatalf("Unexpected --except failure: %s", err)
	}
	if !rowsEqual(source.Rows, expected) {
		t.Fatalf("Unexpected --except results: %v", source.Rows)
	}

	source = &Source{
		Previous: &Source{
			Rows: [][]string{
				{"id", "name"},
				{"1", "Oof"},
				{"2", "Bar"},
				{"3", "Zab"},
				{"4", "Foobar"},
				{"5", "Foo"},
			},
		},
		Rows: [][]string{
			{"id", "name"},
			{"1", "Foo"},
			{"2", "Bar"},
			{"3", "Baz"},
			{"4", "Foobar"},
		},
	}
	expected = [][]string{
		{"id", "name"},
		{"1", "Oof"},
		{"3", "Zab"},
	}
	if err, _ := operation.Construct(source, []string{"--except", "name"}); err != nil {
		t.Fatalf("Unexpected --except failure: %s", err)
	}
	if err := operation.Run(source); err != nil {
		t.Fatalf("Unexpected --except failure: %s", err)
	}
	if !rowsEqual(source.Rows, expected) {
		t.Fatalf("Unexpected --except results: %v", source.Rows)
	}
}
