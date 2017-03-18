// this file was generated by gomacro command: import "encoding"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	"encoding"
)

func init() {
	Packages["encoding"] = Package{
	Binds: map[string]Value{
	},
	Types: map[string]Type{
		"BinaryMarshaler":	TypeOf((*encoding.BinaryMarshaler)(nil)).Elem(),
		"BinaryUnmarshaler":	TypeOf((*encoding.BinaryUnmarshaler)(nil)).Elem(),
		"TextMarshaler":	TypeOf((*encoding.TextMarshaler)(nil)).Elem(),
		"TextUnmarshaler":	TypeOf((*encoding.TextUnmarshaler)(nil)).Elem(),
	},
	Proxies: map[string]Type{
		"BinaryMarshaler":	TypeOf((*BinaryMarshaler_encoding)(nil)).Elem(),
		"BinaryUnmarshaler":	TypeOf((*BinaryUnmarshaler_encoding)(nil)).Elem(),
		"TextMarshaler":	TypeOf((*TextMarshaler_encoding)(nil)).Elem(),
		"TextUnmarshaler":	TypeOf((*TextUnmarshaler_encoding)(nil)).Elem(),
	} }
}

// --------------- proxy for encoding.BinaryMarshaler ---------------
type BinaryMarshaler_encoding struct {
	MarshalBinary_	func() (data []byte, err error)
}
func (Obj BinaryMarshaler_encoding) MarshalBinary() (data []byte, err error) {
	return Obj.MarshalBinary_()
}

// --------------- proxy for encoding.BinaryUnmarshaler ---------------
type BinaryUnmarshaler_encoding struct {
	UnmarshalBinary_	func(data []byte) error
}
func (Obj BinaryUnmarshaler_encoding) UnmarshalBinary(data []byte) error {
	return Obj.UnmarshalBinary_(data)
}

// --------------- proxy for encoding.TextMarshaler ---------------
type TextMarshaler_encoding struct {
	MarshalText_	func() (text []byte, err error)
}
func (Obj TextMarshaler_encoding) MarshalText() (text []byte, err error) {
	return Obj.MarshalText_()
}

// --------------- proxy for encoding.TextUnmarshaler ---------------
type TextUnmarshaler_encoding struct {
	UnmarshalText_	func(text []byte) error
}
func (Obj TextUnmarshaler_encoding) UnmarshalText(text []byte) error {
	return Obj.UnmarshalText_(text)
}
