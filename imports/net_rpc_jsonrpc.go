// this file was generated by gomacro command: import "net/rpc/jsonrpc"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	"net/rpc/jsonrpc"
)

func init() {
	Packages["net/rpc/jsonrpc"] = Package{
	Binds: map[string]Value{
		"Dial":	ValueOf(jsonrpc.Dial),
		"NewClient":	ValueOf(jsonrpc.NewClient),
		"NewClientCodec":	ValueOf(jsonrpc.NewClientCodec),
		"NewServerCodec":	ValueOf(jsonrpc.NewServerCodec),
		"ServeConn":	ValueOf(jsonrpc.ServeConn),
	},
	Types: map[string]Type{
	},
	Proxies: map[string]Type{
	} }
}
