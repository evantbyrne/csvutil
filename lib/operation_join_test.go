package lib

import (
	"testing"
)

func TestOperationJoin(t *testing.T) {
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
	operation := &OperationJoin{}
	if err, _ := operation.Construct(source, []string{"--join"}); err == nil {
		t.Fatal("Expected --join to fail without previous source column name.")
	}
	if err, _ := operation.Construct(source, []string{"--join", "group_id"}); err == nil {
		t.Fatal("Expected --join to fail without operator.")
	}
	if err, _ := operation.Construct(source, []string{"--join", "group_id", "=="}); err == nil {
		t.Fatal("Expected --join to fail without source column name.")
	}
	if err, _ := operation.Construct(source, []string{"--join", "group_id", "==", "id"}); err != nil {
		t.Fatalf("Unexpected --join failure: %s", err)
	}
	if err := operation.Run(source); err == nil {
		t.Fatal("Expected --join to fail when run on first source.")
	}

	expected := [][]string{
		[]string{"username", "id", "group_id", "name", "id"},
		[]string{"foo", "1", "10", "Admin", "10"},
		[]string{"baz", "3", "20", "Moderator", "20"},
		[]string{"four", "4", "20", "Moderator", "20"},
	}
	source = &Source{
		Previous: &Source{
			Rows: [][]string{
				[]string{"username", "id", "group_id"},
				[]string{"foo", "1", "10"},
				[]string{"bar", "2", ""},
				[]string{"baz", "3", "20"},
				[]string{"four", "4", "20"},
				[]string{"five", "5", ""},
			},
		},
		Rows: [][]string{
			[]string{"name", "id"},
			[]string{"Admin", "10"},
			[]string{"Moderator", "20"},
		},
	}
	if err, _ := operation.Construct(source, []string{"--join", "group_id", "==", "id"}); err != nil {
		t.Fatalf("Unexpected --join failure: %s", err)
	}
	if err := operation.Run(source); err != nil {
		t.Fatalf("Unexpected --join failure: %s", err)
	}
	if !rowsEqual(source.Rows, expected) {
		t.Fatalf("Unexpected --join results: %v", source.Rows)
	}
}
