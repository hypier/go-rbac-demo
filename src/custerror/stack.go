package custerror

import (
	"fmt"
	"io"
	"path"
	"runtime"
	"strconv"
	"strings"
)

type stack []uintptr
type subStack uintptr

func callers() *stack {
	const depth = 32
	var pcs [depth]uintptr
	n := runtime.Callers(4, pcs[:])
	var st stack = pcs[0:n]
	return &st

}

func (sub subStack) Format(state fmt.State, verb rune) {
	switch verb {
	case 's':
		switch {
		case state.Flag('+'):
			io.WriteString(state, "\n\t")
			io.WriteString(state, sub.name())
			io.WriteString(state, "  ")
			io.WriteString(state, sub.file())
		default:
			io.WriteString(state, path.Base(sub.file()))
		}
	case 'd':
		io.WriteString(state, strconv.Itoa(sub.line()))
	case 'n':
		io.WriteString(state, funcName(sub.name()))
	case 'v':
		sub.Format(state, 's')
		io.WriteString(state, ":")
		sub.Format(state, 'd')
	}
}

func (sub subStack) name() string {
	fn := runtime.FuncForPC(sub.pc())
	if fn == nil {
		return "unknown"
	}
	return fn.Name()
}

func (sub subStack) file() string {
	fn := runtime.FuncForPC(sub.pc())
	if fn == nil {
		return "unknown"
	}
	file, _ := fn.FileLine(sub.pc())
	return file
}

func (sub subStack) line() int {
	fn := runtime.FuncForPC(sub.pc())
	if fn == nil {
		return 0
	}
	_, line := fn.FileLine(sub.pc())
	return line
}

func (sub subStack) pc() uintptr {
	return uintptr(sub) - 1
}

func funcName(name string) string {
	i := strings.LastIndex(name, "/")
	name = name[i+1:]
	i = strings.Index(name, ".")
	return name[i+1:]
}
