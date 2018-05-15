//Package action ...
package action

<<<<<<< HEAD
import (
	"fmt"
	"os/exec"
)

func shell() {
	out, _ := exec.Command("bash", "-c", "docker ps > log.txt").Output()
	fmt.Println(string(out))
=======
func shell() {

>>>>>>> 7af5e911c5e4794f3f274991005513898cc4d6e5
}
