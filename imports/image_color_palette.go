// this file was generated by gomacro command: import "image/color/palette"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	"image/color/palette"
)

func init() {
	Packages["image/color/palette"] = Package{
	Binds: map[string]Value{
		"Plan9":	ValueOf(&palette.Plan9).Elem(),
		"WebSafe":	ValueOf(&palette.WebSafe).Elem(),
	},
	Types: map[string]Type{
	},
	Proxies: map[string]Type{
	} }
}
