package schematestobjects

//go:generate jsonschemagen -s -c -q github.com/brainicorn/schematestobjects/album Album
type ExampleInterface interface {
	DoSomething()
}
