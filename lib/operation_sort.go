package lib

import (
	"errors"
	"sort"
	"strconv"
	"strings"
)

type OperationSort struct {
	column    string
	algorithm string
	order     string
}

func (this *OperationSort) Construct(source *Source, args []string) (error, []string) {
	if len(args) < 4 {
		return errors.New("--sort operation requires three arguments."), []string{}
	}

	this.column = args[1]
	this.algorithm = strings.ToUpper(args[2])
	this.order = strings.ToUpper(args[3])

	if this.algorithm != "ALPHA" && this.algorithm != "FLOAT" && this.algorithm != "INT" {
		return errors.New("--sort expects a valid sort algorithm as its second argument. Options include \"ALPHA\", \"FLOAT\", and \"INT\"."), []string{}
	}

	if this.order != "ASC" && this.order != "DESC" {
		return errors.New("--sort expects a valid sort order as its third argument. Options include \"ASC\" and \"DESC\"."), []string{}
	}

	return nil, args[4:]
}

func (this *OperationSort) Run(source *Source) error {
	columnIndex := source.ColumnIndex(this.column)
	header := source.Rows[0]
	rows := source.Rows[1:]
	sort.Slice(rows, func(i, j int) bool {
		// FLOAT
		if this.algorithm == "FLOAT" {
			left, errLeft := strconv.ParseFloat(rows[i][columnIndex], 64)
			right, errRight := strconv.ParseFloat(rows[j][columnIndex], 64)
			if errLeft != nil && errRight != nil {
				// Fallback to ALPHA sort.
			} else if errLeft != nil {
				if this.order == "DESC" {
					return false
				}
				return true
			} else if errRight != nil {
				if this.order == "DESC" {
					return true
				}
				return false
			} else if this.order == "DESC" {
				return left > right
			} else {
				return left < right
			}
		}

		// INT
		if this.algorithm == "INT" {
			left, errLeft := strconv.ParseInt(rows[i][columnIndex], 10, 64)
			right, errRight := strconv.ParseInt(rows[j][columnIndex], 10, 64)
			if errLeft != nil && errRight != nil {
				// Fallback to ALPHA sort.
			} else if errLeft != nil {
				if this.order == "DESC" {
					return false
				}
				return true
			} else if errRight != nil {
				if this.order == "DESC" {
					return true
				}
				return false
			} else if this.order == "DESC" {
				return left > right
			} else {
				return left < right
			}
		}

		// ALPHA
		if this.order == "DESC" {
			return strings.ToLower(rows[i][columnIndex]) > strings.ToLower(rows[j][columnIndex])
		}
		return strings.ToLower(rows[i][columnIndex]) < strings.ToLower(rows[j][columnIndex])
	})
	source.Rows = [][]string{header}
	source.Rows = append(source.Rows, rows...)
	return nil
}
