package sysinfo

type SysInfo struct {
	Hostname  string
	Arch      string
	OSName    string
	OSVersion string
	MaxProcs  int
	NumCPU    int
}
