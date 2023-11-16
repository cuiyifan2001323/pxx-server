// @Author 齐静春
// @Date 2023-11-14 22:17:00

package common

type Result struct {
	Data any    `json:"data"`
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func Success(data any, msg string) Result {
	return Result{
		Data: data,
		Msg:  msg,
		Code: 200,
	}
}

func Err(code int, msg string) Result {
	return Result{
		Data: nil,
		Code: code,
		Msg:  msg,
	}
}
