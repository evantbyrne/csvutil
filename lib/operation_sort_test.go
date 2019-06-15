package lib

import (
	"testing"
)

func TestOperationSort(t *testing.T) {
	expected := [][]string{
		[]string{"year", "cost"},
		[]string{"10", "500"},
		[]string{"2", "200.10"},
		[]string{"2000", "50.01"},
		[]string{"2019", "1000.99"},
		[]string{"x", "x"},
		[]string{"y", "z"},
		[]string{"z", "y"},
	}
	source := &Source{
		Rows: [][]string{
			[]string{"year", "cost"},
			[]string{"10", "500"},
			[]string{"2", "200.10"},
			[]string{"2000", "50.01"},
			[]string{"x", "x"},
			[]string{"2019", "1000.99"},
			[]string{"y", "z"},
			[]string{"z", "y"},
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
		[]string{"year", "cost"},
		[]string{"z", "y"},
		[]string{"y", "z"},
		[]string{"x", "x"},
		[]string{"2019", "1000.99"},
		[]string{"2000", "50.01"},
		[]string{"2", "200.10"},
		[]string{"10", "500"},
	}
	if err, _ := operation.Construct(source, []string{"--sort", "year", "ALPHA", "DESC"}); err != nil {
		t.Fatalf("Unexpected --sort failure: %s", err)
	}
	if err := operation.Run(source); err != nil {
		t.Fatalf("Unexpected --sort failure: %s", err)
	}

	expected = [][]string{
		[]string{"year", "cost"},
		[]string{"2", "200.10"},
		[]string{"10", "500"},
		[]string{"2000", "50.01"},
		[]string{"2019", "1000.99"},
		[]string{"x", "x"},
		[]string{"y", "z"},
		[]string{"z", "y"},
	}
	if err, _ := operation.Construct(source, []string{"--sort", "year", "INT", "ASC"}); err != nil {
		t.Fatalf("Unexpected --sort failure: %s", err)
	}
	if err := operation.Run(source); err != nil {
		t.Fatalf("Unexpected --sort failure: %s", err)
	}

	expected = [][]string{
		[]string{"year", "cost"},
		[]string{"10", "500"},
		[]string{"y", "z"},
		[]string{"z", "y"},
		[]string{"x", "x"},
		[]string{"2000", "50.01"},
		[]string{"2", "200.10"},
		[]string{"2019", "1000.99"},
	}
	if err, _ := operation.Construct(source, []string{"--sort", "cost", "INT", "ASC"}); err != nil {
		t.Fatalf("Unexpected --sort failure: %s", err)
	}
	if err := operation.Run(source); err != nil {
		t.Fatalf("Unexpected --sort failure: %s", err)
	}

	expected = [][]string{
		[]string{"year", "cost"},
		[]string{"x", "x"},
		[]string{"y", "z"},
		[]string{"z", "y"},
		[]string{"2000", "50.01"},
		[]string{"2", "200.10"},
		[]string{"10", "500"},
		[]string{"2019", "1000.99"},
	}
	if err, _ := operation.Construct(source, []string{"--sort", "cost", "FLOAT", "ASC"}); err != nil {
		t.Fatalf("Unexpected --sort failure: %s", err)
	}
	if err := operation.Run(source); err != nil {
		t.Fatalf("Unexpected --sort failure: %s", err)
	}
}
