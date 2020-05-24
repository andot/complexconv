package complexconv

import (
	"testing"
)

var str = []string{
	"0i",
	"011i",
	"-1i",
	"2.71828",
	"1E6i",
	"1.e+0i",
	"0",
	"3+2i",
	"2.7e-4+2i",
	"7.7e+4+0.9i",
	"3i+4.0",
	"2.1i+4.55e+10",
	""}
var c128 = []complex128{
	0i,
	011i,
	-1i,
	2.71828,
	1E6i,
	1.e+0i,
	0,
	3.0 + 2i,
	2.7e-4 + 2i,
	7.7e+4 + 0.9i,
	3i + 4.0,
	2.1i + 4.55e+10,
	0}
var c64 = []complex64{
	0i,
	011i,
	-1i,
	2.71828,
	1E6i,
	1.e+0i,
	0,
	3.0 + 2i,
	2.7e-4 + 2i,
	7.7e+4 + 0.9i,
	3i + 4.0,
	2.1i + 4.55e+10,
	0}
var fstr = []string{
	"(0+0i)",
	"(0+11i)",
	"(0-1i)",
	"(2.71828+0i)",
	"(0+1e+06i)",
	"(0+1i)",
	"(0+0i)",
	"(3+2i)",
	"(0.00027+2i)",
	"(77000+0.9i)",
	"(4+3i)",
	"(4.55e+10+2.1i)",
	"(0+0i)"}

func TestParseComplex(t *testing.T) {
	for i := range str {
		c, e := ParseComplex(str[i], 128)
		if e != nil {
			t.Error(e)
		}
		if c != c128[i] {
			t.Fatalf("%v did not parse correctly to %v, instead: %v", str[i], c128[i], c)
		}
	}
	for i := range str {
		c, e := ParseComplex(str[i], 64)
		if e != nil {
			t.Error(e)
		}
		if complex64(c) != c64[i] {
			t.Fatalf("%v did not parse correctly to %v, instead: %v", str[i], c64[i], complex64(c))
		}
	}
	for i := range str {
		c, e := ParseComplex(fstr[i], 128)
		if e != nil {
			t.Error(e)
		}
		if c != c128[i] {
			t.Fatalf("%v did not parse correctly to %v, instead: %v", fstr[i], c128[i], c)
		}
	}
	for i := range str {
		c, e := ParseComplex(fstr[i], 64)
		if e != nil {
			t.Error(e)
		}
		if complex64(c) != c64[i] {
			t.Fatalf("%v did not parse correctly to %v, instead: %v", fstr[i], c64[i], complex64(c))
		}
	}
}

func TestFormatComplex(t *testing.T) {
	for i := range c128 {
		s := FormatComplex(c128[i])
		if s != fstr[i] {
			t.Fatalf("%v did not format correctly to %v, instead: %v", c128[i], fstr[i], s)
		}
	}
}
