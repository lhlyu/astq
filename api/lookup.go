package api

import (
	"github.com/lhlyu/appstoreserverapi"
	"io"
	"net/http"
)

// Lookup 查找订单 ID
// Look Up Order ID
// doc: https://developer.apple.com/documentation/appstoreserverapi/look_up_order_id
func Lookup(w http.ResponseWriter, r *http.Request) {
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
	result, err := c.ApiLookUpOrderId(req.ID)
	if err != nil {
		io.WriteString(w, NewResp(req, err).Json())
		return
	}
	io.WriteString(w, NewResp(result, nil).Json())
}
