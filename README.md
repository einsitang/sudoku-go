# sudoku-go 
 [![License](https://img.shields.io/badge/License-Anti%20996-blue.svg)](https://github.com/996icu/996.ICU/blob/master/LICENSE) [![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0) [![Page Views Count](https://badges.toozhao.com/badges/01FT3Z973THHC20KF0D6MDQGWE/blue.svg)](https://badges.toozhao.com/stats/01FT3Z973THHC20KF0D6MDQGWE "Get your own page views count badge on badges.toozhao.com")

使用 `golang` 实现的数独`计算器`和`生成器`

opensource sudoku calculator and puzzle generator golang library

## 功能 features
- 数独解题器 - sodoku calculator  / solver
- 题目生成器 - random puzzle generator with goroutinue , multi-core support

## 使用 tutorial

`require github.com/einsitang/sudoku-go`

### 计算器 solver

输入 `[81]int8` 的数组题目,`-1`为需要的填空,`1`-`9`为题面，输出一个包含答案的 `Sudoku`

input `[81]int8` array and return a `sudoku.Sudoku`structure with full answer

use `-1` to mark the position mean **computation item**

#### Init

`Init` use backtrack algorithm to solve puzzle , no matter is many solution or not , find one then return

#### StrictInit

`StrictInit` is only can solve one solution sudoku puzzle , more then one will return error message : puzzle is not aone-solution sudoku , if you only want solve sudoku puzzle , just use `Init`

```golang
// test case : main_test.go
import sudoku "github.com/einsitang/sudoku-go/core"

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

  _sudoku := sudoku.Sudoku{}
  err := _sudoku.Init(puzzle)
  // err := _sudoku.StrictInit(puzzle)
  if err != nil {
    fmt.Println(err)
  } else {
   _sudoku.Debug()
    
    // origin puzzle
   _sudoku.Puzzle() 
    // with answer sudoku
   _sudoku.Answer()
 }
}
```

### 生成器 generator

可以随机生成四种不同难度的数独题目(唯一解数独)

make four level random one solution sudoku puzzle function `generator.Generate` 

#### level constant

- 简单 `LEVEL_EASY`
- 中等 `LEVEL_MEDIUM`
- 困难 `LEVEL_HARD`
- 大师 `LEVEL_EXPERT`

```golang
// test case : generator_test.go
import generator "github.com/einsitang/sudoku-go/generator"

func main(){
  sudoku1, err1 := generator.Generate(generator.LEVEL_EXPERT)
  if err1 != nil {
    fmt.Println(err1)
  }
  if err2 != nil {
    fmt.Println(err2)
  }
}
```

## More

with any idea welcome open issue to make me know

if you want same project with other language like js / dart and flutter app , here they are :
- [einsitang/sudoku-nodejs](https://github.com/einsitang/sudoku-nodejs)
- [einsitang/sudoku-dart](https://github.com/einsitang/sudoku-dart)
- [einsitang/sudoku-flutter](https://github.com/einsitang/sudoku-flutter)