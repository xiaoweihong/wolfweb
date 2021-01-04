package service

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/xiaoweihong/wolfweb/model"
	"net"
	"regexp"
	"strings"
)

const (
	//排除/var/lib/kubelet,/boot,/etc,/dev挂载点
	diskReg = "^/var/lib/kubelet|/boot|/var/lib/docker|/etc|/dev"
)

func GetSysInfo() (sysinfo model.HostInfo, err error) {
	stat, err := host.Info()
	if err != nil {
		fmt.Println("stat", err)
		return sysinfo, err
	}
	sysinfo.HostName = stat.Hostname
	// 获取ip
	if ip, err := GetIP(); err != nil {
		fmt.Println("ip", err)
		return sysinfo, err
	} else {
		sysinfo.IP = ip
	}

	sysinfo.BootTime = stat.BootTime
	sysinfo.Arch = stat.KernelArch
	sysinfo.OS = fmt.Sprintf("%v-%v", stat.Platform, stat.PlatformVersion)
	// 获取cpu信息
	info, err := cpu.Info()
	if err != nil {
		fmt.Println("cpu", err)
		return sysinfo, err
	}
	counts, err := cpu.Counts(true)
	if err != nil {
		fmt.Println("count", err)
		return sysinfo, err
	}
	sysinfo.Cpu.ModelName = info[0].ModelName
	sysinfo.Cpu.Cores = counts
	sysinfo.Cpu.Mhz = info[0].Mhz

	// 获取磁盘信息
	partitions, _ := disk.Partitions(false)
	var disks []model.Disk
	var diskInfo model.Disk
	for _, p := range partitions {
		diskInfo.Device = p.Device
		d, err := disk.Usage(p.Mountpoint)
		if err != nil {
			fmt.Println("disk", err)
			fmt.Println("disk", diskInfo.MountPoint)
			return sysinfo, err
		}
		// docker容器挂载
		if strings.Contains(p.Mountpoint, "/hostfs") {
			tmpRoot := strings.TrimPrefix(p.Mountpoint, "/hostfs")
			if len(tmpRoot) == 0 {
				p.Mountpoint = "/"
			} else {
				p.Mountpoint = tmpRoot
			}
		}
		diskInfo.MountPoint = p.Mountpoint
		// 排除无用挂载点
		if regexp.MustCompile(diskReg).MatchString(p.Mountpoint) {
			continue
		}
		diskInfo.DiskSizeBytes = d.Total
		diskInfo.DiskAvailBytes = d.Free
		diskInfo.DiskPercent = d.UsedPercent
		disks = append(disks, diskInfo)
	}
	sysinfo.Disk = disks

	// 获取内存信息
	memory, err := mem.VirtualMemory()
	if err != nil {
		return sysinfo, err
	}
	sysinfo.Memory.TotlaSizeBytes = memory.Total
	sysinfo.Memory.FreeSizeBytes = memory.Free
	return sysinfo, nil
}

func GetIP() (ip string, err error) {
	addrSlice, err := net.InterfaceAddrs()
	if nil != err {
		return "127.0.0.1", err
	}
	for _, addr := range addrSlice {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if nil != ipnet.IP.To4() {
				ip := ipnet.IP.String()
				return ip, nil
			}
		}
	}
	return ip, err
}
