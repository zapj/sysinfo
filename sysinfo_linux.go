package sysinfo

import (
	"os"
	"runtime"
)

func GetSysInfo() *SysInfo {
	si := &SysInfo{}
	var err error
	si.Hostname, err = os.Hostname()
	if err != nil {
		si.Hostname = ""
	}
	si.OSName = runtime.GOOS
	si.Arch = runtime.GOARCH
	si.NumCPU = runtime.NumCPU()
	return si
}
