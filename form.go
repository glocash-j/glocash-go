package glocash

import (
	"strconv"
	"time"
)

// All passable fields of the classic payment model
// Only the REQ_EMAIL,REQ_SIGN,REQ_TIMES,CUS_EMAIL,CUS_COUNTRY,
// CUS_FNAME,CUS_LNAME,BIL_METHOD,BIL_PRICE,BIL_CURRENCY,BIL_GOODSNAME,URL_SUCCESS,URL_NOTIFY are necessary
var FormPaymentCla = map[string]string{
	"REQ_EMAIL":            "",
	"REQ_MERCHANT":         "",
	"REQ_SIGN":             "",
	"REQ_SANDBOX":          "",
	"REQ_INVOICE":          "",
	"REQ_TIMES":            "",
	"MCH_URLPOST":          "",
	"CUS_EMAIL":            "",
	"CUS_COUNTRY":          "",
	"CUS_PHONE":            "",
	"CUS_MOBILE":           "",
	"CUS_IMUSR":            "",
	"CUS_STATE":            "",
	"CUS_CITY":             "",
	"CUS_ADDRESS":          "",
	"CUS_POSTAL":           "",
	"CUS_FNAME":            "",
	"CUS_LNAME":            "",
	"CUS_REGISTER":         "",
	"BIL_METHOD":           "",
	"BIL_PRICE":            "",
	"BIL_CURRENCY":         "",
	"BIL_PRCCODE":          "",
	"BIL_GOODSNAME":        "",
	"BIL_QUANTITY":         "",
	"BIL_CC3DS":            "",
	"URL_SUCCESS":          "",
	"URL_PENDING":          "",
	"URL_FAILED":           "",
	"URL_NOTIFY":           "",
	"URL_LOADING":          "",
	"URL_RETURN":           "",
	"BIL_SELLER_EMAIL":     "",
	"BIL_SELLER_URL":       "",
	"BIL_SELLER_GOODSNAME": "",
	"BIL_GOODS_URL":        "",
	"BIL_RAW_PRICE":        "",
	"BIL_RAW_CURRENCY":     "",
	"MCH_DOMAIN_KEY":       "",
	"REQ_MODE":             "",
	"IFS_MODE":             "",
	"IFS_URL":              "",
	"REC_STIME":            "",
	"REC_INPRICE":          "",
	"REC_PERIOD":           "",
	"REC_INTERVAL":         "",
	"REC_RETRIES":          "",
	"BIL_TYPE":             "",
	"CUSTOM_FD0":           "",
	"CUSTOM_FD1":           "",
	"CUSTOM_FD2":           "",
	"CUSTOM_FD3":           "",
	"CUSTOM_FD4":           "",
	"CUSTOM_FD5":           "",
	"CUSTOM_FD6":           "",
	"CUSTOM_FD7":           "",
}

// All passable fields of the direct payment model
// You need PCI certification to pay using this api
var FormPaymentDic = map[string]string{
	"REQ_TIMES":      "",
	"REQ_SIGN":       "",
	"REQ_EMAIL":      "",
	"REQ_INVOICE":    "",
	"REQ_MERCHANT":   "",
	"REQ_SANDBOX":    "",
	"MCH_URLPOST":    "",
	"CUS_EMAIL":      "",
	"CUS_COUNTRY":    "",
	"CUS_PHONE":      "",
	"CUS_MOBILE":     "",
	"CUS_IMUSR":      "",
	"CUS_STATE":      "",
	"CUS_CITY":       "",
	"CUS_ADDRESS":    "",
	"CUS_POSTAL":     "",
	"CUS_FNAME":      "",
	"CUS_LNAME":      "",
	"CUS_REGISTER":   "",
	"BIL_METHOD":     "",
	"BIL_PRICE":      "",
	"BIL_CURRENCY":   "",
	"BIL_PRCCODE":    "",
	"BIL_GOODSNAME":  "",
	"BIL_QUANTITY":   "",
	"BIL_GOODS_URL":  "",
	"BIL_CC3DS":      "",
	"BIL_IPADDR":     "",
	"BIL_CCNUMBER":   "",
	"BIL_CCHOLDER":   "",
	"BIL_CCEXPM":     "",
	"BIL_CCEXPY":     "",
	"BIL_CCCVV2":     "",
	"URL_SUCCESS":    "",
	"URL_PENDING":    "",
	"URL_FAILED":     "",
	"URL_NOTIFY":     "",
	"URL_LOADING":    "",
	"DEV_ACCEPT":     "",
	"DEV_UAGENT":     "",
	"BIL_REQ_KEY":    "",
	"MCH_DOMAIN_KEY": "",
	"REC_STIME":      "",
	"REC_INPRICE":    "",
	"REC_PERIOD":     "",
	"REC_INTERVAL":   "",
	"REC_RETRIES":    "",
	"BIL_TYPE":       "",
	"CUSTOM_FD0":     "",
	"CUSTOM_FD1":     "",
	"CUSTOM_FD2":     "",
	"CUSTOM_FD3":     "",
	"CUSTOM_FD4":     "",
	"CUSTOM_FD5":     "",
	"CUSTOM_FD6":     "",
	"CUSTOM_FD7":     "",
}

var FormRefund = map[string]string{
	"REQ_TIMES": "",
	"REQ_EMAIL": "",
	"TNS_GCID":  "",
	"PGW_PRICE": "",
}

var FormQuery = map[string]string{
	"REQ_TIMES":   "",
	"REQ_EMAIL":   "",
	"REQ_INVOICE": "",
	"TNS_GCID":    "",
}

// Combine map data to prepare post
func (g *Glocash) pack(tpl map[string]string, targetData map[string]string) *Glocash {
	g.RequestParam = map[string]string{}
	for k, v := range tpl {
		g.RequestParam[k] = v
	}
	for k, v := range targetData {
		_, ok := tpl[k]
		if !ok {
			panic(k + " is not allowed")
		}
		g.RequestParam[k] = v
	}
	if g.Environ == LIVE {
		g.RequestParam["REQ_SANDBOX"] = "0"
	} else {
		g.RequestParam["REQ_SANDBOX"] = "ON"
	}
	g.RequestParam["REQ_EMAIL"] = g.Email
	g.RequestParam["REQ_MERCHANT"] = g.MerchantName
	g.RequestParam["REQ_TIMES"] = strconv.FormatInt(time.Now().Unix(), 10)
	return g
}
