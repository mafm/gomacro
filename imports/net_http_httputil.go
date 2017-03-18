// this file was generated by gomacro command: import "net/http/httputil"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	"net/http/httputil"
)

func init() {
	Packages["net/http/httputil"] = Package{
	Binds: map[string]Value{
		"DumpRequest":	ValueOf(httputil.DumpRequest),
		"DumpRequestOut":	ValueOf(httputil.DumpRequestOut),
		"DumpResponse":	ValueOf(httputil.DumpResponse),
		"ErrClosed":	ValueOf(&httputil.ErrClosed).Elem(),
		"ErrLineTooLong":	ValueOf(&httputil.ErrLineTooLong).Elem(),
		"ErrPersistEOF":	ValueOf(&httputil.ErrPersistEOF).Elem(),
		"ErrPipeline":	ValueOf(&httputil.ErrPipeline).Elem(),
		"NewChunkedReader":	ValueOf(httputil.NewChunkedReader),
		"NewChunkedWriter":	ValueOf(httputil.NewChunkedWriter),
		"NewClientConn":	ValueOf(httputil.NewClientConn),
		"NewProxyClientConn":	ValueOf(httputil.NewProxyClientConn),
		"NewServerConn":	ValueOf(httputil.NewServerConn),
		"NewSingleHostReverseProxy":	ValueOf(httputil.NewSingleHostReverseProxy),
	},
	Types: map[string]Type{
		"BufferPool":	TypeOf((*httputil.BufferPool)(nil)).Elem(),
		"ClientConn":	TypeOf((*httputil.ClientConn)(nil)).Elem(),
		"ReverseProxy":	TypeOf((*httputil.ReverseProxy)(nil)).Elem(),
		"ServerConn":	TypeOf((*httputil.ServerConn)(nil)).Elem(),
	},
	Proxies: map[string]Type{
		"BufferPool":	TypeOf((*BufferPool_net_http_httputil)(nil)).Elem(),
	} }
}

// --------------- proxy for net/http/httputil.BufferPool ---------------
type BufferPool_net_http_httputil struct {
	Get_	func() []byte
	Put_	func([]byte) 
}
func (Obj BufferPool_net_http_httputil) Get() []byte {
	return Obj.Get_()
}
func (Obj BufferPool_net_http_httputil) Put(unnamed0 []byte)  {
	Obj.Put_(unnamed0)
}
