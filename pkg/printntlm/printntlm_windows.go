// +build windows

package printntlm

import (
	"bytes"
	"fmt"
	"os/exec"
)

func SelfDAV(port int) {
	args := fmt.Sprintf(" /c net use * http://127.0.0.1:%d", port)
	cmd := exec.Command("cmd.exe", args)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	cmd.Run()
	fmt.Println("Ran: ", cmd, args)
}
