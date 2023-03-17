package response

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"strings"
)

/*
Response 实现error用于统一HTTP返回数据
未出错的情况下，Code统一返回1，Msg返回OK，接收数据写入Data
出错后则提取错误信息，Code用于定位错误位置，Msg为错误类别（可用于统一的提示回显），部分错误的报错信息（如缺失的字段）放在Data中
*/

type Response struct {
	Code int    `json:"code"`
	Msg  ErrMsg `json:"msg"`
	Data any    `json:"data,omitempty"`
}

func (r Response) Error() string {
	return string(r.Msg)
}

func Write(w http.ResponseWriter, err error, data any) {
	var body Response
	if err != nil {
		if e, ok := err.(Response); ok {
			body = e
		} else {
			body = Response{Code: -1, Msg: UnknownError}
		}
	} else {
		body = Response{Code: 1, Msg: OK, Data: data}
	}

	httpx.OkJson(w, body)
}

// ErrResp Logic内返回错误码从2开始，Handle占前两个位置
func ErrResp(pos, baseCode int, msg ErrMsg, note ...string) (resp Response) {
	resp.Code = baseCode + pos + 2
	resp.Msg = msg
	if len(note) > 0 {
		resp.Data = strings.Join(note, " ")
	}
	return
}
