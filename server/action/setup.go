//Package action ...
package action

import (
	"fmt"
	"os/exec"
)

func setup() {
	cmd := string("bash script.sh")
	out, _ := exec.Command("bash", "-c", cmd).Output()
	fmt.Println(string(out))
}
