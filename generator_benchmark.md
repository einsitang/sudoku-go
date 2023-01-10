# Generator Benchmark

command : `go test -bench=. ./generator/* -benchtime=5s`

concurrency with goroutine : 8(CPU CORE)

```
goos: darwin
goarch: amd64
cpu: Intel(R) Core(TM) i7-6700HQ CPU @ 2.60GHz
BenchmarkGenerateLevelExpert-8   	       8	 754413634 ns/op
--- BENCH: BenchmarkGenerateLevelExpert-8
    generator_test.go:31: BenchmarkGenerateLevelExpert N : 1
    generator_test.go:40: generated (N:1) total time : 636 ms
    generator_test.go:31: BenchmarkGenerateLevelExpert N : 8
    generator_test.go:40: generated (N:8) total time : 6035 ms
BenchmarkGenerateLevelMedium-8   	     669	   9231044 ns/op
--- BENCH: BenchmarkGenerateLevelMedium-8
    generator_test.go:44: BenchmarkGenerateLevelMedium N : 1
    generator_test.go:53: generated (N:1) total time : 9 ms
    generator_test.go:44: BenchmarkGenerateLevelMedium N : 100
    generator_test.go:53: generated (N:100) total time : 895 ms
    generator_test.go:44: BenchmarkGenerateLevelMedium N : 669
    generator_test.go:53: generated (N:669) total time : 6175 ms
BenchmarkGenerateLevelHard-8     	     240	  27400842 ns/op
--- BENCH: BenchmarkGenerateLevelHard-8
    generator_test.go:57: BenchmarkGenerateLevelHard N : 1
    generator_test.go:66: generated (N:1) total time : 10 ms
    generator_test.go:57: BenchmarkGenerateLevelHard N : 100
    generator_test.go:66: generated (N:100) total time : 2489 ms
    generator_test.go:57: BenchmarkGenerateLevelHard N : 240
    generator_test.go:66: generated (N:240) total time : 6576 ms
BenchmarkGenerateLevelEasy-8     	     621	   9576866 ns/op
--- BENCH: BenchmarkGenerateLevelEasy-8
    generator_test.go:70: BenchmarkGenerateLevelEasy N : 1
    generator_test.go:79: generated (N:1) total time : 3 ms
    generator_test.go:70: BenchmarkGenerateLevelEasy N : 100
    generator_test.go:79: generated (N:100) total time : 964 ms
    generator_test.go:70: BenchmarkGenerateLevelEasy N : 621
    generator_test.go:79: generated (N:621) total time : 5946 ms
PASS
ok  	command-line-arguments	30.194s

goos: darwin
goarch: amd64
cpu: Intel(R) Core(TM) i7-6700HQ CPU @ 2.60GHz
BenchmarkGenerateLevelExpert-8   	       6	 855295180 ns/op
--- BENCH: BenchmarkGenerateLevelExpert-8
    generator_test.go:31: BenchmarkGenerateLevelExpert N : 1
    generator_test.go:40: generated (N:1) total time : 3043 ms
    generator_test.go:31: BenchmarkGenerateLevelExpert N : 2
    generator_test.go:40: generated (N:2) total time : 2632 ms
    generator_test.go:31: BenchmarkGenerateLevelExpert N : 3
    generator_test.go:40: generated (N:3) total time : 3127 ms
    generator_test.go:31: BenchmarkGenerateLevelExpert N : 4
    generator_test.go:40: generated (N:4) total time : 3630 ms
    generator_test.go:31: BenchmarkGenerateLevelExpert N : 6
    generator_test.go:40: generated (N:6) total time : 5131 ms
	... [output truncated]
BenchmarkGenerateLevelMedium-8   	     614	   9473716 ns/op
--- BENCH: BenchmarkGenerateLevelMedium-8
    generator_test.go:44: BenchmarkGenerateLevelMedium N : 1
    generator_test.go:53: generated (N:1) total time : 9 ms
    generator_test.go:44: BenchmarkGenerateLevelMedium N : 100
    generator_test.go:53: generated (N:100) total time : 974 ms
    generator_test.go:44: BenchmarkGenerateLevelMedium N : 614
    generator_test.go:53: generated (N:614) total time : 5816 ms
BenchmarkGenerateLevelHard-8     	     271	  21736504 ns/op
--- BENCH: BenchmarkGenerateLevelHard-8
    generator_test.go:57: BenchmarkGenerateLevelHard N : 1
    generator_test.go:66: generated (N:1) total time : 43 ms
    generator_test.go:57: BenchmarkGenerateLevelHard N : 100
    generator_test.go:66: generated (N:100) total time : 2206 ms
    generator_test.go:57: BenchmarkGenerateLevelHard N : 271
    generator_test.go:66: generated (N:271) total time : 5890 ms
BenchmarkGenerateLevelEasy-8     	     720	   8468314 ns/op
--- BENCH: BenchmarkGenerateLevelEasy-8
    generator_test.go:70: BenchmarkGenerateLevelEasy N : 1
    generator_test.go:79: generated (N:1) total time : 7 ms
    generator_test.go:70: BenchmarkGenerateLevelEasy N : 100
    generator_test.go:79: generated (N:100) total time : 832 ms
    generator_test.go:70: BenchmarkGenerateLevelEasy N : 720
    generator_test.go:79: generated (N:720) total time : 6097 ms
PASS
ok  	command-line-arguments	40.575s
```

concurrency with goroutine : 1

```
goos: darwin
goarch: amd64
cpu: Intel(R) Core(TM) i7-6700HQ CPU @ 2.60GHz
BenchmarkGenerateLevelExpert-8   	       2	8753649982 ns/op
--- BENCH: BenchmarkGenerateLevelExpert-8
    generator_test.go:31: BenchmarkGenerateLevelExpert N : 1
    generator_test.go:40: generated (N:1) total time : 3166 ms
    generator_test.go:31: BenchmarkGenerateLevelExpert N : 2
    generator_test.go:40: generated (N:2) total time : 17507 ms
BenchmarkGenerateLevelMedium-8   	    1677	   3621061 ns/op
--- BENCH: BenchmarkGenerateLevelMedium-8
    generator_test.go:44: BenchmarkGenerateLevelMedium N : 1
    generator_test.go:53: generated (N:1) total time : 1 ms
    generator_test.go:44: BenchmarkGenerateLevelMedium N : 100
    generator_test.go:53: generated (N:100) total time : 357 ms
    generator_test.go:44: BenchmarkGenerateLevelMedium N : 1677
    generator_test.go:53: generated (N:1677) total time : 6072 ms
BenchmarkGenerateLevelHard-8     	     306	  25170063 ns/op
--- BENCH: BenchmarkGenerateLevelHard-8
    generator_test.go:57: BenchmarkGenerateLevelHard N : 1
    generator_test.go:66: generated (N:1) total time : 14 ms
    generator_test.go:57: BenchmarkGenerateLevelHard N : 100
    generator_test.go:66: generated (N:100) total time : 1955 ms
    generator_test.go:57: BenchmarkGenerateLevelHard N : 306
    generator_test.go:66: generated (N:306) total time : 7702 ms
BenchmarkGenerateLevelEasy-8     	    2086	   2952080 ns/op
--- BENCH: BenchmarkGenerateLevelEasy-8
    generator_test.go:70: BenchmarkGenerateLevelEasy N : 1
    generator_test.go:79: generated (N:1) total time : 2 ms
    generator_test.go:70: BenchmarkGenerateLevelEasy N : 100
    generator_test.go:79: generated (N:100) total time : 287 ms
    generator_test.go:70: BenchmarkGenerateLevelEasy N : 2086
    generator_test.go:79: generated (N:2086) total time : 6157 ms
PASS
ok  	command-line-arguments	45.729s

goos: darwin
goarch: amd64
cpu: Intel(R) Core(TM) i7-6700HQ CPU @ 2.60GHz
BenchmarkGenerateLevelExpert-8   	       9	4548247150 ns/op
--- BENCH: BenchmarkGenerateLevelExpert-8
    generator_test.go:31: BenchmarkGenerateLevelExpert N : 1
    generator_test.go:40: generated (N:1) total time : 613 ms
    generator_test.go:31: BenchmarkGenerateLevelExpert N : 9
    generator_test.go:40: generated (N:9) total time : 40934 ms
BenchmarkGenerateLevelMedium-8   	    1616	   3428536 ns/op
--- BENCH: BenchmarkGenerateLevelMedium-8
    generator_test.go:44: BenchmarkGenerateLevelMedium N : 1
    generator_test.go:53: generated (N:1) total time : 2 ms
    generator_test.go:44: BenchmarkGenerateLevelMedium N : 100
    generator_test.go:53: generated (N:100) total time : 371 ms
    generator_test.go:44: BenchmarkGenerateLevelMedium N : 1616
    generator_test.go:53: generated (N:1616) total time : 5540 ms
BenchmarkGenerateLevelHard-8     	     198	  26367724 ns/op
--- BENCH: BenchmarkGenerateLevelHard-8
    generator_test.go:57: BenchmarkGenerateLevelHard N : 1
    generator_test.go:66: generated (N:1) total time : 6 ms
    generator_test.go:57: BenchmarkGenerateLevelHard N : 100
    generator_test.go:66: generated (N:100) total time : 3028 ms
    generator_test.go:57: BenchmarkGenerateLevelHard N : 198
    generator_test.go:66: generated (N:198) total time : 5220 ms
BenchmarkGenerateLevelEasy-8     	    2407	   2579200 ns/op
--- BENCH: BenchmarkGenerateLevelEasy-8
    generator_test.go:70: BenchmarkGenerateLevelEasy N : 1
    generator_test.go:79: generated (N:1) total time : 2 ms
    generator_test.go:70: BenchmarkGenerateLevelEasy N : 100
    generator_test.go:79: generated (N:100) total time : 249 ms
    generator_test.go:70: BenchmarkGenerateLevelEasy N : 2407
    generator_test.go:79: generated (N:2407) total time : 6208 ms
PASS
ok  	command-line-arguments	71.643s
```

mix only above level MEDIUM use concurrency `current`

```
goos: darwin
goarch: amd64
cpu: Intel(R) Core(TM) i7-6700HQ CPU @ 2.60GHz
BenchmarkGenerateLevelExpert-8   	      18	 963635523 ns/op
--- BENCH: BenchmarkGenerateLevelExpert-8
    generator_test.go:31: BenchmarkGenerateLevelExpert N : 1
    generator_test.go:40: generated (N:1) total time : 325 ms
    generator_test.go:31: BenchmarkGenerateLevelExpert N : 18
    generator_test.go:40: generated (N:18) total time : 17345 ms
BenchmarkGenerateLevelMedium-8   	    1600	   3980637 ns/op
--- BENCH: BenchmarkGenerateLevelMedium-8
    generator_test.go:44: BenchmarkGenerateLevelMedium N : 1
    generator_test.go:53: generated (N:1) total time : 6 ms
    generator_test.go:44: BenchmarkGenerateLevelMedium N : 100
    generator_test.go:53: generated (N:100) total time : 374 ms
    generator_test.go:44: BenchmarkGenerateLevelMedium N : 1600
    generator_test.go:53: generated (N:1600) total time : 6368 ms
BenchmarkGenerateLevelHard-8     	     237	  21261854 ns/op
--- BENCH: BenchmarkGenerateLevelHard-8
    generator_test.go:57: BenchmarkGenerateLevelHard N : 1
    generator_test.go:66: generated (N:1) total time : 57 ms
    generator_test.go:57: BenchmarkGenerateLevelHard N : 100
    generator_test.go:66: generated (N:100) total time : 2521 ms
    generator_test.go:57: BenchmarkGenerateLevelHard N : 237
    generator_test.go:66: generated (N:237) total time : 5039 ms
BenchmarkGenerateLevelEasy-8     	    1993	   2694292 ns/op
--- BENCH: BenchmarkGenerateLevelEasy-8
    generator_test.go:70: BenchmarkGenerateLevelEasy N : 1
    generator_test.go:79: generated (N:1) total time : 5 ms
    generator_test.go:70: BenchmarkGenerateLevelEasy N : 100
    generator_test.go:79: generated (N:100) total time : 300 ms
    generator_test.go:70: BenchmarkGenerateLevelEasy N : 1993
    generator_test.go:79: generated (N:1993) total time : 5369 ms
PASS
ok  	command-line-arguments	40.036s

goos: darwin
goarch: amd64
cpu: Intel(R) Core(TM) i7-6700HQ CPU @ 2.60GHz
BenchmarkGenerateLevelExpert-8   	      26	 893480413 ns/op
--- BENCH: BenchmarkGenerateLevelExpert-8
    generator_test.go:31: BenchmarkGenerateLevelExpert N : 1
    generator_test.go:40: generated (N:1) total time : 220 ms
    generator_test.go:31: BenchmarkGenerateLevelExpert N : 26
    generator_test.go:40: generated (N:26) total time : 23230 ms
BenchmarkGenerateLevelMedium-8   	    1592	   3613410 ns/op
--- BENCH: BenchmarkGenerateLevelMedium-8
    generator_test.go:44: BenchmarkGenerateLevelMedium N : 1
    generator_test.go:53: generated (N:1) total time : 5 ms
    generator_test.go:44: BenchmarkGenerateLevelMedium N : 100
    generator_test.go:53: generated (N:100) total time : 376 ms
    generator_test.go:44: BenchmarkGenerateLevelMedium N : 1592
    generator_test.go:53: generated (N:1592) total time : 5752 ms
BenchmarkGenerateLevelHard-8     	     278	  22245556 ns/op
--- BENCH: BenchmarkGenerateLevelHard-8
    generator_test.go:57: BenchmarkGenerateLevelHard N : 1
    generator_test.go:66: generated (N:1) total time : 14 ms
    generator_test.go:57: BenchmarkGenerateLevelHard N : 100
    generator_test.go:66: generated (N:100) total time : 2154 ms
    generator_test.go:57: BenchmarkGenerateLevelHard N : 278
    generator_test.go:66: generated (N:278) total time : 6183 ms
BenchmarkGenerateLevelEasy-8     	    1958	   2813535 ns/op
--- BENCH: BenchmarkGenerateLevelEasy-8
    generator_test.go:70: BenchmarkGenerateLevelEasy N : 1
    generator_test.go:79: generated (N:1) total time : 5 ms
    generator_test.go:70: BenchmarkGenerateLevelEasy N : 100
    generator_test.go:79: generated (N:100) total time : 306 ms
    generator_test.go:70: BenchmarkGenerateLevelEasy N : 1958
    generator_test.go:79: generated (N:1958) total time : 5508 ms
PASS
ok  	command-line-arguments	44.206s
```

