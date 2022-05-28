package api

import (
	"github.com/lhlyu/appstoreserverapi"
	"io"
	"net/http"
)

// Refund 获取退款历史
// Get Refund History
// doc: https://developer.apple.com/documentation/appstoreserverapi/get_refund_history
func Refund(w http.ResponseWriter, r *http.Request) {
	req, err := getReq(r)
	w.Header().Set("content-type", "application/json")
	if err != nil {
		io.WriteString(w, NewResp(nil, err).Json())
		return
	}
	c, err := appstoreserverapi.NewClient(&appstoreserverapi.Config{
		Iss:      req.Iss,
		Kid:      req.Kid,
		Bid:      req.Kid,
		Pk:       req.Pk,
		Aud:      req.Aud,
	})
	if err != nil {
		io.WriteString(w, NewResp(nil, err).Json())
		return
	}
	result,err := c.ApiGetRefundHistory(req.ID, false)
	if err != nil {
		io.WriteString(w, NewResp(nil, err).Json())
		return
	}
	io.WriteString(w, NewResp(result, nil).Json())
}
