// this file was generated by gomacro command: import "go/parser"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	"go/parser"
)

func init() {
	Packages["go/parser"] = Package{
	Binds: map[string]Value{
		"AllErrors":	ValueOf(parser.AllErrors),
		"DeclarationErrors":	ValueOf(parser.DeclarationErrors),
		"ImportsOnly":	ValueOf(parser.ImportsOnly),
		"PackageClauseOnly":	ValueOf(parser.PackageClauseOnly),
		"ParseComments":	ValueOf(parser.ParseComments),
		"ParseDir":	ValueOf(parser.ParseDir),
		"ParseExpr":	ValueOf(parser.ParseExpr),
		"ParseExprFrom":	ValueOf(parser.ParseExprFrom),
		"ParseFile":	ValueOf(parser.ParseFile),
		"SpuriousErrors":	ValueOf(parser.SpuriousErrors),
		"Trace":	ValueOf(parser.Trace),
	},
	Types: map[string]Type{
		"Mode":	TypeOf((*parser.Mode)(nil)).Elem(),
	},
	Proxies: map[string]Type{
	} }
}
