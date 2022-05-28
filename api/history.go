package api

import (
	"github.com/lhlyu/appstoreserverapi"
	"io"
	"net/http"
)

// History 获取历史交易记录
// Get Transaction History
// doc: https://developer.apple.com/documentation/appstoreserverapi/get_transaction_history
func History(w http.ResponseWriter, r *http.Request) {
	req, err := getReq(r)
	w.Header().Set("content-type", "application/json")
	if err != nil {
		io.WriteString(w, NewResp(nil, err).Json())
		return
	}
	c, err := appstoreserverapi.NewClient(&appstoreserverapi.Config{
		Iss: req.Iss,
		Kid: req.Kid,
		Bid: req.Bid,
		Pk:  req.Pk,
		Aud: req.Aud,
	})
	if err != nil {
		io.WriteString(w, NewResp(req, err).Json())
		return
	}
	result, err := c.ApiGetTransactionHistory(req.ID, false)
	if err != nil {
		io.WriteString(w, NewResp(req, err).Json())
		return
	}
	io.WriteString(w, NewResp(result, nil).Json())
}
