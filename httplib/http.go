package httplib

import (
	"math/rand"
	"net"
	"net/url"

	lib "github.com/astaxie/beego/httplib"
)

func Get(target string) *lib.BeegoHTTPRequest {
	m, _ := url.Parse(target)
	host := m.Host
	return lib.Get(NewUrl(m)).Header("Host", host)
}

func Post(target string) *lib.BeegoHTTPRequest {
	m, _ := url.Parse(target)
	host := m.Host
	return lib.Post(NewUrl(m)).Header("Host", host)
}

func NewUrl(target *url.URL) string {
	ip_list, err := net.LookupIP(target.Host)
	if err != nil || len(ip_list) <= 0 {
		return target.String()
	}
	ip := ip_list[rand.Intn(len(ip_list))]
	target.Host = ip.String()
	return target.String()
}
