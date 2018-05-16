//Package action ...
package action

import (
	"fmt"
	"os/exec"
)

func setup() {
	cmd := "make test"
	out, _ := exec.Command("/bin/bash", "-c", cmd).Output()
	fmt.Println(string(out))
}
