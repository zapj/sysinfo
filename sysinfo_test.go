package sysinfo

import (
	"testing"
)

func TestSysinfo(t *testing.T) {
	si := GetSysInfo()
	t.Logf("%#v", si)

}
