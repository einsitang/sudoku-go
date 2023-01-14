# Generator Benchmark

command : `go test -bench=. ./generator/* -benchtime=5s`

mix only above level MEDIUM use concurrency `current`

```
goos: darwin
goarch: amd64
cpu: Intel(R) Core(TM) i7-6700HQ CPU @ 2.60GHz
BenchmarkGenerateLevelExpert
    generator_test.go:30: BenchmarkGenerateLevelExpert N : 1
    generator_test.go:39: generated (N:1) total time : 6 ms
    generator_test.go:30: BenchmarkGenerateLevelExpert N : 100
    generator_test.go:39: generated (N:100) total time : 2945 ms
    generator_test.go:30: BenchmarkGenerateLevelExpert N : 202
    generator_test.go:39: generated (N:202) total time : 5537 ms
BenchmarkGenerateLevelExpert-8               202          27413273 ns/op
BenchmarkGenerateLevelHard
    generator_test.go:43: BenchmarkGenerateLevelHard N : 1
    generator_test.go:52: generated (N:1) total time : 6 ms
    generator_test.go:43: BenchmarkGenerateLevelHard N : 100
    generator_test.go:52: generated (N:100) total time : 606 ms
    generator_test.go:43: BenchmarkGenerateLevelHard N : 987
    generator_test.go:52: generated (N:987) total time : 5640 ms
BenchmarkGenerateLevelHard-8                 987           5715086 ns/op
BenchmarkGenerateLevelMedium
    generator_test.go:56: BenchmarkGenerateLevelMedium N : 1
    generator_test.go:65: generated (N:1) total time : 1 ms
    generator_test.go:56: BenchmarkGenerateLevelMedium N : 100
    generator_test.go:65: generated (N:100) total time : 107 ms
    generator_test.go:56: BenchmarkGenerateLevelMedium N : 5599
    generator_test.go:65: generated (N:5599) total time : 5988 ms
BenchmarkGenerateLevelMedium-8              5599           1069496 ns/op
BenchmarkGenerateLevelEasy
    generator_test.go:69: BenchmarkGenerateLevelEasy N : 1
    generator_test.go:78: generated (N:1) total time : 0 ms
    generator_test.go:69: BenchmarkGenerateLevelEasy N : 100
    generator_test.go:78: generated (N:100) total time : 82 ms
    generator_test.go:69: BenchmarkGenerateLevelEasy N : 7252
    generator_test.go:78: generated (N:7252) total time : 5655 ms
BenchmarkGenerateLevelEasy-8                7252            779866 ns/op
PASS
ok      command-line-arguments  26.764s

goos: darwin
goarch: amd64
cpu: Intel(R) Core(TM) i7-6700HQ CPU @ 2.60GHz
BenchmarkGenerateLevelExpert
    generator_test.go:30: BenchmarkGenerateLevelExpert N : 1
    generator_test.go:39: generated (N:1) total time : 5 ms
    generator_test.go:30: BenchmarkGenerateLevelExpert N : 100
    generator_test.go:39: generated (N:100) total time : 2499 ms
    generator_test.go:30: BenchmarkGenerateLevelExpert N : 240
    generator_test.go:39: generated (N:240) total time : 7017 ms
BenchmarkGenerateLevelExpert-8               240          29240097 ns/op
BenchmarkGenerateLevelHard
    generator_test.go:43: BenchmarkGenerateLevelHard N : 1
    generator_test.go:52: generated (N:1) total time : 3 ms
    generator_test.go:43: BenchmarkGenerateLevelHard N : 100
    generator_test.go:52: generated (N:100) total time : 577 ms
    generator_test.go:43: BenchmarkGenerateLevelHard N : 1039
    generator_test.go:52: generated (N:1039) total time : 5320 ms
BenchmarkGenerateLevelHard-8                1039           5120495 ns/op
BenchmarkGenerateLevelMedium
    generator_test.go:56: BenchmarkGenerateLevelMedium N : 1
    generator_test.go:65: generated (N:1) total time : 1 ms
    generator_test.go:56: BenchmarkGenerateLevelMedium N : 100
    generator_test.go:65: generated (N:100) total time : 110 ms
    generator_test.go:56: BenchmarkGenerateLevelMedium N : 5419
    generator_test.go:65: generated (N:5419) total time : 5656 ms
BenchmarkGenerateLevelMedium-8              5419           1043884 ns/op
BenchmarkGenerateLevelEasy
    generator_test.go:69: BenchmarkGenerateLevelEasy N : 1
    generator_test.go:78: generated (N:1) total time : 0 ms
    generator_test.go:69: BenchmarkGenerateLevelEasy N : 100
    generator_test.go:78: generated (N:100) total time : 75 ms
    generator_test.go:69: BenchmarkGenerateLevelEasy N : 7908
    generator_test.go:78: generated (N:7908) total time : 6251 ms
BenchmarkGenerateLevelEasy-8                7908            790503 ns/op
PASS
ok      command-line-arguments  27.686s
```

