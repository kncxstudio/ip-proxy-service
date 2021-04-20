package proxysvc

import (
	"fmt"
	"testing"
)

func TestGetProxyPool(t *testing.T) {
	fmt.Println(GetProxyPool())
}

func TestFormatProxy(t *testing.T) {
	fmt.Println(FormatProxy("142.93.16.163:3128 US-N-S + "))
}
