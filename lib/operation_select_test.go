package lib

import (
	"testing"
)

func TestOperationSelect(t *testing.T) {
	source := &Source{
		Rows: [][]string{
			{"username", "id", "group_id"},
			{"foo", "1", "10"},
			{"bar", "2", ""},
			{"baz", "3", "20"},
			{"four", "4", "20"},
			{"five", "5", ""},
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
		{"group_id", "username"},
		{"10", "foo"},
		{"", "bar"},
		{"20", "baz"},
		{"20", "four"},
		{"", "five"},
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
