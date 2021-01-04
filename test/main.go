package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	mountpoint()
	//regpoint()
}

func regpoint() {
	point := "/dev/termination-log"

	diskReg := "^/var/lib/kubelet|/boot|/var/lib/docker|/etc|/dev"
	if regexp.MustCompile(diskReg).MatchString(point) {
		fmt.Println("ok")
	}
}

func mountpoint() {
	str := "/hostfs"
	if strings.Contains(str, "/hostfs") {
		str = strings.TrimPrefix(str, "/hostfs")
		if len(str) != 0 {
			str = strings.TrimPrefix(str, "/hostfs")
		} else {
			str = str + "/"
		}
	}
	fmt.Println(str)
}
