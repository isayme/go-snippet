package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(CallerInfo())
}

type callerInfo struct {
	file string
	line int
	fn   string
}

func CallerInfo() string {
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		return "-"
	}

	info := callerInfo{
		file: file,
		line: line,
	}

	funcForPc := runtime.FuncForPC(pc)
	if funcForPc != nil {
		info.fn = funcForPc.Name()
	}

	return info.String()
}

func (c callerInfo) String() string {
	return fmt.Sprintf("%s:%d:%s", c.file, c.line, c.fn)
}
