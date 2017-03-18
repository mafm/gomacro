// this file was generated by gomacro command: import "image/draw"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	"image"
	"image/color"
	"image/draw"
)

func init() {
	Packages["image/draw"] = Package{
	Binds: map[string]Value{
		"Draw":	ValueOf(draw.Draw),
		"DrawMask":	ValueOf(draw.DrawMask),
		"FloydSteinberg":	ValueOf(&draw.FloydSteinberg).Elem(),
		"Over":	ValueOf(draw.Over),
		"Src":	ValueOf(draw.Src),
	},
	Types: map[string]Type{
		"Drawer":	TypeOf((*draw.Drawer)(nil)).Elem(),
		"Image":	TypeOf((*draw.Image)(nil)).Elem(),
		"Op":	TypeOf((*draw.Op)(nil)).Elem(),
		"Quantizer":	TypeOf((*draw.Quantizer)(nil)).Elem(),
	},
	Proxies: map[string]Type{
		"Drawer":	TypeOf((*Drawer_image_draw)(nil)).Elem(),
		"Image":	TypeOf((*Image_image_draw)(nil)).Elem(),
		"Quantizer":	TypeOf((*Quantizer_image_draw)(nil)).Elem(),
	} }
}

// --------------- proxy for image/draw.Drawer ---------------
type Drawer_image_draw struct {
	Draw_	func(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) 
}
func (Obj Drawer_image_draw) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point)  {
	Obj.Draw_(dst, r, src, sp)
}

// --------------- proxy for image/draw.Image ---------------
type Image_image_draw struct {
	At_	func(x int, y int) color.Color
	Bounds_	func() image.Rectangle
	ColorModel_	func() color.Model
	Set_	func(x int, y int, c color.Color) 
}
func (Obj Image_image_draw) At(x int, y int) color.Color {
	return Obj.At_(x, y)
}
func (Obj Image_image_draw) Bounds() image.Rectangle {
	return Obj.Bounds_()
}
func (Obj Image_image_draw) ColorModel() color.Model {
	return Obj.ColorModel_()
}
func (Obj Image_image_draw) Set(x int, y int, c color.Color)  {
	Obj.Set_(x, y, c)
}

// --------------- proxy for image/draw.Quantizer ---------------
type Quantizer_image_draw struct {
	Quantize_	func(p color.Palette, m image.Image) color.Palette
}
func (Obj Quantizer_image_draw) Quantize(p color.Palette, m image.Image) color.Palette {
	return Obj.Quantize_(p, m)
}
