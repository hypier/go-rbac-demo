package custerror

import (
	"fmt"
	"io"
	"time"
)

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
func (c *customerError) Error() string {
	return c.msg
}

func (c *customerError) Format(state fmt.State, verb rune) {
	switch verb {
	case 'v':
		if state.Flag('+') {
			io.WriteString(state, c.msg)
			c.stackTrace(state, verb)
			return
		}
		fallthrough
	case 's':
		io.WriteString(state, c.msg)
	case 'q':
		fmt.Fprintf(state, "%q", c.msg)
	}
}

// 创建一个error
func New(msg string) error {
	return &customerError{msg: msg, when: time.Now(), stack: callers()}
}

func NewError(err error) error {
	return New(err.Error())
}

func (c *customerError) stackTrace(state fmt.State, verb rune) {
	switch verb {
	case 'v':
		switch {
		case state.Flag('+'):
			for _, pc := range *c.stack {
				f := subStack(pc)
				fmt.Printf("%+v", f)
				//fmt.Fprintf(state, "\n%+v", f)
			}
		}
	}

}

//func Newf(format string, args ...interface{}) error {
//	return &customerError{
//		msg:   fmt.Sprintf(format, args...),
//		stack: callers(),
//	}
//}
