// this file was generated by gomacro command: import "io/ioutil"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	"io/ioutil"
)

func init() {
	Packages["io/ioutil"] = Package{
	Binds: map[string]Value{
		"Discard":	ValueOf(&ioutil.Discard).Elem(),
		"NopCloser":	ValueOf(ioutil.NopCloser),
		"ReadAll":	ValueOf(ioutil.ReadAll),
		"ReadDir":	ValueOf(ioutil.ReadDir),
		"ReadFile":	ValueOf(ioutil.ReadFile),
		"TempDir":	ValueOf(ioutil.TempDir),
		"TempFile":	ValueOf(ioutil.TempFile),
		"WriteFile":	ValueOf(ioutil.WriteFile),
	},
	Types: map[string]Type{
	},
	Proxies: map[string]Type{
	} }
}
