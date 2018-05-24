package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	test2()
}
func test1() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cmd := exec.Command("git")
		// deepin指令
		// google-chrome	谷歌
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

	})
	http.ListenAndServe(":8080", nil)
}
func test2() {
	cmd := exec.Command("ping", "baidu.com")
	cmd.Stdout = os.Stdout
	/* file, err := os.Open("4.exec/command.txt")
	if err != nil {
		os.Exit(1)
	}
	cmd.Stdin = file */
	cmd.Run()
}
