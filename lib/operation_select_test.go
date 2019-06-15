package lib

import (
	"testing"
)

func TestOperationSelect(t *testing.T) {
	source := &Source{
		Rows: [][]string{
			[]string{"username", "id", "group_id"},
			[]string{"foo", "1", "10"},
			[]string{"bar", "2", ""},
			[]string{"baz", "3", "20"},
			[]string{"four", "4", "20"},
			[]string{"five", "5", ""},
		},
	}
	operation := &OperationSelect{}
	if err, _ := operation.Construct(source, []string{"--select"}); err == nil {
		t.Fatal("Expected --select to fail without a comma-separated list of column names")
	}
	if err, _ := operation.Construct(source, []string{"--select", "foo"}); err != nil {
		t.Fatalf("Unexpected --select failure: %s", err)
	}
	if err := operation.Run(source); err == nil {
		t.Fatal("Expected --select to fail when run with invalid column.")
	}

	expected := [][]string{
		[]string{"group_id", "username"},
		[]string{"10", "foo"},
		[]string{"", "bar"},
		[]string{"20", "baz"},
		[]string{"20", "four"},
		[]string{"", "five"},
	}
	if err, _ := operation.Construct(source, []string{"--select", "group_id,username"}); err != nil {
		t.Fatalf("Unexpected --select failure: %s", err)
	}
	if err := operation.Run(source); err != nil {
		t.Fatalf("Unexpected --select failure: %s", err)
	}
	if !rowsEqual(source.Rows, expected) {
		t.Fatalf("Unexpected --select results: %v", source.Rows)
	}
}
