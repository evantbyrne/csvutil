package lib

type Operation interface {
	Construct(*Source, []string) (error, []string)
	Run(*Source) error
}
