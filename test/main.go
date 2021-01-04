package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	point:="/etc/resolv.conf"

	diskReg := "^/var/lib/kubelet|/boot|/var/lib/docker|/etc|/dev$"
	if regexp.MustCompile(diskReg).MatchString(point) {
		fmt.Println("ok")
	}
}

func mountpoint() {
	str:="/platformData"
	if strings.Contains(str,"/hostfs/"){
		str = strings.TrimPrefix(str,"/hostfs")
	}
	fmt.Println(str)
}
