//Package action ...
package action

import (
	"fmt"
	"os/exec"
)

func setup() {
	str := string("ls" + " -a")
	out, _ := exec.Command("bash", "-c", str).Output()
	fmt.Println(string(out))
}
