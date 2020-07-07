package custerror

type customerError struct {
	msg string
}

// 获取error的信息
func (b *customerError) Error() string {
	return b.msg
}

// 创建一个error
func New(msg string) error {
	return &customerError{msg}
}
