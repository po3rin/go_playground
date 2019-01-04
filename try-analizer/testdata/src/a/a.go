package a

type StructTagTest struct {
	A   int "hello"            // want "`hello` not compatible with reflect.StructTag.Get: bad syntax for struct tag pair"
	B   int "\tx:\"y\""        // want "not compatible with reflect.StructTag.Get: bad syntax for struct tag key"
	C   int "x:\"y\"\tx:\"y\"" // want "not compatible with reflect.StructTag.Get"
	D   int "x:`y`"            // want "not compatible with reflect.StructTag.Get: bad syntax for struct tag value"
	E   int "ct\brl:\"char\""  // want "not compatible with reflect.StructTag.Get: bad syntax for struct tag pair"
	F   int `:"emptykey"`      // want "not compatible with reflect.StructTag.Get: bad syntax for struct tag key"
	G   int `x:"noEndQuote`    // want "not compatible with reflect.StructTag.Get: bad syntax for struct tag value"
	H   int `x:"trunc\x0"`     // want "not compatible with reflect.StructTag.Get: bad syntax for struct tag value"
	I   int `x:"foo",y:"bar"`  // want "not compatible with reflect.StructTag.Get: key:.value. pairs not separated by spaces"
	J   int `x:"foo"y:"bar"`   // want "not compatible with reflect.StructTag.Get: key:.value. pairs not separated by spaces"
	OK0 int `x:"y" u:"v" w:""`
	OK1 int `x:"y:z" u:"v" w:""` // note multiple colons.
	OK2 int "k0:\"values contain spaces\" k1:\"literal\ttabs\" k2:\"and\\tescaped\\tabs\""
	OK3 int `under_scores:"and" CAPS:"ARE_OK"`
}
