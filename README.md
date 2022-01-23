# sudoku-go

使用 `golang` 实现的数独`计算器`和`生成器`

## 使用

`require github.com/forfuns/sudoku-go`

### 计算器

输入 `[81]int8` 的数组题目,`-1`为需要的填空,`1`-`9`为题面，输出一个包含答案的 `Sudoku`

```go
// test case : main_test.go
import sudoku "github.com/forfuns/sudoku-go/core"

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

	_sudoku := new(sudoku.Sudoku)
	err := _sudoku.Init(puzzle)
	if err != nil {
		fmt.Println(err)
	} else {
		_sudoku.Debug()
        // t.Log("sudo puzzle is : %v",_sudoku.Puzzle())
		// t.Log("sudo answer is : %v",_sudoku.Answer())
		fmt.Println("sudoku is done")
	}
}
```

### 生成器

可以随机生成四种不同难度的数独题目 

分别为
- 简单 `LEVEL_EASY`
- 中等 `LEVEL_MEDIUM`
- 困难 `LEVEL_HARD`
- 大师 `LEVEL_EXPERT`

```go
// test case : generator_test.go
import generator "github.com/forfuns/sudoku-go/generator"

func main(){
    sudoku, err := generator.Generate(generator.LEVEL_EXPERT)
	if err != nil {
		fmt.Println(err)
	}
}
```

## More