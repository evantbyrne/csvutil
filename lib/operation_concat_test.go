package lib

import (
	"testing"
)

func TestOperationConcat(t *testing.T) {
	source := &Source{
		Rows: [][]string{
			{"id", "name"},
			{"2", "Bar"},
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
			{"name", "id"},
			{"Bar", "id"},
		},
	}
	if err := operation.Run(source); err == nil {
		t.Fatal("Expected --concat to fail when source columns mismatch.")
	}

	source.Previous = &Source{
		Rows: [][]string{
			{"id", "name"},
			{"1", "Foo"},
		},
	}
	if err := operation.Run(source); err != nil {
		t.Fatalf("Unexpected --concat failure: %s", err)
	}
	expected := [][]string{
		{"id", "name"},
		{"1", "Foo"},
		{"2", "Bar"},
	}
	if !rowsEqual(source.Rows, expected) {
		t.Fatalf("Unexpected --concat results: %v", source.Rows)
	}
}
