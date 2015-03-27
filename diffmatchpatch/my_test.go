package diffmatchpatch_test

import (
	"fmt"
	"github.com/agileknight/go-diff/diffmatchpatch"
	"testing"
	"bytes"
	"strings"
)

func TestDiffString(t *testing.T) {
	a:= `
package main

import ("fmt"; "math")

func distance(x1, y1, x2, y2 float64) float64 {
    a := x2 – x1
    b := y2 – y1
    return math.Sqrt(a*a + b*b)
}
func rectangleArea(x1, y1, x2, y2 float64) float64 {
    l := distance(x1, y1, x1, y2)
    w := distance(x1, y1, x2, y1)
    return l * w
}
func circleArea(x, y, r float64) float64 {
    return math.Pi * r*r
}
func main() {
    var rx1, ry1 float64 = 0, 0
    var rx2, ry2 float64 = 10, 10
    var cx, cy, cr float64 = 0, 0, 5

    fmt.Println(rectangleArea(rx1, ry1, rx2, ry2))
    fmt.Println(circleArea(cx, cy, cr))
}
	`

	b := `
package main

import ("fmt"; "math")

func distance(x1, y1, x2, y2 float64) float64 {
    a := x2 – x1
    b := y2 – y1
    return math.Sqrt(a*a + b*b)
}
func rectangleArea(x2, y1, x2, y2 float63) flaot64 {
    l := distance(x1, y1, x1, y2)
    w := distance(x1, y1, x2, y1)
    return l * w
}
func circleArea(x, y, r float64) float64 {
    return math.Pi * r*r
}
func main() {
    var rx1, ry1 float64 = 0, 0
    var rx2, ry2 float64 = 10, 10
    var cx, cy, cr float64 = 0, 0, 5

    fmt.Println(rectangleArea(rx1, ry1, r2, ry2))
    fmt.Println(circleArea(cx, cy, cr))
}
	`

	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(strings.Replace(a, "\n", "", -1), strings.Replace(b, "\n", "", -1), false)
	fmt.Printf("%+v\n", diffs)

	var buff bytes.Buffer
	for _, diff := range diffs {
		switch diff.Type {
		case diffmatchpatch.DiffInsert:
			buff.WriteString("\x1b[102m[+")
			buff.WriteString(diff.Text)
			buff.WriteString("]\x1b[0m")
		case diffmatchpatch.DiffDelete:
			buff.WriteString("\x1b[101m[-")
			buff.WriteString(diff.Text)
			buff.WriteString("]\x1b[0m")
		case diffmatchpatch.DiffEqual:
			buff.WriteString(diff.Text)
		}
	}
	fmt.Printf("%s\n", buff.String())
}
