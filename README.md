# xorm_assassin

## to start benchmarking
8 cpus is used for the benchmark
go test -v -count=10 -bench=. -benchmem -benchtime=10000x > xxxx.txt

document: http://www.inanzzz.com/index.php/post/yz8n/using-golang-bench-benchstat-and-benchcmp-to-measure-performance

## xorm & goqu
benchstat xorm.txt goqu.txt

name           old time/op    new time/op    delta
DBExecution-8    2.39ms ± 3%    2.53ms ± 3%   +5.59%  (p=0.000 n=9+9)

name           old alloc/op   new alloc/op   delta
DBExecution-8    25.8kB ± 0%    29.3kB ± 0%  +13.82%  (p=0.000 n=8+10)

name           old allocs/op  new allocs/op  delta
DBExecution-8       810 ± 0%       569 ± 0%  -29.75%  (p=0.000 n=10+10)

## xorm & sqlx
benchstat xorm.txt sqlx.txt

name           old time/op    new time/op    delta
DBExecution-8    2.39ms ± 3%    2.42ms ± 3%     ~     (p=0.297 n=9+9)

name           old alloc/op   new alloc/op   delta
DBExecution-8    25.8kB ± 0%     8.9kB ± 0%  -65.46%  (p=0.002 n=8+10)

name           old allocs/op  new allocs/op  delta
DBExecution-8       810 ± 0%       316 ± 0%  -60.99%  (p=0.000 n=10+10)

## When we use in memory for sqlite
benchstat xorm.txt sqlx.txt

name           old time/op    new time/op    delta
DBExecution-8     151µs ± 1%      93µs ±16%  -38.66%  (p=0.000 n=8+10)

name           old alloc/op   new alloc/op   delta
DBExecution-8    25.9kB ± 0%     9.0kB ± 0%  -65.38%  (p=0.000 n=10+9)

name           old allocs/op  new allocs/op  delta
DBExecution-8       809 ± 0%       316 ± 0%  -60.94%  (p=0.000 n=10+10)

benchstat xorm.txt goqu.txt

name           old time/op    new time/op    delta
DBExecution-8     151µs ± 1%     144µs ± 0%   -5.22%  (p=0.000 n=8+9)

name           old alloc/op   new alloc/op   delta
DBExecution-8    25.9kB ± 0%    28.3kB ± 0%   +9.34%  (p=0.000 n=10+10)

name           old allocs/op  new allocs/op  delta
DBExecution-8       809 ± 0%       911 ± 0%  +12.61%  (p=0.000 n=10+9)