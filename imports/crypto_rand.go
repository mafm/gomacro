// this file was generated by gomacro command: import "crypto/rand"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	"crypto/rand"
)

func init() {
	Packages["crypto/rand"] = Package{
	Binds: map[string]Value{
		"Int":	ValueOf(rand.Int),
		"Prime":	ValueOf(rand.Prime),
		"Read":	ValueOf(rand.Read),
		"Reader":	ValueOf(&rand.Reader).Elem(),
	},
	Types: map[string]Type{
	},
	Proxies: map[string]Type{
	} }
}
