//Package action ...
package action

import (
	"fmt"
	"os/exec"
)

func shell() {
	out, _ := exec.Command("bash", "-c", "docker ps > log.txt").Output()
	fmt.Println(string(out))
}
