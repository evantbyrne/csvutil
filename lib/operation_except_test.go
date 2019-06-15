package lib

import (
	"testing"
)

func TestOperationExcept(t *testing.T) {
	source := &Source{
		Rows: [][]string{
			[]string{"id", "name"},
			[]string{"1", "Oof"},
			[]string{"2", "Bar"},
			[]string{"3", "Zab"},
			[]string{"4", "Foobar"},
			[]string{"5", "Foo"},
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
		[]string{"id", "name"},
		[]string{"1", "Foo"},
		[]string{"3", "Baz"},
	}
	source.Previous = &Source{
		Rows: [][]string{
			[]string{"id", "name"},
			[]string{"1", "Foo"},
			[]string{"2", "Bar"},
			[]string{"3", "Baz"},
			[]string{"4", "Foobar"},
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
				[]string{"id", "name"},
				[]string{"1", "Foo"},
				[]string{"2", "Bar"},
				[]string{"3", "Baz"},
				[]string{"4", "Foobar"},
			},
		},
		Rows: [][]string{
			[]string{"id", "name"},
			[]string{"1", "Oof"},
			[]string{"2", "Bar"},
			[]string{"3", "Zab"},
			[]string{"4", "Foobar"},
			[]string{"5", "Foo"},
		},
	}
	expected = [][]string{
		[]string{"id", "name"},
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
				[]string{"id", "name"},
				[]string{"1", "Oof"},
				[]string{"2", "Bar"},
				[]string{"3", "Zab"},
				[]string{"4", "Foobar"},
				[]string{"5", "Foo"},
			},
		},
		Rows: [][]string{
			[]string{"id", "name"},
			[]string{"1", "Foo"},
			[]string{"2", "Bar"},
			[]string{"3", "Baz"},
			[]string{"4", "Foobar"},
		},
	}
	expected = [][]string{
		[]string{"id", "name"},
		[]string{"1", "Oof"},
		[]string{"3", "Zab"},
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
