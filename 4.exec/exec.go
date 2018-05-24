package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("free")
	// deepin指令
	// free
	// name	args
	// ifconfig
	// ls	-a
	// ls
	// cat	/etc/deepin-version
	// uname -a
	// man	chmod
	// locale
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
}
