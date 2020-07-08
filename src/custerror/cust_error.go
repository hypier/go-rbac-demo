package custerror

import (
	"runtime"
	"time"
)

type stack []uintptr

/**
1. 可传入更多参数
2. 可定位具体报错位置
3. 可输出报错调用链
4. 可对错误进行分类
5. 可自定义错误输出
*/
type customerError struct {
	msg  string
	when time.Time
	*stack
}

// 获取error的信息
func (b *customerError) Error() string {
	return b.msg
}

// 创建一个error
func New(msg string) error {
	return &customerError{msg: msg, when: time.Now(), stack: callers()}
}

func callers() *stack {
	const depth = 32
	var pcs [depth]uintptr
	n := runtime.Callers(3, pcs[:])
	var st stack = pcs[0:n]
	return &st

}
