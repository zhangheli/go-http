package httplib

import (
	"context"
	"crypto/tls"
	"math/rand"
	"net"
	"net/url"
	"time"

	lib "github.com/astaxie/beego/httplib"
)

var (
	ReSolver *net.Resolver
)

func Get(target string) *lib.BeegoHTTPRequest {
	m, _ := url.Parse(target)
	host := m.Host
	return HandleReturn(lib.Get(NewUrl(m)), host)
}

func Post(target string) *lib.BeegoHTTPRequest {
	m, _ := url.Parse(target)
	host := m.Host
	return HandleReturn(lib.Post(NewUrl(m)), host)
}

func NewUrl(target *url.URL) string {
	return NewUrlWithDNS(target, "ip4", "114.114.114.114:53")
}

// NewUrlWithDNS
// ipProtocal -> "ip", "ip4", "ip6"
// dns -> 8.8.8.8:53
func NewUrlWithDNS(target *url.URL, ipProtocal, dns string) string {
	if ReSolver == nil {
		ReSolver = &net.Resolver{
			PreferGo: true,
			Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
				d := net.Dialer{
					Timeout: 10 * time.Second,
				}
				return d.DialContext(ctx, "udp", dns)
			},
		}
	}

	ip_list, err := ReSolver.LookupIP(context.Background(), ipProtocal, target.Host)
	if err != nil || len(ip_list) <= 0 {
		return target.String()
	}
	ip := ip_list[rand.Intn(len(ip_list))]
	target.Host = ip.String()
	return target.String()
}

func HandleReturn(req *lib.BeegoHTTPRequest, host string) *lib.BeegoHTTPRequest {
	return req.SetHost(host).SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
}
