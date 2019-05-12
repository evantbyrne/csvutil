package lib

type OperationWhere struct {
	comparison *Comparison
}

func (this *OperationWhere) Construct(arg string) error {
	err, comparison := ConstructComparison("where", arg)
	if err != nil {
		return err
	}

	this.comparison = comparison
	return nil
}

func (this *OperationWhere) Run(source *Source) error {
	var rightIndex int
	if this.comparison.RightType == "column" {
		rightIndex = source.ColumnIndex(this.comparison.Right)
	}
	leftIndex := source.ColumnIndex(this.comparison.Left)

	var rows [][]string
	for _, row := range source.Rows {
		if this.comparison.RightType == "value" {
			if row[leftIndex] == this.comparison.Right {
				rows = append(rows, row)
			}
		} else {
			if row[leftIndex] == row[rightIndex] {
				rows = append(rows, row)
			}
		}
	}
	source.Rows = rows
	return nil
}
