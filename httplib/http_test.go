package httplib

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	s, err := Get("http://www.baidu.com/robots.txt").String()
	fmt.Println(s, err)
}
