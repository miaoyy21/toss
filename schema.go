package toss

import "fmt"

var (
	SchemaInvalid  Schema = '-'
	SchemaPositive Schema = 'P'
	SchemaNegative Schema = 'N'
)

type Schema byte

func (o Schema) Reverse() Schema {
	if o == SchemaPositive {
		return SchemaNegative
	} else if o == SchemaNegative {
		return SchemaPositive
	}

	panic(fmt.Sprintf("unreachable Schema of %q", string(o)))
}

func (o Schema) Byte() byte {
	if o == SchemaInvalid || o == SchemaPositive || o == SchemaNegative {
		return byte(o)
	}

	panic(fmt.Sprintf("unreachable Schema of %q", string(o)))
}

func (o Schema) String() string {
	if o == SchemaInvalid || o == SchemaPositive || o == SchemaNegative {
		return string(o)
	}

	panic(fmt.Sprintf("unreachable Schema of %q", string(o)))
}
