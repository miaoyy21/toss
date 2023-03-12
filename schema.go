package toss

import "fmt"

var SchemaPositive Schema = 'P'
var SchemaNegative Schema = 'N'

type Schema byte

func (o Schema) Byte() byte {
	if o == SchemaPositive || o == SchemaNegative {
		return byte(o)
	}

	panic(fmt.Sprintf("unreachable Schema of %q", string(o)))
}

func (o Schema) String() string {
	if o == SchemaPositive || o == SchemaNegative {
		return string(o)
	}

	panic(fmt.Sprintf("unreachable Schema of %q", string(o)))
}
