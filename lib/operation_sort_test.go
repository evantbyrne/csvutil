package lib

import (
	"testing"
)

func TestOperationSort(t *testing.T) {
	expected := [][]string{
		{"year", "cost"},
		{"10", "500"},
		{"2", "200.10"},
		{"2000", "50.01"},
		{"2019", "1000.99"},
		{"x", "x"},
		{"y", "z"},
		{"z", "y"},
	}
	source := &Source{
		Rows: [][]string{
			{"year", "cost"},
			{"10", "500"},
			{"2", "200.10"},
			{"2000", "50.01"},
			{"x", "x"},
			{"2019", "1000.99"},
			{"y", "z"},
			{"z", "y"},
		},
	}
	operation := &OperationSort{}
	if err, _ := operation.Construct(source, []string{"--sort"}); err == nil {
		t.Fatal("Expected --sort to fail without a column name")
	}
	if err, _ := operation.Construct(source, []string{"--sort", "year"}); err == nil {
		t.Fatal("Expected --sort to fail without an algorithm")
	}
	if err, _ := operation.Construct(source, []string{"--sort", "year", "FOO"}); err == nil {
		t.Fatal("Expected --sort to fail without a valid algorithm")
	}
	if err, _ := operation.Construct(source, []string{"--sort", "year", "ALPHA"}); err == nil {
		t.Fatal("Expected --sort to fail without an order")
	}
	if err, _ := operation.Construct(source, []string{"--sort", "year", "ALPHA", "BAR"}); err == nil {
		t.Fatal("Expected --sort to fail without a valid order")
	}
	if err, _ := operation.Construct(source, []string{"--sort", "year", "ALPHA", "ASC"}); err != nil {
		t.Fatalf("Unexpected --sort failure: %s", err)
	}
	if err := operation.Run(source); err != nil {
		t.Fatalf("Unexpected --sort failure: %s", err)
	}
	if !rowsEqual(source.Rows, expected) {
		t.Fatalf("Unexpected --sort results: %v", source.Rows)
	}

	expected = [][]string{
		{"year", "cost"},
		{"z", "y"},
		{"y", "z"},
		{"x", "x"},
		{"2019", "1000.99"},
		{"2000", "50.01"},
		{"2", "200.10"},
		{"10", "500"},
	}
	if err, _ := operation.Construct(source, []string{"--sort", "year", "ALPHA", "DESC"}); err != nil {
		t.Fatalf("Unexpected --sort failure: %s", err)
	}
	if err := operation.Run(source); err != nil {
		t.Fatalf("Unexpected --sort failure: %s", err)
	}
	if !rowsEqual(source.Rows, expected) {
		t.Fatalf("Unexpected --sort results: %v", source.Rows)
	}

	expected = [][]string{
		{"year", "cost"},
		{"x", "x"},
		{"y", "z"},
		{"z", "y"},
		{"2", "200.10"},
		{"10", "500"},
		{"2000", "50.01"},
		{"2019", "1000.99"},
	}
	if err, _ := operation.Construct(source, []string{"--sort", "year", "INT", "ASC"}); err != nil {
		t.Fatalf("Unexpected --sort failure: %s", err)
	}
	if err := operation.Run(source); err != nil {
		t.Fatalf("Unexpected --sort failure: %s", err)
	}
	if !rowsEqual(source.Rows, expected) {
		t.Fatalf("Unexpected --sort results: %v", source.Rows)
	}

	expected = [][]string{
		{"year", "cost"},
		{"2019", "1000.99"},
		{"2", "200.10"},
		{"2000", "50.01"},
		{"x", "x"},
		{"z", "y"},
		{"y", "z"},
		{"10", "500"},
	}
	if err, _ := operation.Construct(source, []string{"--sort", "cost", "INT", "ASC"}); err != nil {
		t.Fatalf("Unexpected --sort failure: %s", err)
	}
	if err := operation.Run(source); err != nil {
		t.Fatalf("Unexpected --sort failure: %s", err)
	}
	if !rowsEqual(source.Rows, expected) {
		t.Fatalf("Unexpected --sort results: %v", source.Rows)
	}

	expected = [][]string{
		{"year", "cost"},
		{"x", "x"},
		{"z", "y"},
		{"y", "z"},
		{"2000", "50.01"},
		{"2", "200.10"},
		{"10", "500"},
		{"2019", "1000.99"},
	}
	if err, _ := operation.Construct(source, []string{"--sort", "cost", "FLOAT", "ASC"}); err != nil {
		t.Fatalf("Unexpected --sort failure: %s", err)
	}
	if err := operation.Run(source); err != nil {
		t.Fatalf("Unexpected --sort failure: %s", err)
	}
	if !rowsEqual(source.Rows, expected) {
		t.Fatalf("Unexpected --sort results: %v", source.Rows)
	}
}
