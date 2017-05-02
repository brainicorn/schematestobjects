package xof

// @jsonSchema(allOf=["github.com/brainicorn/schematestobjects/xof/AThing"])
type AllOf interface{}

// @jsonSchema(anyOf=["github.com/brainicorn/schematestobjects/xof/AThing"])
type AnyOf interface{}

// @jsonSchema(oneOf=["github.com/brainicorn/schematestobjects/xof/AThing"])
type OneOf interface{}

// @jsonSchema(oneOf=["github.com/brainicorn/schematestobjects/xof/AThing"]
//		,not="github.com/brainicorn/schematestobjects/xof/AnotherThing")
type Not interface{}

type AThing struct {
	Name string
}

type AnotherThing struct {
	Name string
}
