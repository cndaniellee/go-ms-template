package response

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

// response Implement from error
type response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data any `json:"data,omitempty"`
}

func (r response) Error() string {
	return r.Msg
}

func Write(w http.ResponseWriter, err error, data ...any) {
	var body response
	if err != nil {
		switch e := err.(type) {
		case response:
			body = e
		default:
			body = response{
				Code: -1,
				Msg: ServerError,
			}
		}
	} else {
		body = response{
			Code: 1,
			Msg: OK,
			Data: data,
		}
	}

	httpx.OkJson(w, body)
}

func ErrResp(pos, code int, msg string) (resp response) {
	resp.Code = code * 100 + pos + 2
	resp.Msg = msg
	return
}
