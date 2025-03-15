# sudoku-go 
[![codebeat badge](https://codebeat.co/badges/9beecf62-49dd-4eb5-9566-20cbf5c40143)](https://codebeat.co/projects/github-com-einsitang-sudoku-go-master)
[![Go Report Card](https://goreportcard.com/badge/github.com/einsitang/sudoku-go)](https://goreportcard.com/report/github.com/einsitang/sudoku-go)
[![License](https://img.shields.io/badge/License-Anti%20996-blue.svg)](https://github.com/996icu/996.ICU/blob/master/LICENSE)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![Page Views Count](https://badges.toozhao.com/badges/01FT3Z973THHC20KF0D6MDQGWE/blue.svg)](https://badges.toozhao.com/stats/01FT3Z973THHC20KF0D6MDQGWE "Get your own page views count badge on badges.toozhao.com")

使用 `golang` 实现的数独`解题器`和`生成器`

opensource sudoku solver and puzzle generator `golang` library

## 功能 features
- 数独解题器 - sodoku calculator  / solver
- 题目生成器 - random one-solution puzzle generator with goroutinue , multi-core support

## 安装 install

`go get github.com/einsitang/sudoku-go/v2@latest`

### 计算器 solver

输入 `[81]int8` 的数组题目,`-1`为需要的填空,`1`-`9`为题面，输出一个包含答案的 `Sudoku`

input `[81]int8` array and return a `sudoku.Sudoku`structure with full answer

use `-1` to mark the position mean **computation item**


```golang
// test case : core_test.go
import sudoku "github.com/einsitang/sudoku-go/v2/sudoku"

func main(){
 puzzle := [81]int8{
  -1, -1, 8 /* */, 9, -1, 6 /* */, -1, -1, 5,
  -1, 4, 3, -1 /* */, -1, -1, -1 /* */, 2, -1,
  -1, -1, -1 /* */, -1, -1, -1, -1 /* */, -1, -1,

  -1, -1, 4 /* */, -1, -1, -1 /* */, 9, -1, -1,
  5, -1, -1 /* */, -1, 4, -1 /* */, 6, 8, -1,
  -1, -1, -1 /* */, 1, -1, -1 /* */, -1, -1, -1,

  2, -1, -1 /* */, -1, 8, -1 /* */, -1, 7, -1,
  -1, -1, -1 /* */, -1, 3, 4 /* */, 1, -1, -1,
  -1, 6, -1 /* */, -1, -1, 9 /* */, -1, -1, -1,
 }

  _sudoku,err := sudoku.Solve(puzzle)

  // with [DFS] algorithm solve puzzle  #recommend#
  // _sudoku,err := sudoku.Solve(puzzle)

  // only solve with one-solution puzzle use this function
  // _sudoku,err := sudoku.Solve(puzzle,sudoku.WithStrict())

  // with [DLX] algorithm solve puzzle 
  // _sudoku,err := sudoku.Solve(puzzle,sudoku.WithDLX())
  
  if err != nil {
    fmt.Println(err)
  } else {
   _sudoku.Debug()
    
    // origin puzzle
   _sudoku.Puzzle() 
    // sudoku solution
   _sudoku.Solution()
 }
}
```

#### default solve

use backtrack algorithm to solve puzzle , no matter is many solution or not , find one then return

#### solve WithStrict

`WithStrict` is only can solve one-solution sudoku puzzle , more then one will return error message : puzzle is not one-solution sudoku , if you only want solve sudoku puzzle , just omit this parameter

#### solve WithDLX

~~`WithDLX` is use DLX algorithm to solve puzzle , only for very hard puzzle will faster , recommend use default way is well , and they not verify one-solution~~
There is no need to specifically establish the DLX mode.

Sudoku will calculate the difficulty of the puzzle and automatically select which algorithm to use for calculation.


### 生成器 generator

可以随机生成 **五** 种不同难度的数独题目(唯一解数独)

make **five** level random one-solution sudoku puzzle function `generator.Generate` 

[Generator Benchmark](./generator_benchmark.md)

#### level constant

- 简单 `LEVEL_EASY`
- 中等 `LEVEL_MEDIUM`
- 困难 `LEVEL_HARD`
- 大师 `LEVEL_EXPERT`
- "地狱" `LEVEL_HELL`

> "地狱" 难度的数独生成可能会非常慢,因为是数独的生成是完全离线且随机，花费太长时间将会严重耗损计算资源，所以在"地狱"难度耗费一定计算次数后仍然无法输出数独 , 则会降低其初定难度再次生成(大师 < 难度 < "地狱")，从而保证生成器能正常输出数独，因此耗时长度会有较大波动
> 
> `LEVE_HELL` There is a certain probability that it will take more than 500ms to complete. Be careful, it is using.
> 

```golang
// test case : generator_test.go
import sudoku "github.com/einsitang/sudoku-go/v2/sudoku"

func main(){
  _sudoku, err := sudoku.Generate(sudoku.LEVEL_EXPERT)
  if err != nil {
    fmt.Println(err)
  }
}
```

## More

with any idea welcome open issue to let me know

if you want same project with other language like js / dart and flutter app , here they are :
- [einsitang/sudoku-nodejs](https://github.com/einsitang/sudoku-nodejs)
- [einsitang/sudoku-dart](https://github.com/einsitang/sudoku-dart)
- [einsitang/sudoku-flutter](https://github.com/einsitang/sudoku-flutter)
