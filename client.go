package glocash

import "time"

type Client interface {
	PayClassic(data map[string]string) string
	PayDirect(data map[string]string) string
	Refund(data map[string]string) string
	Query(data map[string]string) string
}

// implement client
type Glocash struct {
	Environ      string
	Email        string
	MerchantName string
	Key          string
	Domain       string
	Timeout      time.Duration
	Scheme       string
	RequestParam map[string]string
	RespJson     string
	url          string
}
