package httplib

import (
	"fmt"
	"net"
	"testing"
)

func TestQueryDns(t *testing.T) {
}

func TestCommonLib(t *testing.T) {
	ipRecords, _ := net.LookupIP("google.com")
	for _, ip := range ipRecords {
		fmt.Println(ip)
	}
}
