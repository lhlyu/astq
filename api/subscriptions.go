package api

import (
	"encoding/json"
	"github.com/lhlyu/appstoreserverapi"
	"io"
	"io/ioutil"
	"net/http"
)

type Req struct {
	// 发行人: 您在 App Store Connect 中的密钥页面中的发行者 ID（例如：" 57246542-96fe-1a63-e053-0824d011072a"）
	// Issuer: Your issuer ID from the Keys page in App Store Connect (Ex: "57246542-96fe-1a63-e053-0824d011072a")
	Iss string `json:"iss"`
	// 秘钥：您在 App Store Connect 中的私钥 ID（例如2X9R4HXF34：）
	// Key ID: Your private key ID from App Store Connect (Ex: 2X9R4HXF34)
	Kid string `json:"kid"`
	// 应用的BundleID（例如：“com.example.testbundleid2021”)
	// Bundle ID: Your app’s bundle ID (Ex: “com.example.testbundleid2021”)
	Bid string `json:"bid"`
	// 签名的秘钥
	// sign private key, eg:
	/*
		-----BEGIN PRIVATE KEY-----
		MIGTAg23kjjh2h3uhuhfduhJHAKJ23JASjhaskjj234hjHKJHS31hkjj
		-----END PRIVATE KEY-----`
	*/
	Pk string `json:"pk"`
	// 受众：appstoreconnect-v1
	// Audience: appstoreconnect-v1
	Aud string `json:"aud"`

	// orderId | transactionId
	ID string `json:"id"`
}

func getReq(r *http.Request) (*Req, error) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	r.Body.Close()
	req := &Req{}
	if err := json.Unmarshal(b, req); err != nil {
		return nil, err
	}
	return req, nil
}

type Resp struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data string `json:"data"`
}

func NewResp(data interface{}, err error) *Resp {
	if err != nil {
		return &Resp{
			Code: 1,
			Msg:  err.Error(),
			Data: "",
		}
	}
	b,err := json.Marshal(data)
	if err != nil {
		return &Resp{
			Code: 2,
			Msg:  err.Error(),
			Data: "",
		}
	}
	return &Resp{
		Code: 0,
		Msg:  "success",
		Data: string(b),
	}
}

func (r *Resp) Json() string {
	b,_ := json.Marshal(r)
	return string(b)
}

// Subscriptions 获取所有的订阅状态
// Get All Subscription Statuses
// doc: https://developer.apple.com/documentation/appstoreserverapi/get_all_subscription_statuses
func Subscriptions(w http.ResponseWriter, r *http.Request) {
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
	result,err := c.ApiGetAllSubscriptionStatuses(req.ID)
	if err != nil {
		io.WriteString(w, NewResp(nil, err).Json())
		return
	}
	io.WriteString(w, NewResp(result, nil).Json())
}
