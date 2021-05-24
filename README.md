# go-big-benchmark

This is a benchmark of the "big" golang package in relation to native float and int operations

## Results

Machine: Macbook Pro - Intel(R) Core(TM) i7-4850HQ CPU @ 2.30GHz

During my tests I observed the following:

* big.Int sums are 50x slower than with native int
* big.Float sums are 100x slower than with native float64
* big.Float multiplications and divisions are 50x slower than with native float64 for small numbers
* big.Float multiplications and divisions are just 5x slower than with native float64 for large numbers

## Usage

### Run in Docker container

* git clone this repository
* docker-compose build
* sample results from my Mac

```sh
goos: linux
goarch: amd64
pkg: github.com/flaviostutz/go-big-benchmark
cpu: Intel(R) Core(TM) i7-4850HQ CPU @ 2.30GHz
BenchmarkIntSumBig-6             	647555434	        17.95 ns/op	       0 B/op	       0 allocs/op
BenchmarkIntSumNative-6          	1000000000	         0.3401 ns/op	       0 B/op	       0 allocs/op
BenchmarkFloatSumBig-6           	98725372	       127.8 ns/op	      48 B/op	       1 allocs/op
BenchmarkFloatSumNative-6        	1000000000	         1.020 ns/op	       0 B/op	       0 allocs/op
BenchmarkFloatMulBig-6           	190391631	        56.58 ns/op	       0 B/op	       0 allocs/op
BenchmarkFloatMulNative-6        	1000000000	         1.682 ns/op	       0 B/op	       0 allocs/op
BenchmarkFloatDivBig-6           	62354732	       189.6 ns/op	      24 B/op	       2 allocs/op
BenchmarkFloatDivNative-6        	1000000000	         3.638 ns/op	       0 B/op	       0 allocs/op
BenchmarkLargeFloatMulBig-6      	1000000000	         7.362 ns/op	       0 B/op	       0 allocs/op
BenchmarkLargeFloatMulNative-6   	1000000000	         1.697 ns/op	       0 B/op	       0 allocs/op
```

### Run locally

* git clone this repository
* run

```sh
go test -bench=. ./... -benchmem -benchtime=10s
```

* sample results from my Mac

```sh
goos: darwin
goarch: amd64
pkg: github.com/flaviostutz/go-big-benchmark
BenchmarkIntSumBig-8             	610815481	        19.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkIntSumNative-8          	1000000000	         0.313 ns/op	       0 B/op	       0 allocs/op
BenchmarkFloatSumBig-8           	100000000	       110 ns/op	      48 B/op	       1 allocs/op
BenchmarkFloatSumNative-8        	1000000000	         0.933 ns/op	       0 B/op	       0 allocs/op
BenchmarkFloatMulBig-8           	200209824	        52.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkFloatMulNative-8        	1000000000	         1.57 ns/op	       0 B/op	       0 allocs/op
BenchmarkFloatDivBig-8           	80568111	       151 ns/op	      24 B/op	       2 allocs/op
BenchmarkFloatDivNative-8        	1000000000	         3.38 ns/op	       0 B/op	       0 allocs/op
BenchmarkLargeFloatMulBig-8      	1000000000	         6.86 ns/op	       0 B/op	       0 allocs/op
BenchmarkLargeFloatMulNative-8   	1000000000	         1.55 ns/op	       0 B/op	       0 allocs/op
```
