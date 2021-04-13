package glocash

import (
	"encoding/json"
	"time"
)

// classic payment api
func (g *Glocash) PayClassic(data map[string]string) *RespPayment {
	g.pack(FormPaymentCla, data).paymentSign()
	urlPath := g.url + CLASSIC
	byt := PostForm(urlPath, g.RequestParam)
	g.RespJson = string(byt)
	d := &RespPayment{}
	err := json.Unmarshal(byt, d)
	if err != nil {
		panic(err)
	}
	return d
}

//use server to server payment api
func (g *Glocash) PayDirect(data map[string]string) *RespDirect {
	g.pack(FormPaymentDic, data).paymentSign()
	urlPath := g.url + DIRECT
	byt := PostForm(urlPath, g.RequestParam)
	g.RespJson = string(byt)
	d := &RespDirect{}
	err := json.Unmarshal(byt, d)
	if err != nil {
		panic(err)
	}
	return d
}

//refund success transaction
func (g *Glocash) Refund(data map[string]string) *RespRefund {
	g.pack(FormRefund, data).refundSign()
	urlPath := g.url + REFUND
	byt := PostForm(urlPath, g.RequestParam)
	g.RespJson = string(byt)
	d := &RespRefund{}
	err := json.Unmarshal(byt, d)
	if err != nil {
		panic(err)
	}
	return d
}

// query exist transaction info
func (g *Glocash) Query(data map[string]string) *RespQueryList {
	g.pack(FormQuery, data).querySign()
	urlPath := g.url + QUERY
	byt := PostForm(urlPath, g.RequestParam)
	g.RespJson = string(byt)
	d := &RespQueryList{}
	if g.RespJson != "" {
		byt = []byte(`{"List":` + g.RespJson + "}")
		err := json.Unmarshal(byt, d)
		if err != nil {
			panic(err)
		}
	}
	return d
}

// Construct for Glocash client
func (g *Glocash) Construct() *Glocash {
	if g.Environ == "" {
		g.Environ = LIVE
	}
	if g.Timeout == 0 {
		g.Timeout = 20 * time.Second
	}
	if g.Domain == "" {
		g.Domain = DOMAIN
	}
	if g.Key == "" {
		panic("Glocash key is must be filled")
	}
	if g.Email == "" {
		panic("Glocash email is must be filled")
	}
	if g.Scheme == "" {
		g.Scheme = SCHEME
	}
	g.url = g.Scheme + g.Environ + "." + g.Domain
	return g
}
