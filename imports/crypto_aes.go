// this file was generated by gomacro command: import "crypto/aes"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	"crypto/aes"
)

func init() {
	Packages["crypto/aes"] = Package{
	Binds: map[string]Value{
		"BlockSize":	ValueOf(aes.BlockSize),
		"NewCipher":	ValueOf(aes.NewCipher),
	},
	Types: map[string]Type{
		"KeySizeError":	TypeOf((*aes.KeySizeError)(nil)).Elem(),
	},
	Proxies: map[string]Type{
	} }
}
