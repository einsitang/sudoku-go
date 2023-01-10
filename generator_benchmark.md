# Generator Benchmark

command : `go test -bench=. ./generator/* -benchtime=5s`

mix only above level MEDIUM use concurrency `current`

```
goos: darwin
goarch: amd64
cpu: Intel(R) Core(TM) i7-6700HQ CPU @ 2.60GHz
BenchmarkGenerateLevelExpert
    generator_test.go:30: BenchmarkGenerateLevelExpert N : 1
    generator_test.go:39: generated (N:1) total time : 538 ms
    generator_test.go:30: BenchmarkGenerateLevelExpert N : 10
    generator_test.go:39: generated (N:10) total time : 2599 ms
    generator_test.go:30: BenchmarkGenerateLevelExpert N : 22
    generator_test.go:39: generated (N:22) total time : 5362 ms
BenchmarkGenerateLevelExpert-8   	      22	 243772013 ns/op
BenchmarkGenerateLevelHard
    generator_test.go:43: BenchmarkGenerateLevelHard N : 1
    generator_test.go:52: generated (N:1) total time : 14 ms
    generator_test.go:43: BenchmarkGenerateLevelHard N : 100
    generator_test.go:52: generated (N:100) total time : 2531 ms
    generator_test.go:43: BenchmarkGenerateLevelHard N : 236
    generator_test.go:52: generated (N:236) total time : 5734 ms
BenchmarkGenerateLevelHard-8     	     236	  24297489 ns/op
BenchmarkGenerateLevelMedium
    generator_test.go:56: BenchmarkGenerateLevelMedium N : 1
    generator_test.go:65: generated (N:1) total time : 50 ms
    generator_test.go:56: BenchmarkGenerateLevelMedium N : 100
    generator_test.go:65: generated (N:100) total time : 111 ms
    generator_test.go:56: BenchmarkGenerateLevelMedium N : 5364
    generator_test.go:65: generated (N:5364) total time : 6043 ms
BenchmarkGenerateLevelMedium-8   	    5364	   1126704 ns/op
BenchmarkGenerateLevelEasy
    generator_test.go:69: BenchmarkGenerateLevelEasy N : 1
    generator_test.go:78: generated (N:1) total time : 0 ms
    generator_test.go:69: BenchmarkGenerateLevelEasy N : 100
    generator_test.go:78: generated (N:100) total time : 79 ms
    generator_test.go:69: BenchmarkGenerateLevelEasy N : 7516
    generator_test.go:78: generated (N:7516) total time : 6113 ms
BenchmarkGenerateLevelEasy-8     	    7516	    813355 ns/op
PASS
ok  	command-line-arguments	30.098s

goos: darwin
goarch: amd64
cpu: Intel(R) Core(TM) i7-6700HQ CPU @ 2.60GHz
BenchmarkGenerateLevelExpert
    generator_test.go:30: BenchmarkGenerateLevelExpert N : 1
    generator_test.go:39: generated (N:1) total time : 38 ms
    generator_test.go:30: BenchmarkGenerateLevelExpert N : 100
    generator_test.go:39: generated (N:100) total time : 29797 ms
BenchmarkGenerateLevelExpert-8   	     100	 297975824 ns/op
BenchmarkGenerateLevelHard
    generator_test.go:43: BenchmarkGenerateLevelHard N : 1
    generator_test.go:52: generated (N:1) total time : 23 ms
    generator_test.go:43: BenchmarkGenerateLevelHard N : 100
    generator_test.go:52: generated (N:100) total time : 2177 ms
    generator_test.go:43: BenchmarkGenerateLevelHard N : 274
    generator_test.go:52: generated (N:274) total time : 5965 ms
BenchmarkGenerateLevelHard-8     	     274	  21773413 ns/op
BenchmarkGenerateLevelMedium
    generator_test.go:56: BenchmarkGenerateLevelMedium N : 1
    generator_test.go:65: generated (N:1) total time : 14 ms
    generator_test.go:56: BenchmarkGenerateLevelMedium N : 100
    generator_test.go:65: generated (N:100) total time : 122 ms
    generator_test.go:56: BenchmarkGenerateLevelMedium N : 4890
    generator_test.go:65: generated (N:4890) total time : 5546 ms
BenchmarkGenerateLevelMedium-8   	    4890	   1134310 ns/op
BenchmarkGenerateLevelEasy
    generator_test.go:69: BenchmarkGenerateLevelEasy N : 1
    generator_test.go:78: generated (N:1) total time : 0 ms
    generator_test.go:69: BenchmarkGenerateLevelEasy N : 100
    generator_test.go:78: generated (N:100) total time : 78 ms
    generator_test.go:69: BenchmarkGenerateLevelEasy N : 7620
    generator_test.go:78: generated (N:7620) total time : 6411 ms
BenchmarkGenerateLevelEasy-8     	    7620	    841371 ns/op
PASS
ok  	command-line-arguments	50.419s
```

