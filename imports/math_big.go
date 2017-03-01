// this file was generated by gomacro command: import "math/big"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	pkg "math/big"
	. "reflect"
)

func Package_math_big() (map[string]Value, map[string]Type) {
	return map[string]Value{
			"Above":         ValueOf(pkg.Above),
			"AwayFromZero":  ValueOf(pkg.AwayFromZero),
			"Below":         ValueOf(pkg.Below),
			"Exact":         ValueOf(pkg.Exact),
			"Jacobi":        ValueOf(pkg.Jacobi),
			"MaxBase":       ValueOf(pkg.MaxBase),
			"MaxExp":        ValueOf(pkg.MaxExp),
			"MaxPrec":       ValueOf(pkg.MaxPrec),
			"MinExp":        ValueOf(pkg.MinExp),
			"NewFloat":      ValueOf(pkg.NewFloat),
			"NewInt":        ValueOf(pkg.NewInt),
			"NewRat":        ValueOf(pkg.NewRat),
			"ParseFloat":    ValueOf(pkg.ParseFloat),
			"ToNearestAway": ValueOf(pkg.ToNearestAway),
			"ToNearestEven": ValueOf(pkg.ToNearestEven),
			"ToNegativeInf": ValueOf(pkg.ToNegativeInf),
			"ToPositiveInf": ValueOf(pkg.ToPositiveInf),
			"ToZero":        ValueOf(pkg.ToZero),
		}, map[string]Type{
			"Accuracy":     TypeOf((*pkg.Accuracy)(nil)).Elem(),
			"ErrNaN":       TypeOf((*pkg.ErrNaN)(nil)).Elem(),
			"Float":        TypeOf((*pkg.Float)(nil)).Elem(),
			"Int":          TypeOf((*pkg.Int)(nil)).Elem(),
			"Rat":          TypeOf((*pkg.Rat)(nil)).Elem(),
			"RoundingMode": TypeOf((*pkg.RoundingMode)(nil)).Elem(),
			"Word":         TypeOf((*pkg.Word)(nil)).Elem(),
		}
}

func init() {
	binds, types := Package_math_big()
	Binds["math/big"] = binds
	Types["math/big"] = types
}