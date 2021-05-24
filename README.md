# go-big-benchmark

This is a benchmark of some golang decimal number packages in relation to native float and int operations. 

Currently evaluated packages:
* [Big](https://golang.org/pkg/math/big/)
* [Decimal](https://github.com/shopspring/decimal)

## Results

Machine: Macbook Pro - Intel(R) Core(TM) i7-4850HQ CPU @ 2.30GHz

During my tests I observed the following:

* big.Int sums are 50x slower than with native int
* decimal.Int sums are 300x slower than with native int
* big.Float sums are 100x slower than with native float64
* decimal.Float sums are 100x slower than with native float64
* big.Float multiplications and divisions are 50x slower than with native float64 for small numbers
* decimal.Float multiplications are 150.000x (!) slower than with native float64 for small numbers
* decimal.Float divisions are 50x slower than with native float64 for small numbers
* big.Float multiplications and divisions are just 5x slower than with native float64 for large numbers
* decimal.Float multiplications are 5x slower than with native float64 for large numbers
* decimal.Float divisions are 5x slower than with native float64 for large numbers
* there is no performance penaulty running in Docker container

## Usage

### Run in Docker container

* git clone this repository
* `docker-compose up --build`
* sample results from my Mac

```sh
goos: linux
goarch: amd64
pkg: github.com/flaviostutz/go-big-benchmark
cpu: Intel(R) Core(TM) i7-4850HQ CPU @ 2.30GHz
BenchmarkIntSumBig-6             	673371614	        17.85 ns/op	       0 B/op	       0 allocs/op
BenchmarkIntSumNative-6          	1000000000	         0.3456 ns/op	       0 B/op	       0 allocs/op
BenchmarkFloatSumBig-6           	93170161	       128.9 ns/op	      48 B/op	       1 allocs/op
BenchmarkFloatSumNative-6        	1000000000	         1.022 ns/op	       0 B/op	       0 allocs/op
BenchmarkFloatMulBig-6           	190275585	        55.97 ns/op	       0 B/op	       0 allocs/op
BenchmarkFloatMulNative-6        	1000000000	         1.684 ns/op	       0 B/op	       0 allocs/op
BenchmarkFloatDivBig-6           	61973674	       191.8 ns/op	      24 B/op	       2 allocs/op
BenchmarkFloatDivNative-6        	1000000000	         3.679 ns/op	       0 B/op	       0 allocs/op
BenchmarkLargeFloatMulBig-6      	1000000000	         7.322 ns/op	       0 B/op	       0 allocs/op
BenchmarkLargeFloatMulNative-6   	1000000000	         1.719 ns/op	       0 B/op	       0 allocs/op
BenchmarkLargeFloatSumBig-6      	95923098	       126.8 ns/op	      48 B/op	       1 allocs/op
BenchmarkLargeFloatSumNative-6   	1000000000	         1.036 ns/op	       0 B/op	       0 allocs/op
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
BenchmarkIntSumBig-8             	625236583	        19.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkIntSumNative-8          	1000000000	         0.311 ns/op	       0 B/op	       0 allocs/op
BenchmarkFloatSumBig-8           	100000000	       110 ns/op	      48 B/op	       1 allocs/op
BenchmarkFloatSumNative-8        	1000000000	         0.935 ns/op	       0 B/op	       0 allocs/op
BenchmarkFloatMulBig-8           	202502194	        51.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkFloatMulNative-8        	1000000000	         1.55 ns/op	       0 B/op	       0 allocs/op
BenchmarkFloatDivBig-8           	81446204	       148 ns/op	      24 B/op	       2 allocs/op
BenchmarkFloatDivNative-8        	1000000000	         3.32 ns/op	       0 B/op	       0 allocs/op
BenchmarkLargeFloatMulBig-8      	1000000000	         6.89 ns/op	       0 B/op	       0 allocs/op
BenchmarkLargeFloatMulNative-8   	1000000000	         1.55 ns/op	       0 B/op	       0 allocs/op
BenchmarkLargeFloatSumBig-8      	100000000	       107 ns/op	      48 B/op	       1 allocs/op
BenchmarkLargeFloatSumNative-8   	1000000000	         0.942 ns/op	       0 B/op	       0 allocs/op
```
