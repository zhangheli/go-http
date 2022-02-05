package httplib

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	s, err := Get("https://open.feishu.cn/open-apis/bot/v2/hook/672b021b-6056-441e-9ce6-9c4e4c49f51d").
		String()
	fmt.Println(s, err)
}
