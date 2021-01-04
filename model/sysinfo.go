package model

// OS 系统信息
type HostInfo struct {
	HostName string `json:"hostname"`
	IP       string `json:"ip"`
	OS       string `json:"os"`
	BootTime uint64 `json:"bootime"`
	Arch     string `json:"arch"`
	Cpu      CPU    `json:"cpu"`
	Disk     []Disk `json:"disk"`
	Memory   Memory `json:"memory"`
}

type CPU struct {
	ModelName string  `json:"modelname"`
	Mhz       float64 `json:"mhz"`
	Cores     int     `json:"cores"`
}

type Disk struct {
	Device         string  `json:"device"`
	MountPoint     string  `json:"mountpoint"`
	DiskSizeBytes  uint64  `json:"disk_size_bytes"`
	DiskAvailBytes uint64  `json:"disk_avail_bytes"`
	DiskPercent    float64 `json:"disk_percent"`
}

type Memory struct {
	TotlaSizeBytes uint64 `json:"totla_size_bytes"`
	FreeSizeBytes  uint64 `json:"free_size_bytes"`
}
