package main

// Arguments to format are:
//
//	[1]: type name
const bsonMethods = `
// MarshalBSONValue implements the bson.ValueMarshaler interface for %[1]s
func (i %[1]s) MarshalBSONValue() (bsontype.Type, []byte, error) {
	return bson.TypeString, bsoncore.AppendString([]byte{}, i.String()), nil
}

// UnmarshalBSONValue implements the bson.ValueUnmarshaler interface for %[1]s
func (i *%[1]s) UnmarshalBSONValue(t bsontype.Type, src []byte) (err error) {
	if t != bson.TypeString {
		return fmt.Errorf("invalid BSON Type, expecting string got %%s", t.String())
	}
	str, _, ok := bsoncore.ReadString(src)
	if !ok {
		return fmt.Errorf("error reading string from bson")
	}
	*i, err = %[1]sString(str)
	return err
}
`

func (g *Generator) buildBSONMethods(runs [][]Value, typeName string, runsThreshold int) {
	g.Printf(bsonMethods, typeName)
}
