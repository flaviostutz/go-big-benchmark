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
* decimal.Float divisions are 150x slower than with native float64 for small numbers
* big.Float multiplications and divisions are just 5x slower than with native float64 for large numbers
* decimal.Float multiplications are 150.000x (!) slower than with native float64 for large numbers
* decimal.Float divisions are 500x slower than with native float64 for large numbers
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
BenchmarkIntSumDecimal-4                  26741582               135.6 ns/op            80 B/op          2 allocs/op
BenchmarkIntSumBig-4                      190040193               19.26 ns/op            0 B/op          0 allocs/op
BenchmarkIntSumNative-4                   1000000000               0.3717 ns/op          0 B/op          0 allocs/op
BenchmarkFloatSumDecimal-4                26855610               141.0 ns/op            80 B/op          2 allocs/op
BenchmarkFloatSumBig-4                    22970008               137.8 ns/op            48 B/op          1 allocs/op
BenchmarkFloatSumNative-4                 1000000000               1.063 ns/op           0 B/op          0 allocs/op
BenchmarkFloatMulDecimal-4                  386100            325724 ns/op          625688 B/op          2 allocs/op
BenchmarkFloatMulBig-4                    50999302                71.85 ns/op            0 B/op          0 allocs/op
BenchmarkFloatMulNative-4                 1000000000               1.773 ns/op           0 B/op          0 allocs/op
BenchmarkFloatDivDecimal-4                 6842312               529.8 ns/op           216 B/op          8 allocs/op
BenchmarkFloatDivBig-4                    17923671               221.3 ns/op            24 B/op          2 allocs/op
BenchmarkFloatDivNative-4                 947187457                3.759 ns/op           0 B/op          0 allocs/op
BenchmarkLargeFloatMulDecimal-4             169734            312709 ns/op          559259 B/op          2 allocs/op
BenchmarkLargeFloatMulBig-4               129192493               23.44 ns/op            0 B/op          0 allocs/op
BenchmarkLargeFloatMulNative-4            1000000000               1.833 ns/op           0 B/op          0 allocs/op
BenchmarkLargeFloatSumBig-4               24095013               143.0 ns/op            48 B/op          1 allocs/op
BenchmarkLargeFloatSumNative-4            1000000000               1.129 ns/op           0 B/op          0 allocs/op
BenchmarkLargeFloatSumDecimal-4            6122847               591.0 ns/op           296 B/op          9 allocs/op
```

### Run locally

* git clone this repository
* run

```sh
go test -bench=. ./... -benchmem -benchtime=2s
```

(must run 2s bench because Decimal package is very slow on some operations and the test would hang)

* sample results from my Mac

```sh
goos: darwin
goarch: amd64
pkg: github.com/flaviostutz/go-big-benchmark
cpu: Intel(R) Core(TM) i7-4850HQ CPU @ 2.30GHz
BenchmarkIntSumDecimal-8                23467168               103.2 ns/op            80 B/op          2 allocs/op
BenchmarkIntSumBig-8                    136590232               16.93 ns/op            0 B/op          0 allocs/op
BenchmarkIntSumNative-8                 1000000000               0.3180 ns/op          0 B/op          0 allocs/op
BenchmarkFloatSumDecimal-8              23883771                99.04 ns/op           80 B/op          2 allocs/op
BenchmarkFloatSumBig-8                  22014676               111.2 ns/op            48 B/op          1 allocs/op
BenchmarkFloatSumNative-8               1000000000               0.9550 ns/op          0 B/op          0 allocs/op
BenchmarkFloatMulDecimal-8                510510            187914 ns/op          825997 B/op          2 allocs/op
BenchmarkFloatMulBig-8                  42061929                60.19 ns/op            0 B/op          0 allocs/op
BenchmarkFloatMulNative-8               1000000000               1.611 ns/op           0 B/op          0 allocs/op
BenchmarkFloatDivDecimal-8               5867126               405.6 ns/op           216 B/op          8 allocs/op
BenchmarkFloatDivBig-8                  14131184               168.9 ns/op            24 B/op          2 allocs/op
BenchmarkFloatDivNative-8               724208641                3.317 ns/op           0 B/op          0 allocs/op
BenchmarkLargeFloatMulDecimal-8           284547            210236 ns/op          934826 B/op          2 allocs/op
BenchmarkLargeFloatMulBig-8             37865743                56.33 ns/op            0 B/op          0 allocs/op
BenchmarkLargeFloatMulNative-8          1000000000               1.557 ns/op           0 B/op          0 allocs/op
BenchmarkLargeFloatSumDecimal-8          5476526               443.5 ns/op           296 B/op          9 allocs/op
BenchmarkLargeFloatSumBig-8             20575528               104.9 ns/op            48 B/op          1 allocs/op
BenchmarkLargeFloatSumNative-8          1000000000               0.9291 ns/op          0 B/op          0 allocs/op
```
