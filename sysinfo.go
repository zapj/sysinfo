package sysinfo

type SysInfo struct {
	Hostname       string
	Arch           string
	OSName         string
	OSVendor       string
	OSVersion      string
	OSRelease      string
	NumCPU         int
	OSMajorVersion int
	OSMinorVersion int
	OSBuild        int
}
