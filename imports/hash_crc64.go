// this file was generated by gomacro command: import "hash/crc64"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	"hash/crc64"
)

func init() {
	Packages["hash/crc64"] = Package{
	Binds: map[string]Value{
		"Checksum":	ValueOf(crc64.Checksum),
		"ECMA":	ValueOf(uint64(crc64.ECMA)),
		"ISO":	ValueOf(uint64(crc64.ISO)),
		"MakeTable":	ValueOf(crc64.MakeTable),
		"New":	ValueOf(crc64.New),
		"Size":	ValueOf(crc64.Size),
		"Update":	ValueOf(crc64.Update),
	},
	Types: map[string]Type{
		"Table":	TypeOf((*crc64.Table)(nil)).Elem(),
	},
	Proxies: map[string]Type{
	} }
}
