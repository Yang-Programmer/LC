package medium

import (
	"fmt"
	"testing"
)

func TestRestoreIpAddresses(t *testing.T) {
	ipAddresses := restoreIpAddresses("25525511135")
	fmt.Println(ipAddresses)
}
