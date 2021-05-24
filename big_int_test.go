package main

import (
	"math/big"
	"testing"
)

var (
	bigir  *big.Int
	bigif  *big.Float
	intr   int
	floatr float64
)

func BenchmarkIntSumBig(b *testing.B) {
	z := big.NewInt(0)
	x, y := big.NewInt(0), big.NewInt(1)
	for i := 0; i < b.N; i++ {
		x = x.Add(x, y)
		//avoid compiler optimizations
		z = x
	}
	//avoid compiler optimizations
	bigir = z
}

func BenchmarkIntSumNative(b *testing.B) {
	r := 0
	x, y := 0, 1
	for i := 0; i < b.N; i++ {
		x = x + y
		//avoid compiler optimizations
		r = x
	}
	//avoid compiler optimizations
	intr = r
}

func BenchmarkFloatSumBig(b *testing.B) {
	z := big.NewFloat(0)
	x, y := big.NewFloat(1234.5678), big.NewFloat(5678.1234)
	for i := 0; i < b.N; i++ {
		x = x.Add(x, y)
		//avoid compiler optimizations
		z = x
	}
	//avoid compiler optimizations
	bigif = z
}

func BenchmarkFloatSumNative(b *testing.B) {
	r := 0.0
	x, y := 1234.5678, 5678.1234
	for i := 0; i < b.N; i++ {
		x = x + y
		//avoid compiler optimizations
		r = x
	}
	//avoid compiler optimizations
	floatr = r
}

func BenchmarkFloatMulBig(b *testing.B) {
	z := big.NewFloat(0)
	x, y := big.NewFloat(1234.5678), big.NewFloat(5678.1234)
	for i := 0; i < b.N; i++ {
		x = x.Mul(x, y)
		//avoid compiler optimizations
		z = x
	}
	//avoid compiler optimizations
	bigif = z
}

func BenchmarkFloatMulNative(b *testing.B) {
	r := 0.0
	x, y := 1234.5678, 5678.1234
	for i := 0; i < b.N; i++ {
		x = x * y
		//avoid compiler optimizations
		r = x
	}
	//avoid compiler optimizations
	floatr = r
}

func BenchmarkFloatDivBig(b *testing.B) {
	z := big.NewFloat(0)
	x, y := big.NewFloat(1234.5678), big.NewFloat(5678.1234)
	for i := 0; i < b.N; i++ {
		x = x.Quo(x, y)
		//avoid compiler optimizations
		z = x
	}
	//avoid compiler optimizations
	bigif = z
}

func BenchmarkFloatDivNative(b *testing.B) {
	r := 0.0
	x, y := 1234.5678, 5678.1234
	for i := 0; i < b.N; i++ {
		x = x / y
		//avoid compiler optimizations
		r = x
	}
	//avoid compiler optimizations
	floatr = r
}

func BenchmarkLargeFloatMulBig(b *testing.B) {
	z := big.NewFloat(0)
	x, y := big.NewFloat(12345879678.2342342342342), big.NewFloat(5674298865645645654.12342342423423)
	for i := 0; i < b.N; i++ {
		x = x.Mul(x, y)
		//avoid compiler optimizations
		z = x
	}
	//avoid compiler optimizations
	bigif = z
}

func BenchmarkLargeFloatMulNative(b *testing.B) {
	r := 0.0
	x, y := 12345879678.2342342342342, 5674298865645645654.12342342423423
	for i := 0; i < b.N; i++ {
		x = x * y
		//avoid compiler optimizations
		r = x
	}
	//avoid compiler optimizations
	floatr = r
}

func BenchmarkLargeFloatSumBig(b *testing.B) {
	z := big.NewFloat(0)
	x, y := big.NewFloat(12345879678.2342342342342), big.NewFloat(5674298865645645654.12342342423423)
	for i := 0; i < b.N; i++ {
		x = x.Add(x, y)
		//avoid compiler optimizations
		z = x
	}
	//avoid compiler optimizations
	bigif = z
}

func BenchmarkLargeFloatSumNative(b *testing.B) {
	r := 0.0
	x, y := 12345879678.2342342342342, 5674298865645645654.12342342423423
	for i := 0; i < b.N; i++ {
		x = x + y
		//avoid compiler optimizations
		r = x
	}
	//avoid compiler optimizations
	floatr = r
}
