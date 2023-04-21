package models

type Result struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	// 成功
	CODE_SUCCESS int = 1

	//失败
	CODE_ERROR int = 0

	//自定义...
)

func (res *Result) SetCode(code int) *Result {
	res.Code = code
	return res
}

func (res *Result) SetMessage(msg string) *Result {
	res.Msg = msg
	return res
}

func (res *Result) SetData(data interface{}) *Result {
	res.Data = data
	return res
}
