// Copyright 2012 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT
// it seem DLX algorithm work for very hard(hell mode) puzzle solve is well , but simple puzzle use dfs will better
// and this algorithm can't verify one-soluion sudoku , so better not use , unless U ensure only solve puzzle without one-solution demand
package core

import "fmt"

func sudokuGoPuzzle2str(puzzle *[81]int8) string {
	puzzleStr := ""
	for _, p := range *puzzle {
		if p == -1 {
			puzzleStr = puzzleStr + "."
		} else {
			puzzleStr = puzzleStr + fmt.Sprint(p)
		}
	}
	return puzzleStr
}

func SudokuGo2str(sudoku *[81]int8) string {
	return sudokuGoPuzzle2str(sudoku)
}

func Str2sudokuGo(sudoku *string) (sudokuGo [81]int8) {
	for i, str := range *sudoku {
		if str == '.' {
			sudokuGo[i] = -1
		} else {
			sudokuGo[i] = int8(str - '0')
		}
	}
	return
}

func DLXSolve(puzzle [81]int8) string {
	return solve(sudokuGoPuzzle2str(&puzzle))
}

// solve puzzle in 81 character string format.
// if solved, result is 81 character string.
// if not solved, result is the empty string.
func solve(u string) string {
	// construct an dlx object with 324 constraint columns.
	// other than the number 324, this is not specific to sudoku.
	d := newDlxObject(324)
	// now add constraints that define sudoku rules.
	for r, i := 0, 0; r < 9; r++ {
		for c := 0; c < 9; c, i = c+1, i+1 {
			b := r/3*3 + c/3
			n := int(u[i] - '1')
			if n >= 0 && n < 9 {
				d.addRow([]int{i, 81 + r*9 + n, 162 + c*9 + n,
					243 + b*9 + n})
			} else {
				for n = 0; n < 9; n++ {
					d.addRow([]int{i, 81 + r*9 + n, 162 + c*9 + n,
						243 + b*9 + n})
				}
			}
		}
	}
	// run dlx.  not sudoku specific.
	d.search()
	// extract the sudoku-specific 81 character result from the dlx solution.
	return d.text()
}

// Knuth's data object
type x struct {
	c          *y
	u, d, l, r *x
	// except x0 is not Knuth's.  it's pointer to first constraint in row,
	// so that the sudoku string can be constructed from the dlx solution.
	x0 *x
}

// Knuth's column object
type y struct {
	x
	s int // size
	n int // name
}

// an object to hold the matrix and solution
type dlx struct {
	ch []y  // all column headers
	h  *y   // ch[0], the root node
	o  []*x // solution
}

// constructor creates the column headers but no rows.
func newDlxObject(nCols int) *dlx {
	ch := make([]y, nCols+1)
	h := &ch[0]
	d := &dlx{ch, h, nil}
	h.c = h
	h.l = &ch[nCols].x
	ch[nCols].r = &h.x
	nh := ch[1:]
	for i := range ch[1:] {
		hi := &nh[i]
		ix := &hi.x
		hi.n = i
		hi.c = hi
		hi.u = ix
		hi.d = ix
		hi.l = &h.x
		h.r = ix
		h = hi
	}
	return d
}

// rows define constraints
func (d *dlx) addRow(nr []int) {
	if len(nr) == 0 {
		return
	}
	r := make([]x, len(nr))
	x0 := &r[0]
	for x, j := range nr {
		ch := &d.ch[j+1]
		ch.s++
		np := &r[x]
		np.c = ch
		np.u = ch.u
		np.d = &ch.x
		np.l = &r[(x+len(r)-1)%len(r)]
		np.r = &r[(x+1)%len(r)]
		np.u.d, np.d.u, np.l.r, np.r.l = np, np, np, np
		np.x0 = x0
	}
}

// extracts 81 character sudoku string
func (d *dlx) text() string {
	b := make([]byte, len(d.o))
	for _, r := range d.o {
		x0 := r.x0
		b[x0.c.n] = byte(x0.r.c.n%9) + '1'
	}
	return string(b)
}

// the dlx algorithm
func (d *dlx) search() bool {
	h := d.h
	j := h.r.c
	if j == h {
		return true
	}
	c := j
	for minS := j.s; ; {
		j = j.r.c
		if j == h {
			break
		}
		if j.s < minS {
			c, minS = j, j.s
		}
	}

	cover(c)
	k := len(d.o)
	d.o = append(d.o, nil)
	for r := c.d; r != &c.x; r = r.d {
		d.o[k] = r
		for j := r.r; j != r; j = j.r {
			cover(j.c)
		}
		if d.search() {
			return true
		}
		r = d.o[k]
		c = r.c
		for j := r.l; j != r; j = j.l {
			uncover(j.c)
		}
	}
	d.o = d.o[:len(d.o)-1]
	uncover(c)
	return false
}

func cover(c *y) {
	c.r.l, c.l.r = c.l, c.r
	for i := c.d; i != &c.x; i = i.d {
		for j := i.r; j != i; j = j.r {
			j.d.u, j.u.d = j.u, j.d
			j.c.s--
		}
	}
}

func uncover(c *y) {
	for i := c.u; i != &c.x; i = i.u {
		for j := i.l; j != i; j = j.l {
			j.c.s++
			j.d.u, j.u.d = j, j
		}
	}
	c.r.l, c.l.r = &c.x, &c.x
}
