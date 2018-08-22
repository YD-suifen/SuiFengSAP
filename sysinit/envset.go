package sysinit

import (
	"fmt"
	"os/exec"
)

func Execshell() {

	cmd := exec.Command("sh", "init.sh")
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("system init finish")

}
