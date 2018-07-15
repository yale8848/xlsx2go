package format

import "strconv"

//FontFamilyType is a type to encode XSD ST_FontFamily
type FontFamilyType byte

//List of all possible values for FontFamilyType
const (
	_ FontFamilyType = iota
	FontFamilyRoman
	FontFamilySwiss
	FontFamilyModern
	FontFamilyScript
	FontFamilyDecorative
)

func (v FontFamilyType) String() string {
	return strconv.Itoa(int(v))
}
