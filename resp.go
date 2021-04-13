package glocash

import (
	"encoding/json"
	"strconv"
)

// Data returned by the classic payment model
// you can use the judgment condition: whether UrlPayment is empty to determine whether the request is successful
type RespPayment struct {
	ReqError   string `json:"REQ_ERROR"`
	PgwMessage string `json:"PGW_MESSAGE"`
	TnsGcid    string `json:"TNS_GCID"`
	UrlPayment string `json:"URL_PAYMENT"`
}

// Data returned by server to server payment
// you can use the judgment condition: whether TnsGcid is empty to judge whether the request is successful
type RespDirect struct {
	ReqTimes    string `json:"REQ_TIMES"`
	ReqSandbox  string `json:"REQ_SANDBOX"`
	ReqInvoice  string `json:"REQ_INVOICE"`
	ReqSign     string `json:"REQ_SIGN"`
	ReqError    string `json:"REQ_ERROR"`
	TnsGcid     string `json:"TNS_GCID"`
	UrlPayment  string `json:"URL_PAYMENT"`
	TnsUtimes   string `json:"TNS_UTIME"`
	BilStatus   string `json:"BIL_STATUS"`
	BilMethod   string `json:"BIL_METHOD"`
	FdlDecision string `json:"FDL_DECISION"`
	FdlReview   string `json:"FDL_REVIEW"`
	PgwPrice    string `json:"PGW_PRICE"`
	PgwCurrency string `json:"PGW_CURRENCY"`
	PgwCc3ds    string `json:"PGW_CC3DS"`
	PgwMessage  string `json:"PGW_MESSGE"`
	PgwError    string `json:"PGW_ERROR"`
	CusEmail    string `json:"CUS_EMAIL"`
	CustomFd0   string `json:"CUSTOM_FD0"`
	CustomFd1   string `json:"CUSTOM_FD1"`
	CustomFd2   string `json:"CUSTOM_FD2"`
	CustomFd3   string `json:"CUSTOM_FD3"`
	CustomFd4   string `json:"CUSTOM_FD4"`
	CustomFd5   string `json:"CUSTOM_FD5"`
	CustomFd6   string `json:"CUSTOM_FD6"`
	CustomFd7   string `json:"CUSTOM_FD7"`
}

// RespRefund is different from other api returned data
// only RespRefund.ReqCode == 200 then this request is success
type RespRefund struct {
	ReqCode  int    `json:"REQ_CODE"`
	ReqError string `json:"REQ_ERROR"`
}

// Data returned by query
// you can use the judgment condition: whether TnsGcid is empty to judge whether the request is successful
type RespQuery struct {
	ReqTimes    int     `json:"REQ_TIMES"`
	ReqSandbox  string  `json:"REQ_SANDBOX"`
	ReqInvoice  string  `json:"REQ_INVOICE"`
	ReqSign     string  `json:"REQ_SIGN"`
	ReqError    string  `json:"REQ_ERROR"`
	ReqEmail    string  `json:"REQ_EMAIL"`
	TnsGcid     string  `json:"TNS_GCID"`
	TnsUtimes   float32 `json:"TNS_UTIMES"`
	BilStatus   string  `json:"BIL_STATUS"`
	BilMethod   string  `json:"BIL_METHOD"`
	BilPrice    string  `json:"BIL_PRICE"`
	BilCurrency string  `json:"BIL_CURRENCY"`
	FdlDecision string  `json:"FDL_DECISION"`
	FdlReview   string  `json:"FDL_REVIEW"`
	PgwPrice    string  `json:"PGW_PRICE"`
	PgwCurrency string  `json:"PGW_CURRENCY"`
	PgwCc3ds    string  `json:"PGW_CC3DS"`
	PgwMessage  string  `json:"PGW_MESSAGE"`
	PgwError    string  `json:"PGW_ERROR"`
	CusEmail    string  `json:"CUS_EMAIL"`
	CustomFd0   string  `json:"CUSTOM_FD0"`
	CustomFd1   string  `json:"CUSTOM_FD1"`
	CustomFd2   string  `json:"CUSTOM_FD2"`
	CustomFd3   string  `json:"CUSTOM_FD3"`
	CustomFd4   string  `json:"CUSTOM_FD4"`
	CustomFd5   string  `json:"CUSTOM_FD5"`
	CustomFd6   string  `json:"CUSTOM_FD6"`
	CustomFd7   string  `json:"CUSTOM_FD7"`
}

type RespQueryList struct {
	List     map[string]RespQuery
	ReqError string `json:"REQ_ERROR"`
}

// interface{} cover to string
func StrVal(value interface{}) string {
	var key string
	if value == nil {
		goto end
	}
	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = value.(string)
	case []byte:
		key = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}
end:
	return key
}
