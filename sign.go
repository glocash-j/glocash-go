package glocash

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

// Generate sha256 of asynchronous notification request parameters to verify whether the asynchronous notification is legal
func (g *Glocash) NotifySign(notifyData map[interface{}]interface{}) string {
	var p = map[string]string{}
	for k, v := range notifyData {
		p[StrVal(k)] = StrVal(v)
	}
	var signStr = g.Key + p["REQ_TIMES"] + p["REQ_EMAIL"] + p["CUS_EMAIL"] + p["TNS_GCID"] + p["BIL_STATUS"] +
		p["BIL_METHOD"] + p["PGW_PRICE"] + p["PGW_CURRENCY"]
	return g.sha256(signStr)
}

// Add payment sign str to  Glocash.RequestParam
func (g *Glocash) paymentSign() *Glocash {
	var signStr = g.Key + g.RequestParam["REQ_TIMES"] + g.RequestParam["REQ_EMAIL"] + g.RequestParam["REQ_INVOICE"] +
		g.RequestParam["CUS_EMAIL"] + g.RequestParam["BIL_METHOD"] + g.RequestParam["BIL_PRICE"] + g.RequestParam["BIL_CURRENCY"]
	g.RequestParam["REQ_SIGN"] = g.sha256(signStr)
	return g
}

// Add sign str to Glocash.RequestParam
func (g *Glocash) refundSign() *Glocash {
	var signStr = g.Key + g.RequestParam["REQ_TIMES"] + g.RequestParam["REQ_EMAIL"] + g.RequestParam["TNS_GCID"] +
		g.RequestParam["PGW_PRICE"]
	g.RequestParam["REQ_SIGN"] = g.sha256(signStr)
	return g
}

func (g *Glocash) querySign() *Glocash {
	var signStr = g.Key + g.RequestParam["REQ_TIMES"] + g.RequestParam["REQ_EMAIL"] + g.RequestParam["REQ_INVOICE"] +
		g.RequestParam["TNS_GCID"]
	g.RequestParam["REQ_SIGN"] = g.sha256(signStr)
	return g
}

func (g *Glocash) sha256(signStr string) string {
	sha := sha256.New()
	sha.Write([]byte(signStr))
	return strings.ToLower(string(hex.EncodeToString(sha.Sum(nil))))
}
