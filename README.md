# xorm_assassin

## to start benchmarking
8 cpus is used for the benchmark
go test -v -count=50 -bench=. -benchmem -benchtime=10000x > xxxx.txt

document: http://www.inanzzz.com/index.php/post/yz8n/using-golang-bench-benchstat-and-benchcmp-to-measure-performance