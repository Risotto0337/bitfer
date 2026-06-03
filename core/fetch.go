package core

import (
	"os/exec"
)

func Fetch(url string, dest string) error {
	cmd := exec.Command("bash", "-c",
		"mkdir -p "+dest+" && cd "+dest+" && curl -L "+url+" | tar xz")

	return cmd.Run()
}
