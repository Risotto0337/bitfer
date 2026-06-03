package core

import (
	"os"
	"os/exec"
)

func Build(steps []string, dir string) error {
	for _, step := range steps {
		cmd := exec.Command("bash", "-c", step)
		cmd.Dir = dir
		cmd.Env = append(os.Environ(),
			"DESTDIR="+dir+"/pkg")

		if err := cmd.Run(); err != nil {
			return err
		}
	}
	return nil
}
