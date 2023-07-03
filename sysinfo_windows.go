package sysinfo

import (
	"os"
	"runtime"
	"syscall"
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
	si.OSMajorVersion, si.OSMinorVersion, si.OSBuild = WinGetVersion()

	return si
}

func WinGetVersion() (int, int, int) {
	v, err := syscall.GetVersion()
	if err != nil {
		panic(err)
	}
	return int(uint8(v)), int(uint8(v >> 8)), int(uint16(v >> 16))
}
