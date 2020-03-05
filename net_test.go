package toolibs

import "testing"

func TestGetAddrs(t *testing.T) {
	addrs := GetAddrs()
	if len(addrs) <= 0 {
		t.Fail()
		t.Log("获取不到地址")
		return
	}
	t.Log(addrs)
}
