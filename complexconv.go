package complexconv

import (
	"errors"
	"fmt"
	"go/ast"
	"go/constant"
	"go/parser"
)

type complexParser struct {
	expr    ast.Expr
	bitSize int
}

var errCannotParse = errors.New("cannot parse")

func (p *complexParser) parse() (c complex128, err error) {
	defer func() {
		if r := recover(); r != nil {
			if e, ok := r.(error); ok {
				c, err = 0, e
			} else {
				c, err = 0, fmt.Errorf("%v", err)
			}
		}
	}()
	v, err := p.parseExpr(p.expr)
	if err != nil {
		return 0, err
	}
	if v.Kind() == constant.Int || v.Kind() == constant.Float {
		v = constant.ToComplex(v)
	}
	if v.Kind() == constant.Complex {
		return p.complexVal(v)
	}
	return 0, errCannotParse
}

func (p *complexParser) complexVal(v constant.Value) (complex128, error) {
	rv := constant.Real(v)
	iv := constant.Imag(v)
	// complex64
	if p.bitSize == 64 {
		r, _ := constant.Float32Val(rv)
		i, _ := constant.Float32Val(iv)
		return complex128(complex(r, i)), nil
	}
	// complex128
	r, _ := constant.Float64Val(rv)
	i, _ := constant.Float64Val(iv)
	return complex(r, i), nil
}

func (p *complexParser) parseExpr(expr ast.Expr) (constant.Value, error) {
	switch expr := expr.(type) {
	case *ast.UnaryExpr:
		return p.parseUnaryExpr(expr)
	case *ast.BinaryExpr:
		return p.parseBinaryExpr(expr)
	case *ast.BasicLit:
		return p.parseBasicLit(expr), nil
	case *ast.ParenExpr:
		return p.parseExpr(expr.X)
	}
	return constant.MakeUnknown(), errCannotParse
}

func (p *complexParser) parseUnaryExpr(expr *ast.UnaryExpr) (constant.Value, error) {
	x, err := p.parseExpr(expr.X)
	if err != nil {
		return constant.MakeUnknown(), err
	}
	return constant.UnaryOp(expr.Op, x, 0), nil
}

func (p *complexParser) parseBinaryExpr(expr *ast.BinaryExpr) (constant.Value, error) {
	x, err := p.parseExpr(expr.X)
	if err != nil {
		return constant.MakeUnknown(), err
	}
	y, err := p.parseExpr(expr.Y)
	if err != nil {
		return constant.MakeUnknown(), err
	}
	return constant.BinaryOp(x, expr.Op, y), nil
}

func (p *complexParser) parseBasicLit(expr *ast.BasicLit) constant.Value {
	return constant.MakeFromLiteral(expr.Value, expr.Kind, 0)
}

// ParseComplex returns the complex value represented by the string.
func ParseComplex(s string, bitSize int) (complex128, error) {
	if s == "" {
		return 0, nil
	}
	expr, err := parser.ParseExpr(s)
	if err != nil {
		return 0, err
	}
	p := &complexParser{
		expr:    expr,
		bitSize: bitSize,
	}
	c, err := p.parse()
	if err != nil {
		if err == errCannotParse {
			return 0, errors.New(`cannot parse "` + s + `" to complex.`)
		}
		return 0, err
	}
	return c, nil
}

// FormatComplex returns the string representation of c.
func FormatComplex(c complex128) string {
	return fmt.Sprintf("%g", c)
}
