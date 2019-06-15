package lib

import (
	"testing"
)

func TestOperationConcat(t *testing.T) {
	source := &Source{
		Rows: [][]string{
			[]string{"id", "name"},
			[]string{"2", "Bar"},
		},
	}
	operation := &OperationConcat{}
	_, args := operation.Construct(source, []string{"--concat"})
	if len(args) != 0 {
		t.Fatal("Expected --concat to not have any remaining args.")
	}

	if err := operation.Run(source); err == nil {
		t.Fatal("Expected --concat to fail when used on first source.")
	}

	source.Previous = &Source{
		Rows: [][]string{
			[]string{"name", "id"},
			[]string{"Bar", "id"},
		},
	}
	if err := operation.Run(source); err == nil {
		t.Fatal("Expected --concat to fail when source columns mismatch.")
	}

	source.Previous = &Source{
		Rows: [][]string{
			[]string{"id", "name"},
			[]string{"1", "Foo"},
		},
	}
	if err := operation.Run(source); err != nil {
		t.Fatalf("Unexpected --concat failure: %s", err)
	}
	expected := [][]string{
		[]string{"id", "name"},
		[]string{"1", "Foo"},
		[]string{"2", "Bar"},
	}
	if !rowsEqual(source.Rows, expected) {
		t.Fatalf("Unexpected --concat results: %v", source.Rows)
	}
}
