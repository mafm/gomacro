/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017 Massimiliano Ghilardi
 *
 *     This program is free software you can redistribute it and/or modify
 *     it under the terms of the GNU General Public License as published by
 *     the Free Software Foundation, either version 3 of the License, or
 *     (at your option) any later version.
 *
 *     This program is distributed in the hope that it will be useful,
 *     but WITHOUT ANY WARRANTY; without even the implied warranty of
 *     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *     GNU General Public License for more details.
 *
 *     You should have received a copy of the GNU General Public License
 *     along with this program.  If not, see <http//www.gnu.org/licenses/>.
 *
 * declaration.go
 *
 *  Created on Apr 01, 2017
 *      Author Massimiliano Ghilardi
 */

package fast_interpreter

import (
	"go/ast"
	"go/token"
	r "reflect"
)

// Decl compiles a constant, variable, function or type declaration - or an import
func (c *Comp) Decl(node ast.Decl) X {
	switch node := node.(type) {
	case *ast.GenDecl:
		return ExprStmtsToX(c.DeclGen(node))
	case *ast.FuncDecl:
		return c.DeclFunc(node, node.Type, node.Body)
	default:
		c.Errorf("Compile: unsupported declaration, expecting <*ast.GenDecl> or <*ast.FuncDecl>, found: %v <%v>", node, TypeOf(node))
		return nil
	}
}

// Decl compiles a constant, variable or type declaration - or an import
func (c *Comp) DeclGen(node *ast.GenDecl) []X {
	var funs []X
	switch node.Tok {
	case token.IMPORT:
		for _, decl := range node.Specs {
			c.Import(decl)
		}
	case token.CONST:
		for _, decl := range node.Specs {
			c.DeclConsts(decl)
		}
	case token.TYPE:
		for _, decl := range node.Specs {
			c.DeclType(decl)
		}
	case token.VAR:
		for _, decl := range node.Specs {
			funs = append(funs, c.DeclVars(decl)...)
		}
	default:
		c.Errorf("Compile: unsupported declaration kind, expecting token.IMPORT, token.CONST, token.TYPE or token.VAR, found %v: %v %<v>",
			node.Tok, node, TypeOf(node))
	}
	return funs
}

// DeclConsts compiles a set of constant declarations
func (c *Comp) DeclConsts(node ast.Spec) {
	switch node := node.(type) {
	case *ast.ValueSpec:
		names, t, inits := c.prepareDeclConstsOrVars(node.Names, node.Type, node.Values)
		c.DeclConsts0(names, t, inits)
	default:
		c.Errorf("Compile: unsupported constant declaration: expecting <*ast.ValueSpec>, found: %v <%v>", node, TypeOf(node))
	}
}

// DeclVars compiles a set of variable declarations
func (c *Comp) DeclVars(node ast.Spec) []X {
	switch node := node.(type) {
	case *ast.ValueSpec:
		names, t, inits := c.prepareDeclConstsOrVars(node.Names, node.Type, node.Values)
		return c.DeclVars0(names, t, inits)
	default:
		c.Errorf("Compile: unsupported variable declaration: expecting <*ast.ValueSpec>, found: %v <%v>", node, TypeOf(node))
		return nil
	}
}

func (c *Comp) prepareDeclConstsOrVars(idents []*ast.Ident, typ ast.Expr, exprs []ast.Expr) (names []string, t r.Type, inits []*Expr) {
	n := len(idents)
	names = make([]string, n)
	for i, ident := range idents {
		names[i] = ident.Name
	}
	if typ != nil {
		t = c.CompileType(typ)
	}
	if exprs != nil {
		inits = c.ExprsMultipleValues(exprs, n)
	}
	return names, t, inits
}

func (c *Comp) DeclConsts0(names []string, t r.Type, inits []*Expr) {
	n := len(names)
	if inits == nil {
		c.Errorf("constants without initialization: %v", names)
	} else if len(inits) != n {
		c.Errorf("cannot declare %d constants with %d initializers: %v", names)
	}
	for i, name := range names {
		init := inits[i]
		if !init.Const() {
			c.Errorf("const initializer for %q is not a constant", name)
		}
		c.DeclConst0(name, t, init.Value)
	}
}

// DeclVars0 compiles a set of variable declarations
func (c *Comp) DeclVars0(names []string, t r.Type, inits []*Expr) []X {
	n := len(names)
	ni := len(inits)
	var funs []X
	if ni == 0 {
		funs = make([]X, n)
		for i := 0; i < n; i++ {
			funs[i] = c.DeclVar0(names[i], t, nil)
		}
	} else if ni == n {
		funs = make([]X, n)
		for i := 0; i < n; i++ {
			funs[i] = c.DeclVar0(names[i], t, inits[i])
		}
	} else if ni == 1 && n > 1 {
		funs = append(funs, c.DeclMultiVar0(names, t, inits[0]))
	} else {
		c.Errorf("cannot declare %d variables from %d expressions: %v", n, ni, names)
	}
	return funs
}

// DeclConst0 compiles a constant declaration
func (c *Comp) DeclConst0(name string, t r.Type, value I) {
	if !isLiteral(value) {
		c.Errorf("const initializer for %q is not a constant", name)
		return
	}
	bind := BindValue(value)
	if t != nil {
		value = bind.ConstTo(t)
	}
	// never define bindings for "_"
	if name == "_" {
		return
	}
	if _, ok := c.Binds[name]; ok {
		c.Warnf("redefined identifier: %v", name)
	} else if c.Binds == nil {
		c.Binds = make(map[string]Bind)
	}
	c.Binds[name] = bind
}

// AddBind reserves space for a subsequent variable declaration
// returns NoBindDescriptor if name == "_"
func (c *Comp) AddVarBind(name string, t r.Type) Bind {
	return c.addBind(name, t, true)
}

// AddBind reserves space for a subsequent function declaration
// returns NoBindDescriptor if name == "_"
func (c *Comp) AddFuncBind(name string, t r.Type) Bind {
	return c.addBind(name, t, false)
}

func (c *Comp) addBind(name string, t r.Type, settable bool) Bind {
	var desc BindDescriptor
	if name != "_" {
		if bind, ok := c.Binds[name]; ok {
			c.Warnf("redefined identifier: %v", name)
			return bind
		} else if c.Binds == nil {
			c.Binds = make(map[string]Bind)
		}
		// allocate a slot either in IntBinds or in Binds
		if isCategory(t.Kind(), r.Bool, r.Int, r.Uint, r.Float64) || t.Kind() == r.Complex64 {
			desc = MakeBindDescriptor(IntBind, c.IntBindNum)
			c.IntBindNum++
		} else if settable {
			desc = MakeBindDescriptor(VarBind, c.BindNum)
			c.BindNum++
		} else {
			desc = MakeBindDescriptor(FuncBind, c.BindNum)
			c.BindNum++
		}
	}
	bind := Bind{Lit: Lit{Type: t}, Desc: desc}
	c.Binds[name] = bind
	return bind
}

// DeclVar0 compiles a variable declaration
func (c *Comp) DeclVar0(name string, t r.Type, init *Expr) X {
	if t == nil {
		if init == nil {
			c.Errorf("no value and no type, cannot declare variable: %v", name)
		}
		t = init.Type
		if t == nil {
			c.Errorf("cannot declare variable as untyped nil: %v", name)
		}
		n := init.NumOut()
		if n == 0 {
			c.Errorf("initializer returns no values, cannot declare variable: %v", name)
		} else if n > 1 {
			c.Errorf("initializer returns %d values, using only the first one to declare variable: %v", name)
		}
	}
	desc := c.AddVarBind(name, t).Desc
	switch desc.Class() {
	case ConstBind, IntBind, FuncBind:
		// no difference between declaration and assignment for these classes
		if init == nil {
			// no initializer... use the zero-value of t
			init = ExprValue(r.Zero(t).Interface())
		}
		return c.AssignVar0(name, desc, t, init)
	default:
		// declaring a variable in Env.Binds[], we must create a settable and addressable reflect.Value
		index := desc.Index()

		if init == nil {
			// no initializer... use the zero-value of t
			return func(env *Env) {
				env.Binds[index] = r.New(t).Elem()
			}
		}
		fun := init.AsX1() // AsX1() panics if init.NumOut() == 0, warns if init.NumOut() > 1
		t := init.Out(0)
		// optimization: no need to wrap multiple-valued function into a single-value function
		if f, ok := init.Fun.(func(*Env) (r.Value, []r.Value)); ok {
			return func(env *Env) {
				ret, _ := f(env)
				place := r.New(t).Elem()
				place.Set(ret)
				env.Binds[index] = place
			}
		}
		return func(env *Env) {
			place := r.New(t).Elem()
			place.Set(fun(env))
			env.Binds[index] = place
		}
	}
}

// DeclMultiVar0 compiles multiple variable declarations from a single multi-valued expression
func (c *Comp) DeclMultiVar0(names []string, t r.Type, init *Expr) X {
	if t == nil {
		if init == nil {
			c.Errorf("no value and no type, cannot declare variables: %v", names)
		}
	}
	n := len(names)
	if n == 1 {
		return c.DeclVar0(names[0], t, init)
	}
	ni := init.NumOut()
	if ni < n {
		c.Errorf("cannot declare %d variables from expression returning %d values: %v", n, ni, names)
	} else if ni > n {
		c.Warnf("declaring %d variables from expression returning %d values: %v", n, ni, names)
	}
	decls := make([]X, n)
	addrs := make([]func(*Env) r.Value, n)
	for i, name := range names {
		ti := init.Out(i)
		if t != nil && t != ti {
			if !ti.AssignableTo(t) {
				c.Errorf("cannot assign <%v> to <%v> in variable declaration: %v", ti, t, names)
				return nil
			} else {
				ti = t // declared variable has type t, not the i-th type returned by multi-valued expression
			}
		}
		bind := c.AddVarBind(name, ti)
		switch bind.Desc.Class() {
		case IntBind:
			// to set the variable, go through its address. a bit slow, uses reflect.Value...
			addrs[i] = c.identIntBindAddress(name, 0, bind).AsX1()
			fallthrough
		case ConstBind, FuncBind: // ConstBind happens if name == "_"
			// no difference between declaration and assignment for these classes
			valuei := ExprValue(r.Zero(ti).Interface())
			decls[i] = c.AssignVar0(name, bind.Desc, ti, valuei)
		default:
			// declaring a variable in Env.Binds[], we must create a settable and addressable reflect.Value
			index := bind.Desc.Index()
			decls[i] = func(env *Env) {
				env.Binds[index] = r.New(ti).Elem()
			}
			// to set the variable, go through its address
			addrs[i] = c.identBindAddress(name, 0, bind).AsX1()
		}

	}
	fun := init.AsXV()
	return func(env *Env) {
		// call the multi-valued function. we know ni > 1, so just use the []r.Value
		_, rets := fun(env)

		// declare and assign the variables one by one. we know n <= ni
		for i, decl := range decls {
			decl(env)
			if addr := addrs[i]; addr != nil {
				addr(env).Elem().Set(rets[i])
			}
		}
	}
}

/*
func (c *Comp) DeclFunc(name string, paramTypes []r.Type, body X) X {
	idx := c.AddBind(name, typeOfInterface) // FIXME need accurate function type
	xf := c.MakeFunc(paramTypes, body)
	return func(env *Env) (r.Value, []r.Value) {
		f := xf(env)
		env.Binds[idx] = r.ValueOf(f)
		return r.ValueOf(f), nil
	}
}

func (c *Comp) MakeFunc(paramTypes []r.Type, body X) XFunc {
	return func(env *Env) Func {
		return func(args ...r.Value) (ret r.Value, rets []r.Value) {
			fenv := NewEnv(env, 10)
			panicking := true // use a flag to distinguish non-panic from panic(nil)
			defer func() {
				if panicking {
					pan := recover()
					switch p := pan.(type) {
					case SReturn:
						// return is implemented with a panic(SReturn{})
						ret = p.result0
						rets = p.results
					default:
						panic(pan)
					}
				}
			}()
			for i, paramType := range paramTypes {
				place := r.New(paramType).Elem()
				place.Set(args[i])
				fenv.Binds[i] = place
			}
			ret, rets = body(fenv)
			panicking = false
			return ret, rets
		}
	}
}

func MakeFuncInt(paramTypes []r.Type, body X) XFuncInt {
	return func(env *Env) FuncInt {
		return func(args ...r.Value) (ret int) {
			fenv := NewEnv(env, 10)
			panicking := true // use a flag to distinguish non-panic from panic(nil)
			defer func() {
				if panicking {
					pan := recover()
					switch p := pan.(type) {
					case SReturn:
						// return is implemented with a panic(cReturn{})
						ret = int(p.result0.Int())
					default:
						panic(pan)
					}
				}
			}()
			for i, paramType := range paramTypes {
				place := r.New(paramType).Elem()
				place.Set(args[i])
				fenv.Binds[i] = place
			}
			ret0, _ := body(fenv)
			panicking = false
			return int(ret0.Int())
		}
	}
}
*/