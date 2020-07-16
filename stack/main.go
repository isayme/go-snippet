package main

import (
	"fmt"
	"runtime"
	"strings"
)

/*
 * https://lingchao.xin/post/go-errors-stack-traces.html
 */

func main() {
	fmt.Println(Stacktrace())
}

func Stacktrace() string {
	var pcs [32]uintptr
	n := runtime.Callers(2, pcs[:])
	st := pcs[0:n]

	var b strings.Builder
	for _, pc := range st {
		fn := runtime.FuncForPC(pc)
		b.WriteString("\n")
		f, n := fn.FileLine(pc)
		b.WriteString(fmt.Sprintf("%s:%d", f, n))
	}
	return b.String()
}
