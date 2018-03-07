package service

import (
	"fmt"
	"testing"
)

func TestAccountService_Get(t *testing.T) {
	fmt.Println(Account.Get("hal_shutter"))
}
