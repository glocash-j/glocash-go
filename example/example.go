package main

import (
	glocash "../../glocash-go"
	"fmt"
	"strconv"
	"time"
)

func main() {
	glocash := &glocash.Glocash{
		Timeout: time.Second * 20,
		Key:     "your keys",
		Email:   "your account email",
		Domain:  "glocashpayment.com",
		Environ: glocash.LIVE,
		Scheme:  "https://",
	}

	// classic payment
	param := map[string]string{
		"REQ_INVOICE":   strconv.FormatInt(time.Now().Unix(), 10),
		"BIL_PRICE":     "98.03",
		"BIL_CURRENCY":  "USD",
		"BIL_GOODSNAME": "DEFAULT GOODS",
		"BIL_QUANTITY":  "0",
		"BIL_CC3DS":     "0",
		"BIL_METHOD":    "C01",
		"CUS_EMAIL":     "customemail@example.com",
		"CUS_COUNTRY":   "US",
		"CUS_ADDRESS":   "Dearborn Stree 8645 code forbidden",
		"CUS_POSTAL":    "854423",
		"URL_FAILED":    "http://localhost/failed",
		"URL_SUCCESS":   "http://localhost/success",
		"URL_PENDING":   "http://localhost/pending",
		"URL_NOTIFY":    "http://localhost/notify",
		"CUSTOM_FD1":    "custom field 1",
		"CUSTOM_FD2":    "custom field 2",
		"CUSTOM_FD3":    "custom field 3",
		"CUSTOM_FD4":    "custom field 4",
		"CUSTOM_FD5":    "custom field 5",
		"CUSTOM_FD6":    "custom field 6",
		"CUSTOM_FD7":    "custom field 7",
	}
	glocash.Construct()

	d := glocash.PayClassic(param)
	// d.UrlPayment is glocash payment page
	// user can input their credit info to pay order
	println(glocash.RespJson)
	if d.UrlPayment != "" {
		// you can location this url to complete payment
	} else {
		panic(d.ReqError)
	}

	/* use 3DS */
	// param["BIL_CC3DS"] = "1"
	// d := glocash.PayClassic(param)

	/* use server to server */
	param["BIL_CCNUMBER"] = "4200000000000000"
	param["BIL_CCHOLDER"] = "json bicker"
	param["BIL_CCEXPM"] = "06"
	param["BIL_CCEXPY"] = "2023"
	param["BIL_CCCVV2"] = "123"
	//custom user ip address
	param["BIL_IPADDR"] = "56.33.69.15"
	// goods detail url page
	param["BIL_GOODS_URL"] = "https://goods.info.com/detail/id/1"
	dd := glocash.PayDirect(param)
	if dd.UrlPayment == "" {
		// direct result
	} else {
		// you need location to dd.UrlPayment to complete 3DS safe check
	}

	/* query transaction list*/
	param = map[string]string{
		"TNS_GCID": "C014X13SH7B8HCMX", // If you want to query multiple transaction list, please combine multiple gcid with commas
		//  "TNS_GCID":"C014X13SH7B8HCMX,C014X13SH7B8CCVP",
	}
	c := glocash.Query(param)
	for gcid, trans := range c.List {
		println(gcid)
		println(trans.PgwMessage)
	}

	/* refund request */
	// refund all
	param = map[string]string{
		"TNS_GCID": "C014X13SH7B8HCMX",
	}
	// refund some
	param = map[string]string{
		"TNS_GCID":  "C014X13SH7B8HCMX",
		"PGW_PRICE": "2",
	}
	r := glocash.Refund(param)
	if r.ReqCode == 200 {
		// refund success
	} else {
		fmt.Println(r.ReqError)
	}

	// Example of verifying the legality of asynchronous notification data
	post := map[interface{}]interface{}{
		"REQ_EMAIL":    "xxxxxxxxxxx",
		"REQ_INVOICE":  "2sVQIHx2Za243",
		"CUS_EMAIL":    "lt@am.com",
		"BIL_PRICE":    "6.90",
		"BIL_CURRENCY": "CAD",
		"BIL_METHOD":   "CCL",
		"URL_NOTIFY":   "http://localhost/result.php?notify",
		"BIL_IPADDR":   "172.18.0.1",
		"CUSTOM_FD1":   "for",
		"CUSTOM_FD2":   "sendent",
		"CUSTOM_FD3":   "west",
		"CUSTOM_FD4":   "dcd",
		"PGW_PMTYPE":   "visa",
		"TNS_UTIMES":   1618194637.682819,
		"TNS_GCID":     "C014X13SH7B8HCMX",
		"REQ_REFERER":  "xxxxxxxxx",
		"BIP_COUNTRY":  "xxxxx",
		"BIP_PROXY":    "xxxxx e",
		"BIP_CITY":     "xxxxxx",
		"PGW_CCHOLDER": "FFSCF MOD",
		"PGW_MESSAGE":  "Request successfully processed",
		"BIL_STATUS":   "refunding",
		"PGW_CCNMASK":  "420000******0000",
		"PGW_PRICE":    "5.64",
		"PGW_CURRENCY": "USD",
		"FDL_DECISION": "ACP",
		"REQ_TIMES":    1618306203,
		"REQ_SIGN":     "8a9533fd06f3f58a50c7d03533930c874c7222105bd15f3c4efa59eab7409817",
	}
	if glocash.NotifySign(post) != post["REQ_SIGN"] {
		// Sign error
	}
}
