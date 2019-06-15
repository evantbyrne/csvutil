package lib

import (
	"testing"
)

func TestOperationValues(t *testing.T) {
	source := &Source{
		Rows: [][]string{
			[]string{"foo", "bar"},
			[]string{"one", "two"},
		},
	}
	expected := [][]string{
		[]string{"one", "two"},
	}
	operation := &OperationValues{}
	if err, _ := operation.Construct(source, []string{"--values"}); err != nil {
		t.Fatalf("Unexpected --values failure: %s", err)
	}
	if err := operation.Run(source); err != nil {
		t.Fatalf("Unexpected --select failure: %s", err)
	}
	if !rowsEqual(source.Rows, expected) {
		t.Fatalf("Unexpected --select results: %v", source.Rows)
	}

	source = &Source{
		Rows: [][]string{
			[]string{"total"},
			[]string{"2"},
		},
	}
	expected = [][]string{
		[]string{"2"},
	}
	if err := operation.Run(source); err != nil {
		t.Fatalf("Unexpected --select failure: %s", err)
	}
	if !rowsEqual(source.Rows, expected) {
		t.Fatalf("Unexpected --select results: %v", source.Rows)
	}
}
