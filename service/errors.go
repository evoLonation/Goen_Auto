package service

import "errors"

// 这个err是具体的值，需要用errors.Is检查
var ErrPreConditionUnsatisfied = errors.New("pre condition dissatisfy")

// 这个err是一个实现了error的类型，需要用errors.As检查，同时实现了Wrapper
type ErrPostCondition struct {
	specificErr error
}

func (e *ErrPostCondition) Error() string {
	return "post condition error"
}

func (e *ErrPostCondition) Unwrap() error { return e.specificErr }

func NewErrPostCondition(err error) *ErrPostCondition {
	return &ErrPostCondition{specificErr: err}
}
