package lib

type OperationValues struct {
	columns []string
}

func (this *OperationValues) Construct(source *Source, args []string) (error, []string) {
	return nil, args[1:]
}

func (this *OperationValues) Run(source *Source) error {
	source.Rows = source.Rows[1:]
	return nil
}
